package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// ProductCategories 商品分类
type ProductCategories struct {
	Model
	SellerID                  string `json:"seller_id" sql:"index" `
	MarketplaceID             string `json:"marketplace_id" sql:"index" `
	SellerSKU                 string `json:"seller_sku" sql:"index" `
	Asin                      string `json:"asin"`
	ProductCategoryID         int64  `json:"product_category_id"`
	ProductCategoryName       string `json:"product_category_name"`
	ParentProductCategoryID   int64  `json:"parent_product_category_id"`
	ParentProductCategoryName string `json:"parent_product_category_name"`
	SecondProductCategoryID   int64  `json:"second_product_category_id"`
	SecondProductCategoryName string `json:"product_category_name"`
}

// ExistProductCategories 检查sku对应的
func ExistProductCategories(par *ProductCategories) (bool, error) {
	var categories ProductCategories
	err := db.Select("uuid,seller_sku").
		Where("seller_id = ? ", par.SellerID).
		Where("marketplace_id = ? ", par.MarketplaceID).
		Where("seller_sku = ? ", par.SellerSKU).
		First(&categories).Error
	if err != nil {
		return false, err
	}

	if categories.SellerSKU != "" {
		return true, nil
	}

	return false, nil
}

// AddProductCategories 添加分类
func AddProductCategories(par *ProductCategories) error {
	newUUID, _ := uuid.NewV4()
	productCategories := ProductCategories{
		Model: Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
			UUID:      newUUID,
		},
		SellerID:                  par.SellerID,
		MarketplaceID:             par.MarketplaceID,
		SellerSKU:                 par.SellerSKU,
		Asin:                      par.Asin,
		ProductCategoryID:         par.ProductCategoryID,
		ProductCategoryName:       par.ProductCategoryName,
		ParentProductCategoryID:   par.ParentProductCategoryID,
		ParentProductCategoryName: par.ParentProductCategoryName,
		SecondProductCategoryID:   par.SecondProductCategoryID,
		SecondProductCategoryName: par.SecondProductCategoryName,
	}
	if err := db.Create(&productCategories).Error; err != nil {
		return err
	}

	return nil
}

// GetProductCategories 获取分类
func GetProductCategories(par *ProductCategories) (categories ProductCategories, err error) {
	err = db.Where("seller_id = ? ", par.SellerID).
		Where("marketplace_id = ? ", par.MarketplaceID).
		Where("seller_sku = ? ", par.SellerSKU).
		First(&categories).Error
	if err != nil {
		return
	}

	return
}
