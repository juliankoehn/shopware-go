package shopware

// OrderCustomer see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/Aggregate/OrderCustomer/OrderCustomerEntity.php
type OrderCustomer struct {
	Entity
	Email          string      `json:"email"`
	OrderID        string      `json:"orderId"`
	SalutationID   string      `json:"salutationId"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Title          string      `json:"title"`
	Company        string      `json:"company"`
	CustomerNumber string      `json:"customerNumber"`
	CustomerID     string      `json:"customerId"`
	Customer       *Customer   `json:"customer"`
	Salutation     *Salutation `json:"salutation"`
	Order          *Order      `json:"order"`
	RemoteAddress  string      `json:"remoteAddress"`
}
