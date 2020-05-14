package shopware

// ProductCrossSellingTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/Aggregate/ProductCrossSellingTranslation/ProductCrossSellingTranslationEntity.php
type ProductCrossSellingTranslation struct {
	TranslationEntity
	ProductCrossSellingID string               `json:"productCrossSellingId"`
	Name                  string               `json:"name"`
	ProductCrossSelling   *ProductCrossSelling `json:"productCrossSelling"`
}
