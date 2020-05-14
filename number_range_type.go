package shopware

// NumberRangeType see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeType/NumberRangeTypeEntity.php
type NumberRangeType struct {
	Entity
	TypeName                 string                        `json:"typeName"`
	TechnicalName            string                        `json:"technicalName"`
	Global                   bool                          `json:"global"`
	NumberRanges             []*NumberRange                `json:"numberRanges"`
	NumberRangeSalesChannels *NumberRangeSalesChannel      `json:"numberRangeSalesChannels"`
	translations             []*NumberRangeTypeTranslation `json:"translations"`
}
