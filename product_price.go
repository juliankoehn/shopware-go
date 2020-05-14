package shopware

// ProductPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductPrice/ProductPriceEntity.php
type ProductPrice struct {
	PriceRule
	ProductID     string   `json:"productId"`
	QuantityStart int      `json:"quantityStart"`
	QuantityEnd   int      `json:"quantityEnd"`
	Product       *Product `json:"product"`
	Rule          Rule     `json:"rule"`
}
