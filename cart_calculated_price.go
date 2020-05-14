package shopware

// CalculatedPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/CalculatedPrice.php
type CalculatedPrice struct {
	UnitPrice       float64          `json:"unitPrice"`
	Quantity        int              `json:"quantity"`
	TotalPrice      float64          `json:"totalPrice"`
	CalculatedTaxes []*CalculatedTax `json:"calculatedTaxes"`
	TaxRules        []*TaxRule       `json:"taxRules"`
	ReferencePrice  *ReferencePrice  `json:"referencePrice"`
	ListPrice       *ListPrice       `json:"listPrice"`
}
