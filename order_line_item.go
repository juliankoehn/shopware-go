package shopware

// OrderLineItem seE:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Order/Aggregate/OrderLineItem/OrderLineItemEntity.php
type OrderLineItem struct {
	Entity
	OrderID                string                    `json:"orderId"`
	Identifier             string                    `json:"identifier"`
	ReferencedID           string                    `json:"referencedId"`
	ProductID              string                    `json:"productId"`
	Quantity               int                       `json:"quantity"`
	UnitPrice              float64                   `json:"unitPrice"`
	TotalPrice             float64                   `json:"totalPrice"`
	Label                  string                    `json:"label"`
	Description            string                    `json:"description"`
	Good                   bool                      `json:"good"`
	Removable              bool                      `json:"removable"`
	CoverID                string                    `json:"coverId"`
	Stackable              bool                      `json:"stackable"`
	Position               int                       `json:"position"`
	Price                  *CalculatedPrice          `json:"price"`
	PriceDefinition        *PriceDefinitionInterface `json:"priceDefinition"`
	Payload                []string                  `json:"payload"`
	ParentID               string                    `json:"parentId"`
	Typ                    string                    `json:"type"` // shopware: type
	Order                  *Order                    `json:"order"`
	OrderDeliveryPositions []*OrderDeliveryPosition  `json:"orderDeliveryPositions"`
	Cover                  *MediaEntity              `json:"cover"`
	Children               []*OrderLineItem          `json:"children"`
	Product                *Product                  `json:"product"`
}
