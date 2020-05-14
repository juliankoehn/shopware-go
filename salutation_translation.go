package shopware

// SalutationTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Salutation/Aggregate/SalutationTranslation/SalutationTranslationEntity.php
type SalutationTranslation struct {
	TranslationEntity
	SalutationID string      `json:"salutationId"`
	DisplayName  string      `json:"displayName"`
	LetterName   string      `json:"letterName"`
	Salutation   *Salutation `json:"salutation"`
}
