package shopware

// SalesChannelTypeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/Aggregate/SalesChannelTypeTranslation/SalesChannelTypeTranslationEntity.php
type SalesChannelTypeTranslation struct {
	TranslationEntity
	SalesChannelTypeID string            `json:"salesChannelTypeId"`
	Name               string            `json:"name"`
	Manufacturer       string            `json:"manufacturer"`
	Description        string            `json:"description"`
	DescriptionLong    string            `json:"descriptionLong"`
	SalesChannelType   *SalesChannelType `json:"salesChannelType"`
}
