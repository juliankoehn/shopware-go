package shopware

// ProductSearchKeyword see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductSearchKeyword/ProductSearchKeywordEntity.php
type ProductSearchKeyword struct {
	Entity
	LanguageID string    `json:"languageId"`
	ProductID  string    `json:"productId"`
	Keyword    string    `json:"keyword"`
	Ranking    float64   `json:"ranking"`
	Product    *Product  `json:"product"`
	Language   *Language `json:"language"`
}
