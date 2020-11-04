package product

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/service/product_service"
	"net/http"
)

func GetProductCategoriesByReport(c *gin.Context) {
	appG := app.Gin{C: c}
	startDateStr := c.DefaultQuery("start_date", "")
	endDateStr := c.DefaultQuery("end_date", "")

	if startDateStr == "" || endDateStr == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	startDate, _ := now.Parse(startDateStr)
	endDate, _ := now.Parse(endDateStr)
	productCategoriesService := product_service.ProductService{}
	rsp, err := productCategoriesService.GetProductCategoriesByReport(startDate, endDate)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, rsp)
}
