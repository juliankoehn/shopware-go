package shopware

// CustomerAddress see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Customer/Aggregate/CustomerAddress/CustomerAddressEntity.php
type CustomerAddress struct {
	Entity
	CustomerID             string        `json:"customerId"`
	CountryID              string        `json:"countryId"`
	CountryStateID         string        `json:"countryStateId"`
	SalutationID           string        `json:"salutationId"`
	FirstName              string        `json:"firstName"`
	LastName               string        `json:"lastName"`
	Zipcode                string        `json:"zipcode"`
	City                   string        `json:"city"`
	Company                string        `json:"company"`
	Department             string        `json:"department"`
	Title                  string        `json:"title"`
	Street                 string        `json:"street"`
	VatID                  string        `json:"vatId"`
	PhoneNumber            string        `json:"phoneNumber"`
	AdditionalAddressLine1 string        `json:"additionalAddressLine1"`
	AdditionalAddressLine2 string        `json:"additionalAddressLine2"`
	Country                *Country      `json:"country"`
	CountryState           *CountryState `json:"countryState"`
	Salutation             *Salutation   `json:"salutation"`
	Customer               *Customer     `json:"customer"`
}
