package shopware

// CurrencyTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Currency/Aggregate/CurrencyTranslation/CurrencyTranslationEntity.php
type CurrencyTranslation struct {
	TranslationEntity
	CurrencyId string    `json:"currencyId"`
	ShortName  string    `json:"shortName"`
	Name       string    `json:"name"`
	Currency   *Currency `json:"currency"`
}
