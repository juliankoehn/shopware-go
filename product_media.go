package shopware

// ProductMediaEntity see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductMedia/ProductMediaEntity.php
type ProductMediaEntity struct {
	Entity
	Media     *MediaEntity `json:"media"`
	MediaID   string       `json:"mediaId"`
	Position  int          `json:"position"`
	Product   *Product     `json:"product"`
	ProductID string       `json:"productId"`
}
