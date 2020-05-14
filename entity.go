package shopware

// Entity is a shopware base entity
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/Entity.php
type Entity struct {
	UniqueIdentifier string      `json:"_uniqueIdentifier"`
	EntityName       string      `json:"_entityName"`
	VersionID        string      `json:"versionId"`
	Translated       interface{} `json:"translated"`
	CreatedAt        string      `json:"createdAt"`
	UpdatedAt        string      `json:"updatedAt"`
}

// TranslationEntity see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Framework/DataAbstractionLayer/TranslationEntity.php
type TranslationEntity struct {
	Entity
	LanguageID string `json:"languageId"`
}
