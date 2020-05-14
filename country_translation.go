package shopware

// CountryTranslation see:
// https://github.com/shopware/platform/tree/6.2/src/Core/System/Country/Aggregate/CountryTranslation
type CountryTranslation struct {
	TranslationEntity
	CountryID string   `json:"countryId"`
	Name      string   `json:"name"`
	Country   *Country `json:"country"`
}
