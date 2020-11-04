package mwsapi

import (
	"encoding/json"
	"fmt"
	"os"
)

//Credential 令牌
type Credential struct {
	SellerID       string //商户: SellerId A1DGTLIXD8902T
	MWSAuthToken   string //商户: 令牌 amzn.mws.06b005f2-1f84-f56f-91d7-cf4875d4a429
	AWSAccessKeyID string //开发者凭证: AWSAccessKeyId AKIAJU6QRROTSUT3T7EA"
	SecretKey      string //开发者凭证: SecretKey 8jFyFpJAo8Fy6Ys8J8DpSDkuCOL/xw7a8r89+KgV
	MarketplaceID  string //站点: MarketplaceId A1F83G8C2ARO7P
}

//GetCredentialFromEnv 从环境变量获取授权令牌
func GetCredentialFromEnv(sellerID, mwsAuthToken string) *Credential {
	return &Credential{
		SellerID:       sellerID,
		MWSAuthToken:   mwsAuthToken,
		AWSAccessKeyID: os.Getenv("AWSAccessKeyId"),
		SecretKey:      os.Getenv("SecretKey"),
		MarketplaceID:  os.Getenv("MarketplaceId"),
	}
}

//GetCredentialForTest 从环境变量获取授权令牌（用于测试）
func GetCredentialForTest() *Credential {
	if os.Getenv("TestSellerId") != "" {
		return GetCredentialFromEnv(os.Getenv("TestSellerId"), os.Getenv("TestMWSAuthToken"))
	} else {
		credential := Credential{}
		fname := ".gomws_config"
		f, err := os.Open(fname)
		if err != nil {
			fmt.Println("Couldn't find an aws config file. you need one.")
		}
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&credential)
		if err != nil {
			fmt.Println(err)
		}
		return &credential
	}

}
