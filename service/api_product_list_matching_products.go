package mws

import (
	"context"
	"fmt"
	"src/mws/practice/core"
)

//ListMatchingProductsResponse 商品价格获取响应体
type ListMatchingProductsResponse struct {
	core.BaseResponse
	ListMatchingProductsResult ListMatchingProductsResult
}

//ListMatchingProductsResult ListMatchingProductsResult
type ListMatchingProductsResult struct {
	Products []*Product `xml:"Products>Product"`
}

//Product 商品信息模型
type Goods struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN 商城
}

// ListMatchingProducts 根据搜索查询，返回商品及其属性列表。
//
// ListMatchingProducts 操作会根据您指定的 SellerSKU 和 MarketplaceId，返回您自己的商品的价格信息。请注意，如果您提交了并未销售的商品的 SellerSKU，则此操作会返回空的 Offers 元素。此操作最多可返回 20 款商品的价格信息。
func (s *ProductService) ListMatchingProducts(ctx context.Context, c *Credential, query string) (requestID string, result ListMatchingProductsResponse, err error) {
	data := ActionValues("ListMatchingProducts")
	data.Set("MarketplaceId", c.MarketplaceID)
	data.Set("Query", query)
	fmt.Println(data)
	if _, err = s.FetchStruct(ctx, c, data, &result); err != nil {
		return
	}
	return
}
