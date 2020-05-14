package shopware

// TaxRuleType see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Tax/Aggregate/TaxRuleType/TaxRuleTypeEntity.php
type TaxRuleType struct {
	Entity
	typeName      string                    `json:"typeName"`
	technicalName string                    `json:"technicalName"`
	position      int                       `json:"position"`
	rules         []*TaxRule                `json:"rules"`
	translations  []*TaxRuleTypeTranslation `json:"translations"`
}
