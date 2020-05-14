package shopware

// PromotionDiscountPrice see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Promotion/Aggregate/PromotionDiscountPrice/PromotionDiscountPriceEntity.php
type PromotionDiscountPrice struct {
	Entity
	CurrencyID        string             `json:"currencyId"`
	DiscountID        string             `json:"discountID"`
	Price             float64            `json:"price"`
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount"`
	Currency          *Currency          `json:"currency"`
}
