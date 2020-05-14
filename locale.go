package shopware

// Locale see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Locale/LocaleEntity.php
type Locale struct {
	Entity
	Code         string               `json:"code"`
	Name         string               `json:"name"`
	Territory    string               `json:"territory"`
	Translations []*LocaleTranslation `json:"translations"`
	Users        []*User              `json:"users"`
	Languages    []*Language          `json:"languages"`
}
