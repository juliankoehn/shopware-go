package shopware

// DocumentTypeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Document/Aggregate/DocumentTypeTranslation/DocumentTypeTranslationEntity.php
type DocumentTypeTranslation struct {
	TranslationEntity
	DocumentTypeID string        `json:"documentTypeId"`
	DocumentType   *DocumentType `json:"documentType"`
	Name           string        `json:"name"`
}
