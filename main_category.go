package shopware

// MainCategory see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Seo/MainCategory/MainCategoryEntity.php
type MainCategory struct {
	Entity
	SalesChannelID string        `json:"salesChannelId"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
	CategoryID     string        `json:"categoryId"`
	Category       *Category     `json:"category"`
	ProductID      string        `json:"productId"`
	Product        *Product      `json:"product"`
}
