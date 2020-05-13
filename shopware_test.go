package shopware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	server   *httptest.Server
	token    = "SWSCOGXOAGGXCMLNODRQTMRQAA"    // testing endpoint
	endpoint = "https://hundemode.julian.pro/" // testing endpoint
)

func checkHeaders(req *http.Request, assert *assert.Assertions) {
	assert.Equal(token, req.Header.Get("sw-access-key"))
	assert.Equal("application/json", req.Header.Get("Content-Type"))
	assert.Equal(fmt.Sprintf("sdk shopware.go/%s", Version), req.Header.Get("X-Shopware-User-Agent"))
}

func TestShopwareNew(t *testing.T) {
	assert := assert.New(t)

	api := New(token, endpoint)
	assert.IsType(Client{}, *api)
	assert.Equal(endpoint, api.BaseURL)
	assert.Equal(token, api.token)
	assert.Equal(token, api.Headers["sw-access-key"])
	assert.Equal("application/json", api.Headers["Content-Type"])
	assert.Equal(fmt.Sprintf("sdk shopware.go/%s", Version), api.Headers["X-Shopware-User-Agent"])
}
