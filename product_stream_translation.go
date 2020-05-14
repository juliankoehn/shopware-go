package shopware

// ProductStreamTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/ProductStream/Aggregate/ProductStreamTranslation/ProductStreamTranslationEntity.php
type ProductStreamTranslation struct {
	TranslationEntity
	ProductStreamID string         `json:"productStreamId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	ProductStream   *ProductStream `json:"productStream"`
}
