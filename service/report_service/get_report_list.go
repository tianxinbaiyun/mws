package report_service

import (
	"context"
	"encoding/json"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/mwsapi"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"time"
)

// GetReportListReq 报告列表请求结构体
type GetReportListReq struct {
	MaxCount            string
	ReportTypeList      string
	Acknowledged        string
	AvailableFromDate   time.Time
	AvailableToDate     time.Time
	ReportRequestIdList string
}

// GetReportRequestList 返回可用于获取报告的 ReportRequestId 的报告请求列表
func (handle *ReportService) GetReportList(req *GetReportListReq) (rsp *mwsapi.ReportListResult, err error) {
	report := mwsapi.Reports()
	ctx := context.Background()
	credential := setting.MwsSetting

	b, err := json.Marshal(req)
	if err != nil {
		logging.Error(err)
		return
	}
	logging.Info(string(b))
	parValues := map[string]string{}
	err = json.Unmarshal(b, &parValues)
	if err != nil {
		logging.Error(err)
		return
	}

	_, rsp, err = report.GetReportList(ctx, credential, parValues)
	if err != nil {
		logging.Error(err)
		return
	}
	return
}
