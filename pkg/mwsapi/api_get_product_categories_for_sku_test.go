package mwsapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetProductCategoriesForSKU(t *testing.T) {

	product := Products()
	ctx := context.Background()
	credential := GetCredentialForTest()
	sellerSKU := "T4ft6-mattress"
	_, res, err := product.GetProductCategoriesForSKU(ctx, credential, sellerSKU)
	if err != nil {
		t.Log(err)
		assert.Error(t, err)
	}
	t.Log(res)

}

func TestGetProductCategoriesForSKU2(t *testing.T) {
	product := Products()
	ctx := context.Background()
	credential := GetCredentialForTest()
	sellerSKU := "T4ft6-mattress"

	for {
		ti := time.NewTimer(time.Second * 60)
		<-ti.C
		// 获取数据
		_, res, err := product.GetProductCategoriesForSKU(ctx, credential, sellerSKU)
		if err != nil {
			t.Log(err)
			assert.Error(t, err)
		}
		t.Log(res)
	}
}
