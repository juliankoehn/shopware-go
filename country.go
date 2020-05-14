package shopware

// Country see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Country/CountryEntity.php
type Country struct {
	Entity
	Name                           string                `json:"name"`
	Iso                            string                `json:"iso"`
	Position                       int                   `json:"position"`
	TaxFree                        bool                  `json:"taxFree"`
	Active                         bool                  `json:"active"`
	ShippingAvailable              bool                  `json:"shippingAvailable"`
	Iso3                           string                `json:"iso3"`
	DisplayStateInRegistration     bool                  `json:"displayStateInRegistration"`
	ForceStateInRegistration       bool                  `json:"forceStateInRegistration"`
	States                         []*CountryState       `json:"states"`
	Translations                   []*CountryTranslation `json:"translations"`
	OrderAddresses                 []*OrderAddress       `json:"orderAddresses"`
	CustomerAddresses              []*CustomerAddress    `json:"customerAddresses"`
	SalesChannelDefaultAssignments []*SalesChannel       `json:"salesChannelDefaultAssignments"`
	SalesChannels                  []*SalesChannel       `json:"salesChannels"`
	TaxRules                       []*TaxRule            `json:"taxRules"`
}
