package shopware

// ProductVisibility see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductVisibility/ProductVisibilityEntity.php
type ProductVisibility struct {
	Entity
	Visibility     int           `json:"visibility"`
	ProductID      string        `json:"productId"`
	SalesChannelID string        `json:"salesChannelId"`
	Product        *Product      `json:"product"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
