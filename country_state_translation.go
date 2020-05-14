package shopware

// CountryStateTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Country/Aggregate/CountryStateTranslation/CountryStateTranslationEntity.php
type CountryStateTranslation struct {
	TranslationEntity
	CountryStateID string        `json:"countryStateId"`
	Name           string        `json:"name"`
	CountryState   *CountryState `json:"countryState"`
}
