package shopware

const DELIVERY_TIME_DAY = "day"
const DELIVERY_TIME_WEEK = "week"
const DELIVERY_TIME_MONTH = "month"

// DeliveryTime see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/DeliveryTime/DeliveryTimeEntity.php
type DeliveryTime struct {
	Entity
	Name            string                     `json:"name"`
	Min             int                        `json:"min"`
	Max             int                        `json:"max"`
	Unit            string                     `json:"unit"`
	Translations    []*DeliveryTimeTranslation `json:"translations"`
	ShippingMethods []*ShippingMethod          `json:"shippingMethods"`
	Products        []*Product                 `json:"products"`
}
