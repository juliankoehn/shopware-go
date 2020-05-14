package shopware

// SeoURLTemplate see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Seo/SeoUrlTemplate/SeoUrlTemplateEntity.php
type SeoURLTemplate struct {
	Entity
	SalesChannelID string        `json:"salesChannelId"`
	EntityName     string        `json:"entityName"`
	RouteName      string        `json:"routeName"`
	Template       string        `json:"template"`
	IsValid        string        `json:"isValid"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
