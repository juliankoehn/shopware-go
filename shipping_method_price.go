package shopware

// ShippingMethodPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Shipping/Aggregate/ShippingMethodPrice/ShippingMethodPriceEntity.php
type ShippingMethodPrice struct {
	Entity
	ShippingMethodID  string          `json:"shippingMethodId"`
	CurrencyID        string          `json:"currencyId"`
	RuleID            string          `json:"ruleId"`
	Calculation       int             `json:"calculation"`
	QuantityStart     float64         `json:"quantityStart"`
	QuantityEnd       float64         `json:"quantityEnd"`
	Price             float64         `json:"price"`
	ShippingMethod    *ShippingMethod `json:"shippingMethod"`
	Rule              *Rule           `json:"rule"`
	Currency          *Currency       `json:"currency"`
	CalculationRuleID string          `json:"calculationRuleId"`
	CalculationRule   *Rule           `json:"calculationRule"`
	CurrencyPrice     []*Price        `json:"currencyPrice"`
}
