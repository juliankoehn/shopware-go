package shopware

// Tag see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/Tag/TagEntity.php
type Tag struct {
	Entity
	Name                 string                 `json:"name"`
	Products             []*Product             `json:"products"`
	Media                []*MediaEntity         `json:"media"`
	Categories           []*Category            `json:"categories"`
	Customers            []*Customer            `json:"customers"`
	orders               []*Order               `json:"orders"`
	shippingMethods      []*ShippingMethod      `json:"shippingMethods"`
	newsletterRecipients []*NewsletterRecipient `json:"newsletterRecipients"`
}
