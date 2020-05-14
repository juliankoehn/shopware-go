package shopware

// Salutation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Salutation/SalutationEntity.php
type Salutation struct {
	Entity
	SalutationKey        string                   `json:"salutationKey"`
	DisplayName          string                   `json:"displayName"`
	ÖetterName           string                   `json:"letterName"`
	Translations         []*SalutationTranslation `json:"translations"`
	Customers            []*Customer              `json:"customers"`
	CustomerAddresses    []*CustomerAddress       `json:"customerAddresses"`
	OrderCustomers       []*OrderCustomer         `json:"orderCustomers"`
	OrderAddresses       []*OrderAddress          `json:"orderAddresses"`
	NewsletterRecipients []*NewsletterRecipient   `json:"newsletterRecipients"`
}
