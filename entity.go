package shopware

// Entity is a shopware base entity
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/Entity.php
type Entity struct {
	ID               string      `json:"id"`
	UniqueIdentifier string      `json:"_uniqueIdentifier"`
	EntityName       string      `json:"_entityName"`
	VersionID        string      `json:"versionId"`
	Translated       interface{} `json:"translated"`
	CreatedAt        string      `json:"createdAt"`
	UpdatedAt        string      `json:"updatedAt"`
}

// GetID returns the ID of entity
func (e *Entity) GetID() string {
	return e.ID
}

// SetID sets the ID of entity
func (e *Entity) SetID(id string) {
	e.ID = id
	e.UniqueIdentifier = id
}

// TranslationEntity see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/TranslationEntity.php
type TranslationEntity struct {
	Entity
	LanguageID string `json:"languageId"`
}
