package shopware

// PromotionSalesChannel see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionSalesChannel/PromotionSalesChannelEntity.php
type PromotionSalesChannel struct {
	Entity
	PromotionID    string        `json:"promotionId"`
	SalesChannelID string        `json:"salesChannelId"`
	Priority       int           `json:"priority"`
	Promotion      *Promotion    `json:"promotion"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
