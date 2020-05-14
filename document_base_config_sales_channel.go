package shopware

// DocumentBaseConfigSalesChannel see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Document/Aggregate/DocumentBaseConfigSalesChannel/DocumentBaseConfigSalesChannelEntity.php
type DocumentBaseConfigSalesChannel struct {
	Entity
	DocumentBaseConfigID string              `json:"documentBaseConfigId"`
	SalesChannelID       string              `json:"salesChannelId"`
	DocumentTypeID       string              `json:"documentTypeId"`
	DocumentType         *DocumentType       `json:"documentType"`
	DocumentBaseConfig   *DocumentBaseConfig `json:"documentBaseConfig"`
	SalesChannel         *SalesChannel       `json:"salesChannel"`
}
