package shopware

const (
	/**
	 * SCOPE_CART defines promotion discounts on
	 * the entire cart and its line items.
	 */
	SCOPE_CART = "cart"
	/**
	 * SCOPE_DELIVERY defines promotion discounts on
	 * the delivery costs.
	 */
	SCOPE_DELIVERY = "delivery"
	/**
	 * SCOPE_SET defines promotion discounts on
	 * the whole set of groups
	 */
	SCOPE_SET = "set"
	/**
	 * SCOPE_SETGROUP defines promotion discounts on
	 * a specific set group.
	 */
	SCOPE_SETGROUP = "setgroup"
	/**
	 * TYPE_PERCENTAGE defines a percentage
	 * price definition of the discount.
	 */
	TYPE_PERCENTAGE = "percentage"
	/**
	 * TYPE_ABSOLUTE defines an absolute price
	 * definition of the discount in the
	 * current context currency.
	 */
	TYPE_ABSOLUTE = "absolute"

	// TYPE_FIXED_UNIT defines an fixed unit price
	// definition of the discount.
	TYPE_FIXED_UNIT = "fixed_unit"

	/* TYPE_FIXED defines a fixed price
	 * definition of the discount.
	 */
	TYPE_FIXED = "fixed"
)

// PromotionDiscount see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionDiscount/PromotionDiscountEntity.php
type PromotionDiscount struct {
	Entity
	PromotionID             string                    `json:"promotionId"`
	Scope                   string                    `json:"scope"`
	Typ                     string                    `json:"type"` // shopware: type
	Value                   float64                   `json:"value"`
	Promotion               *Promotion                `json:"promotion"`
	DiscountRules           []*Rule                   `json:"discountRules"`
	ConsiderAdvancedRules   bool                      `json:"considerAdvancedRules"`
	MaxValue                float64                   `json:"maxValue"`
	PromotionDiscountPrices []*PromotionDiscountPrice `json:"promotionDiscountPrices"`
	SorterKey               string                    `json:"sorterKey"`
	ApplierKey              string                    `json:"applierKey"`
	UsageKey                string                    `json:"usageKey"`
}
