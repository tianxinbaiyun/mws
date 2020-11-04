package mwsapi

/* 公共模型 */

//BaseResponse 基础模型，含请求ID
type BaseResponse struct {
	RequestID string `xml:"ResponseMetadata>RequestId"`
}

//Money 金额信息
type Money struct {
	CurrencyCode string  `xml:"CurrencyCode"` //三位数的货币代码。格式为 ISO 4217。
	Amount       float64 `xml:"Amount"`       //货币金额。
}
