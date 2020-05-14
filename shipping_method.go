package shopware

// ShippingMethod see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Shipping/ShippingMethodEntity.php
type ShippingMethod struct {
	Entity
	Name                           string                       `json:"name"`
	Active                         bool                         `json:"active"`
	Description                    string                       `json:"description"`
	TrackingURL                    string                       `json:"trackingUrl"`
	DeliveryTimeID                 string                       `json:"deliveryTimeId"`
	DeliveryTime                   *DeliveryTime                `json:"deliveryTime"`
	Translations                   []*ShippingMethodTranslation `json:"translations"`
	OrderDeliveries                []*OrderDelivery             `json:"orderDeliveries"`
	SalesChannelDefaultAssignments []*SalesChannel              `json:"salesChannelDefaultAssignments"`
	SalesChannels                  []*SalesChannel              `json:"salesChannels"`
	AvailabilityRule               *Rule                        `json:"availabilityRule"`
	AvailabilityRuleID             string                       `json:"availabilityRuleId"`
	Prices                         []*ShippingMethodPrice       `json:"prices"`
	MediaID                        string                       `json:"mediaId"`
	Media                          *MediaEntity                 `json:"media"`
	Tags                           []*Tag                       `json:"tags"`
}
