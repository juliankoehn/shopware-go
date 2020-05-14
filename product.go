package shopware

import (
	"fmt"
	"net/http"
)

// ProductService model
type ProductsService service

// Product model
type Product struct{}

// Get returns a single product
func (service *ProductsService) Get(productID string) (*Product, error) {
	path := fmt.Sprintf("/sales-channel-api/v1/product/%s", productID)

	req, err := service.c.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var product Product
	if ok := service.c.do(req, &product); ok != nil {
		return nil, ok
	}

	return &product, nil
}
