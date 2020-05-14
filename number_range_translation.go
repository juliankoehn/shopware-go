package shopware

// NumberRangeTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/NumberRange/Aggregate/NumberRangeTranslation/NumberRangeTranslationEntity.php
type NumberRangeTranslation struct {
	TranslationEntity
	NumberRangeID string       `json:"numberRangeId"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	NumberRange   *NumberRange `json:"numberRange"`
}
