package shopware

// UnitTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Unit/Aggregate/UnitTranslation/UnitTranslationEntity.php
type UnitTranslation struct {
	TranslationEntity
	UnitID    string `json:"unitId"`
	ShortCode string `json:"shortCode"`
	Name      string `json:"name"`
	Unit      *Unit  `json:"unit"`
}
