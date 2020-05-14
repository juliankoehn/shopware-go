package shopware

// PromotionTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionTranslation/PromotionTranslationEntity.php
type PromotionTranslation struct {
	TranslationEntity
	PromotionID string     `json:"promotionId"`
	Name        string     `json:"name"`
	Promotion   *Promotion `json:"promotion"`
}
