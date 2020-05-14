package shopware

// ProductKeywordDictionary see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductKeywordDictionary/ProductKeywordDictionaryEntity.php
type ProductKeywordDictionary struct {
	Entity
	ID         string    `json:"id"`
	LanguageID string    `json:"languageId"`
	Keyword    string    `json:"keyword"`
	Reversed   string    `json:"reversed"`
	Language   *Language `json:"language"`
}
