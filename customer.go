package shopware

const ACCOUNT_TYPE_PRIVATE = "private"
const ACCOUNT_TYPE_BUSINESS = "business"

// Customer see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Customer/CustomerEntity.php
type Customer struct {
	Entity
	GroupID                  string             `json:"groupId"`
	DefaultPaymentMethodID   string             `json:"defaultPaymentMethodId"`
	SalesChannelID           string             `json:"salesChannelId"`
	LanguageID               string             `json:"languageId"`
	LastPaymentMethodID      string             `json:"lastPaymentMethodId"`
	DefaultBillingAddressID  string             `json:"defaultBillingAddressId"`
	DefaultShippingAddressID string             `json:"defaultShippingAddressId"`
	CustomerNumber           string             `json:"customerNumber"`
	SalutationID             string             `json:"salutationId"`
	FirstName                string             `json:"firstName"`
	LastName                 string             `json:"lastName"`
	Company                  string             `json:"company"`
	Password                 string             `json:"password"`
	Email                    string             `json:"email"`
	Title                    string             `json:"title"`
	AffiliateCode            string             `json:"affiliateCode"`
	CampaignCode             string             `json:"campaignCode"`
	Active                   bool               `json:"active"`
	DoubleOptInRegistration  bool               `json:"doubleOptInRegistration"`
	DoubleOptInEmailSentDate string             `json:"doubleOptInEmailSentDate"`
	DoubleOptInConfirmDate   string             `json:"doubleOptInConfirmDate"`
	Hash                     string             `json:"hash"`
	Guest                    bool               `json:"guest"`
	FirstLogin               string             `json:"firstLogin"`
	LastLogin                string             `json:"lastLogin"`
	Newsletter               bool               `json:"newsletter"`
	Birthday                 string             `json:"birthday"`
	LastOrderDate            string             `json:"lastOrderDate"`
	OrderCount               int                `json:"orderCount"`
	CreatedAt                string             `json:"createdAt"`
	UpdatedAt                string             `json:"updatedAt"`
	LegacyEncoder            string             `json:"legacyEncoder"`
	LegacyPassword           string             `json:"legacyPassword"`
	Group                    *CustomerGroup     `json:"group"`
	DefaultPaymentMethod     *PaymentMethod     `json:"defaultPaymentMethod"`
	SalesChannel             *SalesChannel      `json:"salesChannel"`
	Language                 *Language          `json:"language"`
	LastPaymentMethod        *PaymentMethod     `json:"lastPaymentMethod"`
	Salutation               *Salutation        `json:"salutation"`
	DefaultBillingAddress    *CustomerAddress   `json:"defaultBillingAddress"`
	DefaultShippingAddress   *CustomerAddress   `json:"defaultShippingAddress"`
	ActiveBillingAddress     *CustomerAddress   `json:"activeBillingAddress"`
	ActiveShippingAddress    *CustomerAddress   `json:"activeShippingAddress"`
	Addresses                []*CustomerAddress `json:"addresses"`
	OrderCustomers           []*OrderCustomer   `json:"orderCustomers"`
	AutoIncrement            int                `json:"autoIncrement"`
	Tags                     []*Tag             `json:"tags"`
	Promotions               []*Promotion       `json:"promotions"`
	RecoveryCustomer         *CustomerRecovery  `json:"recoveryCustomer"`
	ProductReviews           []*ProductReview   `json:"productReviews"`
	RemoteAddress            string             `json:"remoteAddress"`
}
