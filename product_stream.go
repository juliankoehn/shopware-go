package shopware

// ProductStream see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/ProductStream/ProductStreamEntity.php
type ProductStream struct {
	Entity
	Name                 string                      `json:"name"`
	Description          string                      `json:"description"`
	APIFilter            []string                    `json:"apiFilter"`
	Filters              []*ProductStreamFilter      `json:"filters"`
	Invalid              bool                        `json:"invalid"`
	Translations         []*ProductStreamTranslation `json:"translations"`
	ProductExports       []*ProductExport            `json:"productExports"`
	ProductCrossSellings []*ProductCrossSelling      `json:"productCrossSellings"`
}
