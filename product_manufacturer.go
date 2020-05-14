package shopware

// ProductManufacturer is a manufacturer of an product
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductManufacturer/ProductManufacturerEntity.php
type ProductManufacturer struct {
	Entity
	Description  string                            `json:"description"`
	Link         string                            `json:"link"`
	Media        MediaEntity                       `json:"media"`
	MediaID      string                            `json:"mediaId"`
	Name         string                            `json:"name"`
	Products     []*Product                        `json:"products"`
	Translations []*ProductManufacturerTranslation `json:"translations"`
}

// ProductManufacturerTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductManufacturerTranslation/ProductManufacturerTranslationEntity.php
type ProductManufacturerTranslation struct {
	TranslationEntity
	ProductManufacturerID string               `json:"productManufacturerId"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	ProductManufacturer   *ProductManufacturer `json:"productManufacturer"`
}
