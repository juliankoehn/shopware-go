package shopware

// SalesChannel see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/SalesChannelEntity.php
type SalesChannel struct {
	Entity
	TypeID                          string                       `json:"typeId"`
	LanguageID                      string                       `json:"languageId"`
	CurrencyID                      string                       `json:"currencyId"`
	PaymentMethodID                 string                       `json:"paymentMethodId"`
	ShippingMethodID                string                       `json:"shippingMethodId"`
	CountryID                       string                       `json:"countryId"`
	NavigationCategoryID            string                       `json:"navigationCategoryId"`
	NavigationCategoryDepth         int                          `json:"navigationCategoryDepth"`
	FooterCategoryID                string                       `json:"footerCategoryId"`
	ServiceCategoryID               string                       `json:"serviceCategoryId"`
	Name                            string                       `json:"name"`
	ShortName                       string                       `json:"shortName"`
	AccessKey                       string                       `json:"accessKey"`
	Currencies                      []*Currency                  `json:"currencies"`
	Languages                       []*Language                  `json:"languages"`
	Configuration                   []string                     `json:"configuration"`
	Active                          bool                         `json:"active"`
	Maintenance                     bool                         `json:"maintenance"`
	MaintenanceIPWhitelistgo        []string                     `json:"maintenanceIpWhitelist"`
	SalesChannelType                *SalesChannelType            `json:"type"`
	Currency                        *Currency                    `json:"currency"`
	Language                        *Language                    `json:"language"`
	PaymentMethod                   *PaymentMethod               `json:"paymentMethod"`
	ShippingMethod                  *ShippingMethod              `json:"shippingMethod"`
	Country                         *Country                     `json:"country"`
	Orders                          []*Order                     `json:"orders"`
	Customers                       []*Customer                  `json:"customers"`
	Countries                       []*Country                   `json:"countries"`
	PaymentMethods                  []*PaymentMethod             `json:"paymentMethods"`
	ShippingMethods                 []*ShippingMethod            `json:"shippingMethods"`
	Translations                    []*SalesChannelTranslation   `json:"translations"`
	Domains                         []*SalesChannelDomain        `json:"domains"`
	SystemConfigs                   []*SystemConfig              `json:"systemConfigs"`
	NavigationCategory              *Category                    `json:"navigationCategory"`
	FooterCategory                  *Category                    `json:"footerCategory"`
	ServiceCategory                 *Category                    `json:"serviceCategory"`
	ProductVisibilities             []*ProductVisibility         `json:"productVisibilities"`
	MailTemplates                   []*MailTemplateSalesChannel  `json:"mailTemplates"`
	MailHeaderFooterID              string                       `json:"mailHeaderFooterId"`
	NumberRangeSalesChannels        []*NumberRangeSalesChannel   `json:"numberRangeSalesChannels"`
	MailHeaderFooter                string                       `json:"mailHeaderFooter"`
	CustomerGroupID                 string                       `json:"customerGroupId"`
	CustomerGroup                   *CustomerGroup               `json:"customerGroup"`
	NewsletterRecipients            []*NewsletterRecipient       `json:"newsletterRecipients"`
	PromotionSalesChannels          []*PromotionSalesChannel     `json:"promotionSalesChannels"`
	DocumentBaseConfigSalesChannels DocumentBaseConfigDefinition `json:"documentBaseConfigSalesChannels"`
	ProductReviews                  []*ProductReview             `json:"productReviews"`
	SeoUrls                         []*SeoURL                    `json:"seoUrls"`
	SeoURLTemplates                 []*SeoURLTemplate            `json:"seoUrlTemplates"`
	MainCategories                  []*MainCategory              `json:"mainCategories"`
	PaymentMethodIDs                []string                     `json:"paymentMethodIds"`
	ProductExports                  []*ProductExport             `json:"productExports"`
	HreflangActive                  bool                         `json:"hreflangActive"`
	HreflangDefaultDomainID         string                       `json:"hreflangDefaultDomainId"`
	HreflangDefaultDomain           *SalesChannelDomain          `json:"hreflangDefaultDomain"`
	AnalyticsID                     string                       `json:"analyticsId"`
	Analytics                       *SalesChannelAnalytics       `json:"analytics"`
	GoogleShoppingAccount           *GoogleShoppingAccount       `json:"googleShoppingAccount"`
}
