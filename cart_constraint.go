package shopware

// Constraint is used often - kek!
type Constraint struct {
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	IsCalculated bool    `json:"isCalculated"`
	Precision    int     `json:"precision"`
}
