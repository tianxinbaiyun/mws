package product

import (
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/service/product_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
)

func GetProductCategoriesForSKU(c *gin.Context) {
	appG := app.Gin{C: c}
	sku := c.DefaultQuery("sku", "")
	logging.Info(sku)
	if sku == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	productCategoriesService := product_service.ProductService{}
	rsp, err := productCategoriesService.GetProductCategoriesForSKU(sku)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, rsp)
}
