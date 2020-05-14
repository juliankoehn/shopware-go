package shopware

// Document see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Document/DocumentEntity.php
type Document struct {
	Entity
	OrderID              string        `json:"orderId"`
	OrderVersionID       string        `json:"orderVersionId"`
	DocumentTypeID       string        `json:"documentTypeId"`
	DocumentMediaFileID  string        `json:"documentMediaFileId"`
	FileType             string        `json:"fileType"`
	Order                *Order        `json:"order"`
	Config               []string      `json:"config"`
	Sent                 bool          `json:"sent"`
	Static               bool          `json:"static"`
	DeepLinkCode         string        `json:"deepLinkCode"`
	DocumentType         *DocumentType `json:"documentType"`
	ReferencedDocumentID string        `json:"referencedDocumentId"`
	ReferencedDocument   *Document     `json:"referencedDocument"`
	DependentDocuments   []*Document   `json:"dependentDocuments"`
	DocumentMediaFile    *MediaEntity  `json:"documentMediaFile"`
}
