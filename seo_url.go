package shopware

// SeoUrl see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Seo/SeoUrl/SeoUrlEntity.php
type SeoUrl struct {
	Entity
	SalesChannelID string        `json:"salesChannelId"`
	LanguageID     string        `json:"languageId"`
	RouteName      string        `json:"routeName"`
	ForeignKey     string        `json:"foreignKey"`
	PathInfo       string        `json:"pathInfo"`
	SeoPathInfo    string        `json:"seoPathInfo"`
	IsCanonical    bool          `json:"isCanonical"`
	IsModified     bool          `json:"isModified"`
	IsDeleted      bool          `json:"isDeleted"`
	IsValid        bool          `json:"isValid"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
	Language       *Language     `json:"language"`
	URL            string        `json:"url"`
	Error          string        `json:"error"`
}
