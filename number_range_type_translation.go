package shopware

// NumberRangeTypeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeTypeTranslation/NumberRangeTypeTranslationEntity.php
type NumberRangeTypeTranslation struct {
	TranslationEntity
	NumberRangeTypeId string           `json:"numberRangeTypeId"`
	TypeName          string           `json:"typeName"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType"`
}
