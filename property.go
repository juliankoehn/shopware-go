package shopware

// PropertyGroupEntity see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Property/PropertyGroupEntity.php
type PropertyGroupEntity struct {
	Entity
	Name         string                      `json:"name"`
	DisplayType  string                      `json:"displayType"`
	SortingType  string                      `json:"sortingType"`
	Description  string                      `json:"description"`
	Position     int                         `json:"position"`
	Options      []*PropertyGroupOption      `json:"options"`
	Translations []*PropertyGroupTranslation `json:"translations"`
}

// PropertyGroupOption see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Property/Aggregate/PropertyGroupOption/PropertyGroupOptionEntity.php
type PropertyGroupOption struct {
	Entity
	GroupID                     string                            `json:"groupId"`
	Name                        string                            `json:"name"`
	Position                    int                               `json:"position"`
	ColorHexCode                string                            `json:"colorHexCode"`
	MediaID                     string                            `json:"mediaId"`
	Group                       *PropertyGroupEntity              `json:"group"`
	Translations                []*PropertyGroupOptionTranslation `json:"translations"`
	ProductConfiguratorSettings []*ProductConfiguratorSetting     `json:"productConfiguratorSettings"`
	ProductProperties           []*Product                        `json:"productProperties"`
	ProductOptions              []*Product                        `json:"productOptions"`
	Media                       *MediaEntity                      `json:"media"`
}

// PropertyGroupOptionTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Property/Aggregate/PropertyGroupOptionTranslation/PropertyGroupOptionTranslationEntity.php
type PropertyGroupOptionTranslation struct {
	TranslationEntity
	PropertyGroupOptionID string               `json:"propertyGroupOptionId"`
	Name                  string               `json:"name"`
	Position              int                  `json:"position"`
	PropertyGroupOption   *PropertyGroupOption `json:"propertyGroupOption"`
}

// PropertyGroupTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Property/Aggregate/PropertyGroupTranslation/PropertyGroupTranslationEntity.php
type PropertyGroupTranslation struct {
	TranslationEntity
	PropertyGroupID string               `json:"propertyGroupId"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Position        int                  `json:"position"`
	PropertyGroup   *PropertyGroupEntity `json:"propertyGroup"`
}
