package report

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/service/report_service"
	"net/http"
)

func RequestReport(c *gin.Context) {
	appG := app.Gin{C: c}
	reportType := c.DefaultQuery("report_type", "")
	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")
	reportOptions := c.DefaultQuery("report_options", "")
	if reportType == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	// 参数处理
	req := &report_service.ReportTypeReq{
		ReportType:    reportType,
		ReportOptions: reportOptions,
	}
	if startDate != "" {
		req.StartDate, _ = now.Parse(startDate)
	}
	if endDate != "" {
		req.EndDate, _ = now.Parse(endDate)
	}
	reportService := report_service.ReportService{}
	rsp, err := reportService.RequestReport(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, rsp)
}
