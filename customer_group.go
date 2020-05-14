package shopware

// CustomerGroup see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Customer/Aggregate/CustomerGroup/CustomerGroupEntity.php
type CustomerGroup struct {
	Entity
	Name          string                      `json:"name"`
	DisplayGross  bool                        `json:"displayGross"`
	Translations  []*CustomerGroupTranslation `json:"translations"`
	Customers     []*Customer                 `json:"customers"`
	SalesChannels []*SalesChannel             `json:"salesChannels"`
}
