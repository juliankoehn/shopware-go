package shopware

// NumberRangeState see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeState/NumberRangeStateEntity.php
type NumberRangeState struct {
	Entity
	NumberRangeID string       `json:"numberRangeId"`
	LastValue     int          `json:"lastValue"`
	NumberRange   *NumberRange `json:"numberRange"`
}
