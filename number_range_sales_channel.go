package shopware

// NumberRangeSalesChannel see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeSalesChannel/NumberRangeSalesChannelEntity.php
type NumberRangeSalesChannel struct {
	Entity
	NumberRangeID     string           `json:"numberRangeId"`
	SalesChannelID    string           `json:"salesChannelId"`
	NumberRangeTypeID string           `json:"numberRangeTypeId"`
	NumberRange       *NumberRange     `json:"numberRange"`
	SalesChannel      *SalesChannel    `json:"salesChannel"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType"`
}
