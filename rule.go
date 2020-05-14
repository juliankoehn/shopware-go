package shopware

// Rule see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Rule/RuleEntity.php
type Rule struct {
	Entity
	Name                            string                 `json:"name"`
	Description                     string                 `json:"description"`
	Priority                        int                    `json:"priority"`
	Payload                         interface{}            `json:"payload"` // TODO if it's used somewhere within API
	ModuleTypes                     []string               `json:"moduleTypes"`
	ProductPrices                   []*ProductPrice        `json:"productPrices"`
	ShippingMethods                 []*ShippingMethod      `json:"shippingMethods"`
	PaymentMethods                  []*PaymentMethod       `json:"paymentMethods"`
	Conditions                      []*RuleCondition       `json:"conditions"`
	Invalid                         bool                   `json:"invalid"`
	ShippingMethodPrices            []*ShippingMethodPrice `json:"shippingMethodPrices"`
	PromotionDiscounts              []*PromotionDiscount   `json:"promotionDiscounts"`
	PromotionSetGroups              []*PromotionSetGroup   `json:"promotionSetGroups"`
	ShippingMethodPriceCalculations []*ShippingMethodPrice `json:"shippingMethodPriceCalculations"`
	PersonaPromotions               []*Promotion           `json:"personaPromotions"`
	OrderPromotions                 []*Promotion           `json:"orderPromotions"`
	CartPromotions                  []*Promotion           `json:"cartPromotions"`
}
