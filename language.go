package shopware

// Language see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Language/LanguageEntity.php
type Language struct {
	Entity
	ParentID                        string                            `json:"parentId"`
	LocaleID                        string                            `json:"localeId"`
	TranslationCodeID               string                            `json:"translationCodeId"`
	TranslationCode                 *Locale                           `json:"translationCode"`
	Name                            string                            `json:"name"`
	Locale                          *Locale                           `json:"locale"`
	Parent                          *Language                         `json:"parent"`
	Children                        []*Language                       `json:"children"`
	SalesChannels                   []*SalesChannel                   `json:"salesChannels"`
	Customers                       []*Customer                       `json:"customers"`
	SalesChannelDefaultAssignments  []*SalesChannel                   `json:"salesChannelDefaultAssignments"`
	CategoryTranslations            []*CategoryTranslation            `json:"categoryTranslations"`
	CountryStateTranslations        []*CountryStateTranslation        `json:"countryStateTranslations"`
	CountryTranslations             []*CountryTranslation             `json:"countryTranslations"`
	CurrencyTranslations            []*CurrencyTranslation            `json:"currencyTranslations"`
	CustomerGroupTranslations       []*CustomerGroupTranslation       `json:"customerGroupTranslations"`
	LocaleTranslations              []*LocaleTranslation              `json:"localeTranslations"`
	MediaTranslations               []*MediaTranslation               `json:"mediaTranslations"`
	PaymentMethodTranslations       []*PaymentMethodTranslation       `json:"paymentMethodTranslations"`
	ProductManufacturerTranslations []*ProductManufacturerTranslation `json:"productManufacturerTranslations"`
	ProductTranslations             []*ProductTranslation             `json:"productTranslations"`
	ShippingMethodTranslations      []*ShippingMethodTranslation      `json:"shippingMethodTranslations"`
	UnitTranslations                []*UnitTranslation                `json:"unitTranslations"`
	PropertyGroupTranslations       []*PropertyGroupTranslation       `json:"propertyGroupTranslations"`
	PropertyGroupOptionTranslations []*PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations"`
	SalesChannelTranslations        []*SalesChannelTranslation        `json:"salesChannelTranslations"`
	SalesChannelTypeTranslations    []*SalesChannelTypeTranslation    `json:"salesChannelTypeTranslations"`
	SalutationTranslations          []*SalutationTranslation          `json:"salutationTranslations"`
	SalesChannelDomains             []*SalesChannelDomain             `json:"salesChannelDomains"`
	ProductStreamTranslations       []*ProductStreamTranslation       `json:"productStreamTranslations"`
	DeliveryTimeTranslations        []*DeliveryTimeTranslation        `json:"deliveryTimeTranslations"`
	newsletterRecipients            []*NewsletterRecipient            `json:"newsletterRecipients"`
	Orders                          []*Order                          `json:"orders"`
	NumberRangeTypeTranslations     []*NumberRangeTypeTranslation     `json:"numberRangeTypeTranslations"`
	ProductSearchKeywords           []*ProductSearchKeyword           `json:"productSearchKeywords"`
	ProductKeywordDictionaries      []*ProductKeywordDictionary       `json:"productKeywordDictionaries"`
	numberRangeTranslations         []*NumberRangeTranslation         `json:"numberRangeTranslations"`
	ProductReviews                  []*ProductReview                  `json:"productReviews"`
	SeoUrlTranslations              []*SeoUrlTranslation              `json:"seoUrlTranslations"`
	TaxRuleTypeTranslations         []*TaxRuleTypeTranslation         `json:"taxRuleTypeTranslations"`
	ProductCrossSellingTranslations []*ProductCrossSellingTranslation `json:"productCrossSellingTranslations"`
}
