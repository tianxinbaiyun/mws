package mwsapi

import (
	"context"
)

type GetProductCategoriesForSKUResponse struct {
	BaseResponse
	GetProductCategoriesForSKUResult []*Self
}

type GetProductCategoriesForSKUResult struct {
	Self []*Self `xml:"GetProductCategoriesForSKUResult>Self"`
}

type Self struct {
	ProductCategoryId         int64  `xml:"Self>ProductCategoryId"`
	ProductCategoryName       string `xml:"Self>ProductCategoryName"`
	ParentProductCategoryId   int64  `xml:"Self>Parent>ProductCategoryId"`
	ParentProductCategoryName string `xml:"Self>Parent>ProductCategoryName"`
	SecondProductCategoryId   int64  `xml:"Self>Parent>Parent>ProductCategoryId"`
	SecondProductCategoryName string `xml:"Self>Parent>Parent>ProductCategoryName"`
}

// ListMatchingProducts 根据搜索查询，返回商品及其属性列表。
//
// ListMatchingProducts 操作会根据您指定的 SellerSKU 和 MarketplaceId，返回您自己的商品的价格信息。请注意，如果您提交了并未销售的商品的 SellerSKU，则此操作会返回空的 Offers 元素。此操作最多可返回 20 款商品的价格信息。
func (s *ProductService) GetProductCategoriesForSKU(ctx context.Context, c *Credential, sellerSKU string) (
	requestID string, result GetProductCategoriesForSKUResponse, err error) {
	data := ActionValues("GetProductCategoriesForSKU")
	data.Set("MarketplaceId", c.MarketplaceID)
	data.Set("SellerSKU", sellerSKU)
	if _, err = s.FetchStruct(ctx, c, data, &result); err != nil {
		return
	}
	requestID = result.RequestID
	return
}
