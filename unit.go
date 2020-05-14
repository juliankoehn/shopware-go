package shopware

// Unit see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Unit/UnitEntity.php
type Unit struct {
	Entity
	shortCode    string             `json:"shortCode"`
	name         string             `json:"name"`
	translations []*UnitTranslation `json:"translations"`
	products     []*Product         `json:"products"`
}
