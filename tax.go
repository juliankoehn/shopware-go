package shopware

// TODO
type TaxEntity struct {
	TaxRate          float64    `json:"taxRate"`
	Name             string     `json:"name"`
	Products         []*Product `json:"products"`
	UniqueIdentifier string     `json:"_uniqueIdentifier"`
	CreatedAt        string     `json:"createdAt"`
	UpdatedAt        string     `json:"updatedAt"`
	ID               string     `json:"id"`
}
