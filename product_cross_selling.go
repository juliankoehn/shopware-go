package shopware

// ProductCrossSelling see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductCrossSelling/ProductCrossSellingEntity.php
type ProductCrossSelling struct {
	Entity
	Name             string                                `json:"name"`
	Position         int                                   `json:"position"`
	SortBy           string                                `json:"sortBy"`
	SortDirection    string                                `json:"sortDirection"`
	Limit            int                                   `json:"limit"`
	Active           bool                                  `json:"active"`
	ProductID        string                                `json:"productId"`
	Product          *Product                              `json:"product"`
	ProductStreamID  string                                `json:"productStreamId"`
	ProductStream    ProductStream                         `json:"productStream"`
	Typ              string                                `json:"type"` // real: type
	AssignedProducts []*ProductCrossSellingAssignedProduct `json:"assignedProducts"`
	Translations     []*ProductCrossSellingTranslation     `json:"translations"`
}

type ProductStream struct{}
