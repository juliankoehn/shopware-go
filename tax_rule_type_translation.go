package shopware

// TaxRuleTypeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Tax/Aggregate/TaxRuleTypeTranslation/TaxRuleTypeTranslationEntity.php
type TaxRuleTypeTranslation struct {
	TranslationEntity
	TaxRuleTypeID string       `json:"taxRuleTypeId"`
	TypeName      string       `json:"typeName"`
	TaxRuleType   *TaxRuleType `json:"taxRuleType"`
}
