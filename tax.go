package shopware

// Tax see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Tax/TaxEntity.php
type Tax struct {
	Entity
	TaxRate  float64    `json:"taxRate"`
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
	rules    []*TaxRule `json:"rules"`
}
