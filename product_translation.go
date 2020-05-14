package shopware

// ProductTranslation holds translations for a Product
type ProductTranslation struct {
	TranslationEntity
	Keywords        string   `json:"keywords"`
	MetaDescription string   `json:"metaDescription"`
	MetaTitle       string   `json:"metaTitle"`
	Name            string   `json:"name"`
	PackUnit        string   `json:"packUnit"`
	ProductID       string   `json:"productId"`
	Product         *Product `json:"product"`
}
