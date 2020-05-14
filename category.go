package shopware

// Category see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Category/CategoryEntity.php
type Category struct {
	Entity
	ParentID                string                 `json:"parentId"`
	AutoIncrement           int                    `json:"autoIncrement"`
	MediaID                 string                 `json:"mediaId"`
	Name                    string                 `json:"name"`
	Breadcrumb              []string               `json:"breadcrumb"`
	Path                    string                 `json:"path"`
	Level                   int                    `json:"level"`
	Active                  bool                   `json:"active"`
	ChildCount              int                    `json:"childCount"`
	displayNestedProducts   bool                   `json:"displayNestedProducts"`
	parent                  *Category              `json:"parent"`
	children                []*Category            `json:"children"`
	translations            []*CategoryTranslation `json:"translations"`
	Media                   *MediaEntity           `json:"media"`
	Products                []*Product             `json:"products"`
	NestedProducts          []*Product             `json:"nestedProducts"`
	AfterCategoryID         string                 `json:"afterCategoryId"`
	Tags                    []*Tag                 `json:"tags"`
	CmsPageID               string                 `json:"cmsPageId"`
	CmsPage                 *CmsPage               `json:"cmsPage"`
	SlotConfig              []string               `json:"slotConfig"`
	NavigationSalesChannels []*SalesChannel        `json:"navigationSalesChannels"`
	FooterSalesChannels     []*SalesChannel        `json:"footerSalesChannels"`
	ServiceSalesChannels    []*SalesChannel        `json:"serviceSalesChannels"`
	ExternalLink            string                 `json:"externalLink"`
	Visible                 bool                   `json:"visible"`
	Typ                     string                 `json:"type"` // shopware: type
	Description             string                 `json:"description"`
	MetaTitle               string                 `json:"metaTitle"`
	MetaDescription         string                 `json:"metaDescription"`
	Keywords                string                 `json:"keywords"`
	MainCategories          []*MainCategory        `json:"mainCategories"`
	SeoUrls                 []*SeoUrl              `json:"seoUrls"`
}
