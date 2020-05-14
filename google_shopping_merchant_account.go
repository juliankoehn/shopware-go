package shopware

// GoogleShoppingMerchantAccount see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/GoogleShopping/Aggregate/GoogleShoppingMerchantAccount/GoogleShoppingMerchantAccountEntity.php
type GoogleShoppingMerchantAccount struct {
	Entity
	AccountID  string                 `json:"accountId"`
	MerchantID string                 `json:"merchantId"`
	Account    *GoogleShoppingAccount `json:"account"`
}
