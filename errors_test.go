package shopware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundErrorResponse(t *testing.T) {
	var err error
	assert := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, string(readTestData("error-notfound.json")))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	api := New(token, server.URL)

	// test product
	_, err = api.Products.Get("unknown-product-id")
	assert.NotNil(err)
	_, ok := err.(NotFoundError)
	t.Logf(err.Error())
	assert.Equal(true, ok)
	notFoundError := err.(NotFoundError)
	assert.Equal(404, notFoundError.APIError.res.StatusCode)
}
