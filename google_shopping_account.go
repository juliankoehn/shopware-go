package shopware

// GoogleShoppingAccount see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/GoogleShopping/GoogleShoppingAccountEntity.php
type GoogleShoppingAccount struct {
	Entity
	SalesChannelID                string                         `json:"salesChannelId"`
	Email                         string                         `json:"email"`
	name                          string                         `json:"name"`
	credential                    *GoogleAccountCredential       `json:"credential"`
	salesChannel                  *SalesChannel                  `json:"salesChannel"`
	googleShoppingMerchantAccount *GoogleShoppingMerchantAccount `json:"googleShoppingMerchantAccount"`
}
