package shopware

// SalesChannelDomain see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/Aggregate/SalesChannelDomain/SalesChannelDomainEntity.php
type SalesChannelDomain struct {
	Entity
	URL                         string           `json:"url"`
	CurrencyID                  string           `json:"currencyId"`
	currency                    *Currency        `json:"currency"`
	SnippetSetID                string           `json:"snippetSetId"`
	SnippetSet                  *SnippetSet      `json:"snippetSet"`
	SalesChannelID              string           `json:"salesChannelId"`
	SalesChannel                *SalesChannel    `json:"salesChannel"`
	LanguageID                  string           `json:"languageId"`
	Language                    *Language        `json:"language"`
	ProductExports              []*ProductExport `json:"productExports"`
	SalesChannelDefaultHreflang *SalesChannel    `json:"salesChannelDefaultHreflang"`
	HreflangUseOnlyLocale       bool             `json:"hreflangUseOnlyLocale"`
}
