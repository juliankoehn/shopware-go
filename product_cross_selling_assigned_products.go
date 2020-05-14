package shopware

// ProductCrossSellingAssignedProduct see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductCrossSellingAssignedProducts/ProductCrossSellingAssignedProductsEntity.php
type ProductCrossSellingAssignedProduct struct {
	Entity
	CrossSellingID string               `json:"crossSellingId"`
	ProductID      string               `json:"productId"`
	Product        *Product             `json:"product"`
	CrossSelling   *ProductCrossSelling `json:"crossSelling"`
	Position       int                  `json:"position"`
}
