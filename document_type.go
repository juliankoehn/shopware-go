package shopware

// DocumentType see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Document/Aggregate/DocumentType/DocumentTypeEntity.php
type DocumentType struct {
	Entity
	Name                            string                            `json:"name"`
	TechnicalName                   string                            `json:"technicalName"`
	Translations                    []*DocumentTypeTranslation        `json:"translations"`
	Documents                       []*Document                       `json:"documents"`
	DocumentBaseConfigs             []*DocumentBaseConfig             `json:"documentBaseConfigs"`
	DocumentBaseConfigSalesChannels []*DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels"`
}
