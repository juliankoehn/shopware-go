package shopware

// CustomerGroupTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Customer/Aggregate/CustomerGroupTranslation/CustomerGroupTranslationEntity.php
type CustomerGroupTranslation struct {
	TranslationEntity
	CustomerGroupID string         `json:"customerGroupId"`
	Name            string         `json:"name"`
	CustomerGroup   *CustomerGroup `json:"customerGroup"`
}
