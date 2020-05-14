package shopware

// OrderDeliveryPosition see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/Aggregate/OrderDeliveryPosition/OrderDeliveryPositionEntity.php
type OrderDeliveryPosition struct {
	Entity
	OrderDeliveryID string           `json:"orderDeliveryId"`
	OrderLineItemID string           `json:"orderLineItemId"`
	Price           *CalculatedPrice `json:"price"`
	UnitPrice       float64          `json:"unitPrice"`
	TotalPrice      float64          `json:"totalPrice"`
	Quantity        int              `json:"quantity"`
	OrderLineItem   *OrderLineItem   `json:"orderLineItem"`
	OrderDelivery   *OrderDelivery   `json:"orderDelivery"`
}
