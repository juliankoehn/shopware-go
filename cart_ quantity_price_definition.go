package shopware

// QuantityPriceDefinition see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/QuantityPriceDefinition.php
type QuantityPriceDefinition struct {
	Price                    float64                   `json:"price"`
	TaxRules                 []*TaxRule                `json:"taxRules"`
	Quantity                 int                       `json:"quantity"`
	isCalculated             bool                      `json:"isCalculated"`
	Precision                int                       `json:"precision"`
	ReferencePriceDefinition *ReferencePriceDefinition `json:"referencePriceDefinition"`
	ListPrice                float64                   `json:"listPrice"`
}

// NewQuantityPriceDefinition creates a new QuantityPriceDefinition
func NewQuantityPriceDefinition(price, listPrice float64, taxRules []*TaxRule, quantity, precision int, isCalculated bool, referencePrice *ReferencePriceDefinition) *QuantityPriceDefinition {
	return &QuantityPriceDefinition{
		Price:                    price,
		ListPrice:                listPrice,
		TaxRules:                 taxRules,
		Quantity:                 quantity,
		isCalculated:             isCalculated,
		Precision:                precision,
		ReferencePriceDefinition: referencePrice,
	}
}

// GetPrice returns the Price
func (q *QuantityPriceDefinition) GetPrice() float64 {
	return q.Price
}

// GetTaxRules returns the TaxRules
func (q *QuantityPriceDefinition) GetTaxRules() []*TaxRule {
	return q.TaxRules
}

// GetQuantity returns the Quantity
func (q *QuantityPriceDefinition) GetQuantity() int {
	return q.Quantity
}

// SetQuantity sets the Quantity
func (q *QuantityPriceDefinition) SetQuantity(quantity int) {
	q.Quantity = quantity
}

// IsCalculated returns isCalculated
func (q *QuantityPriceDefinition) IsCalculated() bool {
	return q.isCalculated
}

// GetPrecision returns the precision
func (q *QuantityPriceDefinition) GetPrecision() int {
	return q.Precision
}

// SetPrecision sets the Precision
func (q *QuantityPriceDefinition) SetPrecision(precision int) {
	q.Precision = precision
}

// GetType returns the type
func (q *QuantityPriceDefinition) GetType() string {
	return "quantity"
}

// GetPriority returns the priority
func (q *QuantityPriceDefinition) GetPriority() int {
	return 100
}

// GetConstraints returns the constraints
func (q *QuantityPriceDefinition) GetConstraints() *Constraint {
	return &Constraint{
		Price:        0,
		Quantity:     0,
		IsCalculated: false,
		Precision:    0,
	}
}

// Verify that PriceDefinitionInterface implements QuantityPriceDefinition.
var _ PriceDefinitionInterface = QuantityPriceDefinition{}
