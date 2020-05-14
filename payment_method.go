package shopware

// PaymentMethod see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Payment/PaymentMethodEntity.php
type PaymentMethod struct {
	Entity
	PluginID                       string                      `json:"pluginId"`
	HandlerIdentifier              string                      `json:"handlerIdentifier"`
	Name                           string                      `json:"name"`
	Description                    string                      `json:"description"`
	Position                       int                         `json:"position"`
	Active                         bool                        `json:"active"`
	AfterOrderEnabled              bool                        `json:"afterOrderEnabled"`
	Plugin                         *Plugin                     `json:"plugin"`
	Translations                   []*PaymentMethodTranslation `json:"translations"`
	OrderTransactions              []*OrderTransaction         `json:"orderTransactions"`
	Customers                      []*Customer                 `json:"customers"`
	SalesChannelDefaultAssignments []*SalesChannel             `json:"salesChannelDefaultAssignments"`
	SalesChannels                  []*SalesChannel             `json:"salesChannels"`
	AvailabilityRule               *Rule                       `json:"availabilityRule"`
	AvailabilityRuleID             string                      `json:"availabilityRuleId"`
	MediaID                        string                      `json:"mediaId"`
	Media                          *MediaEntity                `json:"media"`
}
