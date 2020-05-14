package shopware

// ShippingMethodTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Shipping/Aggregate/ShippingMethodTranslation/ShippingMethodTranslationEntity.php
type ShippingMethodTranslation struct {
	TranslationEntity
	ShippingMethodID string          `json:"shippingMethodId"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	TrackingURL      string          `json:"trackingUrl"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod"`
}
