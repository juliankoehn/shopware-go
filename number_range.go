package shopware

// NumberRange see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/NumberRangeEntity.php
type NumberRange struct {
	Entity
	TypeID                   string                     `json:"typeId"`
	Global                   bool                       `json:"global"`
	Name                     string                     `json:"name"`
	Description              string                     `json:"description"`
	Pattern                  string                     `json:"pattern"`
	Start                    int                        `json:"start"`
	NumberRangeType          *NumberRangeType           `json:"type"`
	NumberRangeSalesChannels []*NumberRangeSalesChannel `json:"numberRangeSalesChannels"`
	State                    *NumberRangeState          `json:"state"`
	translations             []*NumberRangeTranslation  `json:"translations"`
}
