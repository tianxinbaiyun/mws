package report_service

import (
	"context"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/mwsapi"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"time"
)

type ReportTypeReq struct {
	ReportType        string
	StartDate         time.Time
	EndDate           time.Time
	ReportOptions     string
	MarketplaceIdList string
}

// RequestReport 获取报表
func (handle *ReportService) RequestReport(req *ReportTypeReq) (rsp *mwsapi.ReportRequestInfo, err error) {
	report := mwsapi.Reports()
	ctx := context.Background()
	credential := setting.MwsSetting

	_, rsp, err = report.RequestReport(ctx, credential, req.ReportType)
	if err != nil {
		logging.Error(err)
		return
	}
	return
}
