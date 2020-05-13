package shopware

import "fmt"

// ProductService model
type ProductService service

// Product model
type Product struct{}

// Get returns a single product
func (service *ProductService) Get(productID string) {
	path := fmt.Sprintf("/sales-channel-api/v1/product/%s", productID)
}
