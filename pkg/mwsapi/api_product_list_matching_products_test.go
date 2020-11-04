package mwsapi

import (
	"context"
	"testing"
)

func TestApiProduct(t *testing.T) {
	product := Products()
	ctx := context.Background()
	credential := GetCredentialForTest()
	query := "c2t"
	requestID, res, err := product.ListMatchingProducts(ctx, credential, query)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(requestID)
	if len(res.ListMatchingProductsResult.Products) > 0 {
		for _, item := range res.ListMatchingProductsResult.Products {
			t.Log(item)
		}
	}

	t.Log(err)
}
