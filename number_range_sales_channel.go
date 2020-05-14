package shopware

// NumberRangeSalesChannel see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeSalesChannel/NumberRangeSalesChannelEntity.php
type NumberRangeSalesChannel struct {
	Entity
	NumberRangeId     string           `json:"numberRangeId"`
	SalesChannelId    string           `json:"salesChannelId"`
	NumberRangeTypeId string           `json:"numberRangeTypeId"`
	NumberRange       *NumberRange     `json:"numberRange"`
	SalesChannel      *SalesChannel    `json:"salesChannel"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType"`
}
