package shopware

// https://github.com/shopware/platform/tree/6.2/src/Core/Framework/DataAbstractionLayer/Pricing

// Price see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/Pricing/Price.php
type Price struct {
	CurrencyID string      `json:"currencyId"`
	Net        float64     `json:"net"`
	Gross      float64     `json:"gross"`
	Linked     bool        `json:"linked"`
	ListPrice  interface{} `json:"listPrice"`
}

// ListingPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/Pricing/ListingPrice.php
type ListingPrice struct {
	CurrencyID string `json:"currencyId"`
	RuleID     string `json:"ruleId"`
	From       Price  `json:"from"`
	To         Price  `json:"to"`
}

// PriceRule see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/Pricing/PriceRuleEntity.php
type PriceRule struct {
	Entity
	RuleID string  `json:"ruleId"`
	Price  []Price `json:"price"`
}
