package shopware

// Currency see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Currency/CurrencyEntity.php
type Currency struct {
	Entity
	IsoCode                        string                    `json:"isoCode"`
	Factor                         float64                   `json:"factor"`
	Symbol                         string                    `json:"symbol"`
	ShortName                      string                    `json:"shortName"`
	Name                           string                    `json:"name"`
	Position                       string                    `json:"position"`
	DecimalPrecision               string                    `json:"decimalPrecision"`
	Translations                   []*CurrencyTranslation    `json:"translations"`
	Orders                         []*Order                  `json:"orders"`
	SalesChannels                  []*SalesChannel           `json:"salesChannels"`
	SalesChannelDefaultAssignments []*SalesChannel           `json:"salesChannelDefaultAssignments"`
	SalesChannelDomains            []*SalesChannelDomain     `json:"salesChannelDomains"`
	ShippingMethodPrices           []*ShippingMethodPrice    `json:"shippingMethodPrices"`
	PromotionDiscountPrices        []*PromotionDiscountPrice `json:"promotionDiscountPrices"`
	IsSystemDefault                bool                      `json:"isSystemDefault"`
	ProductExports                 []*ProductExport          `json:"productExports"`
}
