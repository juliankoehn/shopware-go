package shopware

// OrderDelivery see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/Aggregate/OrderDelivery/OrderDeliveryEntity.php
type OrderDelivery struct {
	Entity
	OrderID                string                   `json:"orderId"`
	ShippingOrderAddressID string                   `json:"shippingOrderAddressId"`
	ShippingMethodID       string                   `json:"shippingMethodId"`
	TrackingCodes          []string                 `json:"trackingCodes"`
	ShippingDateEarliest   string                   `json:"shippingDateEarliest"`
	ShippingDateLatest     string                   `json:"shippingDateLatest"`
	ShippingCosts          *CalculatedPrice         `json:"shippingCosts"`
	ShippingOrderAddress   *OrderAddress            `json:"shippingOrderAddress"`
	StateID                string                   `json:"stateId"`
	StateMachineState      *StateMachineState       `json:"stateMachineState"`
	ShippingMethod         *ShippingMethod          `json:"shippingMethod"`
	Order                  *Order                   `json:"order"`
	Positions              []*OrderDeliveryPosition `json:"positions"`
}
