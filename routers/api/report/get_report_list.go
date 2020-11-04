package report

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/service/report_service"
	"net/http"
)

func GetReportList(c *gin.Context) {
	appG := app.Gin{C: c}

	maxCount := c.DefaultQuery("max_count", "")
	reportTypeList := c.DefaultQuery("report_type_list", "")
	acknowledged := c.DefaultQuery("acknowledged", "")
	reportRequestIdList := c.DefaultQuery("MaxCount", "")
	availableFromDate := c.DefaultQuery("available_from_date", "")
	availableToDate := c.DefaultQuery("available_to_date", "")

	req := &report_service.GetReportListReq{
		MaxCount:            maxCount,
		ReportTypeList:      reportTypeList,
		Acknowledged:        acknowledged,
		ReportRequestIdList: reportRequestIdList,
	}
	if availableFromDate != "" {
		req.AvailableFromDate, _ = now.Parse(availableFromDate)
	} else {
		req.AvailableFromDate = now.BeginningOfMonth()
	}
	if availableToDate != "" {
		req.AvailableToDate, _ = now.Parse(availableToDate)
	} else {
		req.AvailableToDate = now.EndOfMonth()
	}
	reportService := report_service.ReportService{}
	rsp, err := reportService.GetReportList(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, rsp)
}
