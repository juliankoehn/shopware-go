package shopware

// ListPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/ListPrice.php
type ListPrice struct {
	Price      float64 `json:"price"`
	Discount   float64 `json:"discount"`
	Percentage float64 `json:"percentage"`
}
