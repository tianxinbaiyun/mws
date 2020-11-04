package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/mws/routers/api"
	"github.com/tianxinbaiyun/mws/routers/api/product"
	"github.com/tianxinbaiyun/mws/routers/api/report"

	_ "github.com/tianxinbaiyun/mws/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/auth", api.GetAuth)
	r.GET("/product/categories", product.GetProductCategoriesForSKU)
	r.GET("/product/categories_by_report", product.GetProductCategoriesByReport)
	r.GET("/report/request_report", report.RequestReport)
	r.GET("/report/get_report", report.GetReport)
	r.GET("/report/get_report_request_list", report.GetReportRequestList)
	r.GET("/report/get_report_list", report.GetReportList)
	return r
}
