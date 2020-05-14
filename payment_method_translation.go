package shopware

// PaymentMethodTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Payment/Aggregate/PaymentMethodTranslation/PaymentMethodTranslationEntity.php
type PaymentMethodTranslation struct {
	TranslationEntity
	PaymentMethodID string         `json:"paymentMethodId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod"`
}
