package shopware

// PromotionSetGroup see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionSetGroup/PromotionSetGroupEntity.php
type PromotionSetGroup struct {
	Entity
	packagerKey   string                 `json:"packagerKey"`
	sorterKey     string                 `json:"sorterKey"`
	value         float64                `json:"value"`
	promotionID   string                 `json:"promotionId"`
	promotion     *Promotion             `json:"promotion"`
	setGroupRules *PromotionSetGroupRule `json:"setGroupRules"`
}
