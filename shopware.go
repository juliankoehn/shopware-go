package shopware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"moul.io/http2curl"
)

// Client model
type Client struct {
	client        *http.Client
	BaseURL       string
	token         string
	Debug         bool
	QueryParams   map[string]string
	Headers       map[string]string
	commonService service

	Products *ProductsService
}

type service struct {
	c *Client
}

// New initialized a new API Client
func New(token, endpoint string) *Client {
	c := &Client{
		client:  http.DefaultClient,
		token:   token,
		BaseURL: endpoint,
		Headers: map[string]string{
			"sw-access-key":         token,
			"Accept":                "application/json",
			"content-type":          "application/json",
			"X-Shopware-User-Agent": fmt.Sprintf("sdk shopware.go/%s", Version),
		},
	}

	c.commonService.c = c

	c.Products = (*ProductsService)(&c.commonService)

	return c
}

// SetHTTPClient sets the underlying http.Client used to make requests.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.client = client
}

func (c *Client) newRequest(method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// set query params
	for key, value := range c.QueryParams {
		query.Set(key, value)
	}

	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	if c.Debug == true {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode < 400 {
		if v != nil {
			defer res.Body.Close()
			err = json.NewDecoder(res.Body).Decode(v)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// parse api response
	apiError := c.handleError(req, res)

	return apiError
}

func (c *Client) handleError(req *http.Request, res *http.Response) error {
	if c.Debug == true {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%q", dump)
	}

	var e ErrorResponse
	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&e)
	if err != nil {
		return err
	}

	apiError := APIError{
		req: req,
		res: res,
		err: &e,
	}

	switch res.StatusCode {
	case 404:
		return NotFoundError{apiError}
	default:
		return e
	}
}
