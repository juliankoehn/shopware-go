package shopware

// ReferencePriceDefinition see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/ReferencePriceDefinition.php
type ReferencePriceDefinition struct {
	PurchaseUnit  float64 `json:"purchaseUnit"`
	ReferenceUnit float64 `json:"referenceUnit"`
	UnitName      string  `json:"unitName"`
}

// NewReferencePriceDefinition returns a new ReferencePriceDefinition
func NewReferencePriceDefinition(purchaseUnit, referenceUnit float64, unitName string) *ReferencePriceDefinition {
	return &ReferencePriceDefinition{
		PurchaseUnit:  purchaseUnit,
		ReferenceUnit: referenceUnit,
		UnitName:      unitName,
	}
}

// GetPurchaseUnit returns the PurchaseUnit
func (r *ReferencePriceDefinition) GetPurchaseUnit() float64 {
	return r.PurchaseUnit
}

// GetReferenceUnit returns the ReferenceUnit
func (r *ReferencePriceDefinition) GetReferenceUnit() float64 {
	return r.ReferenceUnit
}

// GetUnitName returns the UnitName
func (r *ReferencePriceDefinition) GetUnitName() string {
	return r.UnitName
}

// GetAPIAlias returns the API Alias
func (*ReferencePriceDefinition) GetAPIAlias() string {
	return "cart_price_reference_definition"
}
