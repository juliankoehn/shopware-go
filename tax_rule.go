package shopware

// TaxRule see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Tax/Aggregate/TaxRule/TaxRuleEntity.php
type TaxRule struct {
	Entity
	TaxID         string       `json:"taxId"`
	Tax           *Tax         `json:"tax"`
	CountryID     string       `json:"countryId"`
	Country       *Country     `json:"country"`
	TaxRuleTypeID string       `json:"taxRuleTypeId"`
	TaxRuleType   *TaxRuleType `json:"type"`
	TaxRate       float64      `json:"taxRate"`
	Data          []string     `json:"data"`
}
