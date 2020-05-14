package shopware

// OrderAddress see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/Aggregate/OrderAddress/OrderAddressEntity.php
type OrderAddress struct {
	Entity
	CountryID              string           `json:"countryId"`
	CountryStateID         string           `json:"countryStateId"`
	SalutationID           string           `json:"salutationId"`
	FirstName              string           `json:"firstName"`
	LastName               string           `json:"lastName"`
	Street                 string           `json:"street"`
	Zipcode                string           `json:"zipcode"`
	City                   string           `json:"city"`
	Company                string           `json:"company"`
	Department             string           `json:"department"`
	Title                  string           `json:"title"`
	VatID                  string           `json:"vatId"`
	PhoneNumber            string           `json:"phoneNumber"`
	AdditionalAddressLine1 string           `json:"additionalAddressLine1"`
	AdditionalAddressLine2 string           `json:"additionalAddressLine2"`
	Country                *Country         `json:"country"`
	CountryState           *CountryState    `json:"countryState"`
	Order                  *Order           `json:"order"`
	Salutation             *Salutation      `json:"salutation"`
	OrderDeliveries        []*OrderDelivery `json:"orderDeliveries"`
	OrderID                string           `json:"orderId"`
}
