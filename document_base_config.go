package shopware

// DocumentBaseConfig see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Document/Aggregate/DocumentBaseConfig/DocumentBaseConfigEntity.php
type DocumentBaseConfig struct {
	Entity
	Name           string                            `json:"name"`
	FilenamePrefix string                            `json:"filenamePrefix"`
	FilenameSuffix string                            `json:"filenameSuffix"`
	DocumentNumber string                            `json:"documentNumber"`
	Global         bool                              `json:"global"`
	DocumentTypeID string                            `json:"documentTypeId"`
	LogoID         string                            `json:"logoId"`
	config         []string                          `json:"config"`
	salesChannels  []*DocumentBaseConfigSalesChannel `json:"salesChannels"`
	DocumentType   *DocumentType                     `json:"documentType"`
	Logo           *MediaEntity                      `json:"logo"`
}
