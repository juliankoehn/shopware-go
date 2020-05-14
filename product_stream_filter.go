package shopware

// ProductStreamFilter see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/ProductStream/Aggregate/ProductStreamFilter/ProductStreamFilterEntity.php
type ProductStreamFilter struct {
	Entity
	Typ             string                 `json:"type"` // shopware: type
	Field           string                 `json:"field"`
	Operator        string                 `json:"operator"`
	Value           string                 `json:"value"`
	ProductStreamID string                 `json:"productStreamId"`
	ParentID        string                 `json:"parentId"`
	ProductStream   *ProductStream         `json:"productStream"`
	Queries         []*ProductStreamFilter `json:"queries"`
	Parent          *ProductStreamFilter   `json:"parent"`
	Position        int                    `json:"position"`
	Parameters      []string               `json:"parameters"`
}
