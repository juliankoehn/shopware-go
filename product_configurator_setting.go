package shopware

// ProductConfiguratorSetting see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductConfiguratorSetting/ProductConfiguratorSettingEntity.php
type ProductConfiguratorSetting struct {
	Entity
	ProductID string               `json:"productId"`
	OptionID  string               `json:"optionId"`
	MediaID   string               `json:"mediaId"`
	Position  int                  `json:"position"`
	Price     []*Price             `json:"price"`
	Option    *PropertyGroupOption `json:"option"`
	Media     *MediaEntity         `json:"media"`
	Selected  bool                 `json:"selected"`
	Product   *Product             `json:"product"`
}