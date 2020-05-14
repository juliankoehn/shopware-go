package shopware

// LocaleTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Locale/Aggregate/LocaleTranslation/LocaleTranslationEntity.php
type LocaleTranslation struct {
	TranslationEntity
	LocaleID  string  `json:"localeId"`
	Name      string  `json:"name"`
	Territory string  `json:"territory"`
	Locale    *Locale `json:"locale"`
}
