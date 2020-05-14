package shopware

const (
	CODE_TYPE_NO_CODE = "no_code"
)

// Promotion see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/PromotionEntity.php
type Promotion struct {
	Entity
	Name                      string                   `json:"name"`
	Active                    bool                     `json:"active"`
	ValidFrom                 string                   `json:"validFrom"`
	ValidUntil                string                   `json:"validUntil"`
	MaxRedemptionsGlobal      int                      `json:"maxRedemptionsGlobal"`
	MaxRedemptionsPerCustomer int                      `json:"maxRedemptionsPerCustomer"`
	Exclusive                 bool                     `json:"exclusive"`
	UseCodes                  bool                     `json:"useCodes"`
	UseSetGroups              bool                     `json:"useSetGroups"`
	CustomerRestriction       bool                     `json:"customerRestriction"`
	UseIndividualCodes        bool                     `json:"useIndividualCodes"`
	IndividualCodePattern     string                   `json:"individualCodePattern"`
	SalesChannels             []*PromotionSalesChannel `json:"salesChannels"`
	Code                      string                   `json:"code"`
	Discounts                 []*PromotionDiscount     `json:"discounts"`
	Setgroups                 []*PromotionSetGroup     `json:"setgroups"`
	OrderRules                []*Rule                  `json:"orderRules"`
	PersonaRules              []*Rule                  `json:"personaRules"`
	PersonaCustomers          []*Customer              `json:"personaCustomers"`
	CartRules                 []*Rule                  `json:"cartRules"`
	Translations              []*PromotionTranslation  `json:"translations"`
	OrderCount                int                      `json:"orderCount"`
	OrdersPerCustomerCount    []string                 `json:"ordersPerCustomerCount"`
	ExclusionIDs              []string                 `json:"exclusionIds"`
}
