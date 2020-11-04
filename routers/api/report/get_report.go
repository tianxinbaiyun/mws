package report

import (
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/service/report_service"
	"net/http"
)

func GetReport(c *gin.Context) {
	appG := app.Gin{C: c}
	reportID := c.DefaultQuery("report_id", "")
	if reportID == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	reportService := report_service.ReportService{}
	rsp, err := reportService.GetReport(reportID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, rsp)
}
