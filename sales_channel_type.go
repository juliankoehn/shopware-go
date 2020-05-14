package shopware

// SalesChannelType see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/Aggregate/SalesChannelType/SalesChannelTypeEntity.php
type SalesChannelType struct {
	Entity
	Name            string                         `json:"name"`
	Manufacturer    string                         `json:"manufacturer"`
	Description     string                         `json:"description"`
	DescriptionLong string                         `json:"descriptionLong"`
	CoverUrl        string                         `json:"coverUrl"`
	IconName        string                         `json:"iconName"`
	ScreenshotUrls  []string                       `json:"screenshotUrls"`
	SalesChannels   []*SalesChannel                `json:"salesChannels"`
	translations    []*SalesChannelTypeTranslation `json:"translations"`
}
