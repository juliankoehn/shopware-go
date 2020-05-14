package shopware

// ReferencePrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/ReferencePrice.php
type ReferencePrice struct {
	ReferencePriceDefinition
	Price float64 `json:"price"`
}

// NewReferencePrice creates a new NewReferencePrice
func NewReferencePrice(price, purchaseUnit, referenceUnit float64, unitName string) *ReferencePrice {
	return &ReferencePrice{
		Price: price,
		ReferencePriceDefinition: ReferencePriceDefinition{
			PurchaseUnit:  purchaseUnit,
			ReferenceUnit: referenceUnit,
			UnitName:      unitName,
		},
	}
}

// GetPrice returns the Price
func (r *ReferencePrice) GetPrice() float64 {
	return r.Price
}

// GetAPIAlias returns the API Alias
func (r *ReferencePrice) GetAPIAlias() string {
	return "cart_price_reference"
}
