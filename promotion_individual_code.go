package shopware

// PromotionIndividualCode see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionIndividualCode/PromotionIndividualCodeEntity.php
type PromotionIndividualCode struct {
	Entity
	PromotionID string     `json:"promotionId"`
	Code        string     `json:"code"`
	Promotion   *Promotion `json:"promotion"`
	Payload     []string   `json:"payload"`
}
