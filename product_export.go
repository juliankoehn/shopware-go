package shopware

const ENCODING_UTF8 = "UTF-8"
const ENCODING_ISO88591 = "ISO-8859-1"
const FILE_FORMAT_CSV = "csv"
const FILE_FORMAT_XML = "xml"

// ProductExport see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/ProductExport/ProductExportEntity.php
type ProductExport struct {
	Entity
	ProductStreamID          string              `json:"productStreamId"`
	StorefrontSalesChannelID string              `json:"storefrontSalesChannelId"`
	SalesChannelID           string              `json:"salesChannelId"`
	SalesChannelDomainID     string              `json:"salesChannelDomainId"`
	CurrencyID               string              `json:"currencyId"`
	FileName                 string              `json:"fileName"`
	AccessKey                string              `json:"accessKey"`
	Encoding                 string              `json:"encoding"`
	FileFormat               string              `json:"fileFormat"`
	ProductStream            *ProductStream      `json:"productStream"`
	StorefrontSalesChannel   *SalesChannel       `json:"storefrontSalesChannel"`
	SalesChannel             *SalesChannel       `json:"salesChannel"`
	SalesChannelDomain       *SalesChannelDomain `json:"salesChannelDomain"`
	Currency                 *Currency           `json:"currency"`
	IncludeVariants          bool                `json:"includeVariants"`
	GenerateByCronjob        bool                `json:"generateByCronjob"`
	GeneratedAt              string              `json:"generatedAt"`
	Interval                 int                 `json:"interval"`
	HeaderTemplate           string              `json:"headerTemplate"`
	BodyTemplate             string              `json:"bodyTemplate"`
	FooterTemplate           string              `json:"footerTemplate"`
}
