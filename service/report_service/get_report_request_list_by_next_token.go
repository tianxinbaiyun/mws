package report_service

import (
	"context"
	"github.com/tianxinbaiyun/mws/models"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/mwsapi"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"strconv"
)

// GetReportRequestListByNextToken 返回可用于获取报告的 ReportRequestId 的报告请求列表
func (handle *ReportService) GetReportRequestListByNextToken(nextToken string) (rsp *mwsapi.ReportRequestListResult, err error) {
	report := mwsapi.Reports()
	ctx := context.Background()
	credential := setting.MwsSetting
	_, rsp, err = report.GetReportRequestListByNextToken(ctx, credential, nextToken)
	if err != nil {
		logging.Error(err)
		return
	}
	if len(rsp.ReportRequestInfo) > 0 {
		for _, item := range rsp.ReportRequestInfo {
			// 判断报告是否存在
			exist, e := models.ExistReports(&models.Reports{
				SellerID:        setting.MwsSetting.SellerID,
				MarketplaceID:   setting.MwsSetting.MarketplaceID,
				ReportRequestID: item.ReportRequestID,
			})
			if e != nil {
				logging.Error(err)
				return
			}
			if exist {
				continue
			}
			// 保存报告
			report := models.Reports{
				SellerID:               setting.MwsSetting.SellerID,
				MarketplaceID:          setting.MwsSetting.MarketplaceID,
				SellerSKU:              "",
				Asin:                   "",
				ReportRequestID:        item.ReportRequestID,
				ReportType:             item.ReportType,
				StartDate:              item.StartDate,
				EndDate:                item.EndDate,
				Scheduled:              strconv.FormatBool(item.Scheduled),
				SubmittedDate:          item.SubmittedDate,
				ReportProcessingStatus: item.ReportProcessingStatus,
				GeneratedReportID:      item.GeneratedReportID,
				StartedProcessingDate:  item.StartedProcessingDate,
				CompletedDate:          item.CompletedDate,
			}
			err = models.AddReports(&report)
			if err != nil {
				logging.Error(err)
				return
			}
		}

	}
	return
}
