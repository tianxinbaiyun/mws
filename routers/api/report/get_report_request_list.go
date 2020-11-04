package report

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/pkg/app"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/service/report_service"
	"net/http"
	"strings"
	"time"
)

func GetReportRequestList(c *gin.Context) {
	appG := app.Gin{C: c}

	reportRequestIdList := c.DefaultQuery("report_request_id_list", "")
	reportTypeList := c.DefaultQuery("report_type_list", "")
	reportProcessingStatusList := c.DefaultQuery("report_processing_status_list", "")
	maxCount := c.DefaultQuery("max_count", "")
	requestedFromDate := c.DefaultQuery("requested_from_date", "")
	requestedToDate := c.DefaultQuery("requested_to_date", "")
	nextToken := c.DefaultQuery("next_token", "")

	req := map[string]string{}
	// ReportRequestId 列表参数
	if reportRequestIdList != "" {
		reportRequestIds := strings.Split(reportRequestIdList, ",")
		if len(reportRequestIds) > 0 {
			for i, v := range reportRequestIds {
				req[fmt.Sprintf("ReportRequestIdList.Id.%d", i+1)] = v
			}
		}
	}
	// ReportTypeList 类型列表参数
	if reportTypeList != "" {
		reportTypes := strings.Split(reportTypeList, ",")
		if len(reportTypes) > 0 {
			for i, v := range reportTypes {
				req[fmt.Sprintf("ReportTypeList.Type.%d", i+1)] = v
			}
		}
	}
	// ReportProcessingStatusList 状态列表
	if reportProcessingStatusList != "" {
		reportProcessingStatuss := strings.Split(reportProcessingStatusList, ",")
		if len(reportProcessingStatuss) > 0 {
			for i, v := range reportProcessingStatuss {
				req[fmt.Sprintf("ReportProcessingStatusList.status.%d", i+1)] = v
			}
		}
	}
	if maxCount != "" {
		req["MaxCount"] = maxCount
	}
	if requestedFromDate != "" {
		date, _ := now.Parse(requestedFromDate)
		req["RequestedFromDate"] = now.With(date).Format(time.RFC3339)
	}
	if requestedToDate != "" {
		date, _ := now.Parse(requestedToDate)
		req["RequestedToDate"] = now.With(date).Format(time.RFC3339)
	}
	// 获取数据
	reportService := report_service.ReportService{}
	if nextToken != "" {
		rsp, err := reportService.GetReportRequestListByNextToken(nextToken)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, rsp)
	} else {
		rsp, err := reportService.GetReportRequestList(req)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, rsp)
	}
}
