package shopware

// Order see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/OrderEntity.php
type Order struct {
	Entity
	OrderNumber       string              `json:"orderNumber"`
	CurrencyID        string              `json:"currencyId"`
	CurrencyFactor    float64             `json:"currencyFactor"`
	SalesChannelID    string              `json:"salesChannelId"`
	BillingAddressID  string              `json:"billingAddressId"`
	OrderDateTime     string              `json:"orderDateTime"`
	OrderDate         string              `json:"orderDate"`
	Price             *CartPrice          `json:"price"`
	AmountTotal       float64             `json:"amountTotal"`
	AmountNet         float64             `json:"amountNet"`
	PositionPrice     float64             `json:"positionPrice"`
	TaxStatus         string              `json:"taxStatus"`
	ShippingCosts     *CalculatedPrice    `json:"shippingCosts"`
	ShippingTotal     float64             `json:"shippingTotal"`
	OrderCustomer     *Customer           `json:"orderCustomer"`
	Currency          *Currency           `json:"currency"`
	LanguageID        string              `json:"languageId"`
	Language          *Language           `json:"language"`
	SalesChannel      *SalesChannel       `json:"salesChannel"`
	Addresses         []*OrderAddress     `json:"addresses"`
	Deliveries        []*OrderDelivery    `json:"deliveries"`
	LineItems         []*OrderLineItem    `json:"lineItems"`
	Transactions      []*OrderTransaction `json:"transactions"`
	DeepLinkCode      string              `json:"deepLinkCode"`
	AutoIncrement     int                 `json:"autoIncrement"`
	StateMachineState *StateMachineState  `json:"stateMachineState"`
	StateID           string              `json:"stateId"`
	Documents         []*Document         `json:"documents"`
	Tags              []*Tag              `json:"tags"`
	AffiliateCode     string              `json:"affiliateCode"`
	CampaignCode      string              `json:"campaignCode"`
	CustomerComment   string              `json:"customerComment"`
}
