package shopware

const (
	TAX_STATE_GROSS = "gross"
	TAX_STATE_NET   = "net"
	TAX_STATE_FREE  = "tax-free"
)

// CartPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/CartPrice.php
type CartPrice struct {
	NetPrice        float64          `json:"netPrice"`
	TotalPrice      float64          `json:"totalPrice"`
	CalculatedTaxes []*CalculatedTax `json:"calculatedTaxes"`
	TaxRules        []*TaxRule       `json:"taxRules"`
	PositionPrice   float64          `json:"positionPrice"`
	TaxStatus       string           `json:"taxStatus"`
}
