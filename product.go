package shopware

import (
	"fmt"
	"net/http"
)

// ProductsService model
type ProductsService service

// Product model
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Product/ProductEntity.php
// Invalid relation type
// Price is "array" not "float"
// TODO: dates are currently strings
type Product struct {
	Active                       bool                                  `json:"active"`
	AutoIncrement                int                                   `json:"autoIncrement"`
	Available                    bool                                  `json:"available"`
	AvailableStock               int                                   `json:"availableStock"`
	BlacklistIDs                 []string                              `json:"blacklistIds"`
	WhitelistIDs                 []string                              `json:"whitelistIds"`
	Categories                   []*Category                           `json:"categories"`
	CategoriesRo                 []*Category                           `json:"categoriesRo"`
	CategoryTree                 []string                              `json:"categoryTree"`
	Childcount                   int                                   `json:"childCount"`
	Children                     []*Product                            `json:"children"`
	ConfiguratorSettings         []*ProductConfiguratorSetting         `json:"configuratorSettings"`
	CoverID                      string                                `json:"coverId"`
	Cover                        ProductMediaEntity                    `json:"cover"`
	CrossSellings                []*ProductCrossSelling                `json:"crossSellings"`
	CrossSellingAssignedProducts []*ProductCrossSellingAssignedProduct `json:"crossSellingAssignedProducts"`
	Description                  string                                `json:"description"`
	DeliveryTimeID               string                                `json:"deliveryTimeId"`
	DisplayGroup                 string                                `json:"displayGroup"`
	EAN                          string                                `json:"ean"`
	Height                       float64                               `json:"height"`
	ID                           string                                `json:"id"`
	IsCloseout                   bool                                  `json:"isCloseout"`
	IsNew                        bool                                  `json:"isNew"`
	Keywords                     string                                `json:"keywords"`
	Length                       float64                               `json:"length"`
	ListingPrices                []ListingPrice                        `json:"listingPrices"`
	MainCategories               []*MainCategory                       `json:"mainCategories"`
	MainVariantID                string                                `json:"mainVariantId"`
	Manufacturer                 ProductManufacturer                   `json:"manufacturer"`
	ManufacturerID               string                                `json:"manufacturerId"`
	ManufacturerNumber           string                                `json:"manufacturerNumber"`
	MarkAsTopseller              bool                                  `json:"markAsTopseller"`
	MaxPurchase                  int                                   `json:"maxPurchase"`
	Media                        []*ProductMediaEntity                 `json:"media"`
	MetaDescription              string                                `json:"metaDescription"`
	MetaTitle                    string                                `json:"metaTitle"`
	MinPurchase                  int                                   `json:"minPurchase"`
	Name                         string                                `json:"name"`
	OptionIds                    []string                              `json:"optionIds"`
	OrderLineItems               []*OrderLineItem                      `json:"orderLineItems"`
	PackUnit                     string                                `json:"packUnit"`
	Parent                       *Product                              `json:"parent"`
	ParentID                     string                                `json:"parentId"`
	PurchasePrice                float64                               `json:"purchasePrice"`
	PurchaseSteps                int                                   `json:"purchaseSteps"`
	PurchaseUnit                 float64                               `json:"purchaseUnit"`
	Price                        []Price                               `json:"price"`
	Prices                       []Price                               `json:"prices"`
	ProductNumber                string                                `json:"productNumber"`
	ProductReviews               []*ProductReview                      `json:"productReviews"`
	RatingAverage                float64                               `json:"ratingAverage"`
	ReleaseDate                  string                                `json:"releaseDate"`
	ReferenceUnit                float64                               `json:"referenceUnit"`
	RestockTime                  int                                   `json:"restockTime"`
	SeoUrls                      []*SeoURL                             `json:"seoUrls"`
	ShippingFree                 bool                                  `json:"shippingFree"`
	Stock                        int                                   `json:"stock"`
	Tags                         []*Tag                                `json:"tags"`
	TagIDs                       []string                              `json:"tagIds"`
	Tax                          Tax                                   `json:"tax"`
	TaxID                        string                                `json:"taxId"`
	Translated                   ProductTranslation                    `json:"translated"`
	Translations                 []*ProductTranslation                 `json:"translations"`
	UnitID                       string                                `json:"unitId"`
	Visibilities                 []*ProductVisibility                  `json:"visibilities"`
	Weight                       float64                               `json:"weight"`
	Width                        float64                               `json:"width"`
	CreatedAt                    string                                `json:"createdAt"`
	UpdatedAt                    string                                `json:"updatedAt"`
}

// ProductGetResponse is a Get API Response
type ProductGetResponse struct {
	Data Product `json:"data"`
}

// Get returns a single product
func (service *ProductsService) Get(productID string) (*Product, error) {
	path := fmt.Sprintf("/sales-channel-api/v1/product/%s", productID)

	req, err := service.c.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var productResponse ProductGetResponse
	if ok := service.c.do(req, &productResponse); ok != nil {
		return nil, ok
	}

	return &productResponse.Data, nil
}
