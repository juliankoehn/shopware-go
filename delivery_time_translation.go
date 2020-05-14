package shopware

// DeliveryTimeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/DeliveryTime/Aggregate/DeliveryTimeTranslation/DeliveryTimeTranslationEntity.php
type DeliveryTimeTranslation struct {
	TranslationEntity
	DeliveryTime   *DeliveryTime `json:"deliveryTime"`
	DeliveryTimeID string        `json:"deliveryTimeId"`
	Name           string        `json:"name"`
}
