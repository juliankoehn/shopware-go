package shopware

// SalesChannelTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/Aggregate/SalesChannelTranslation/SalesChannelTranslationEntity.php
type SalesChannelTranslation struct {
	TranslationEntity
	SalesChannelID string        `json:"salesChannelId"`
	Name           string        `json:"name"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
