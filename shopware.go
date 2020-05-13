package shopware

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client model
type Client struct {
	client        *http.Client
	BaseURL       string
	token         string
	QueryParams   map[string]string
	Headers       map[string]string
	commonService service
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
			"Content-Type":          "application/json",
			"X-Shopware-User-Agent": fmt.Sprintf("sdk shopware.go/%s", Version),
		},
	}

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
