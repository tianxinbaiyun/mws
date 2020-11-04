package report_service

import (
	"context"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/mwsapi"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"github.com/tianxinbaiyun/mws/pkg/util"
	"strings"
)

// RequestReport 获取报表
func (handle *ReportService) GetReport(reportID string) (rsp [][]string, err error) {
	report := mwsapi.Reports()
	ctx := context.Background()
	credential := setting.MwsSetting

	response, err := report.GetReport(ctx, credential, reportID)
	if err != nil {
		logging.Error(err)
		return
	}
	rspStr := util.Bytes2str(response)
	arr := strings.Split(rspStr, "\r\n")
	arrList := make([][]string, len(arr))
	if len(arr) > 0 {
		for i := 1; i < len(arr); i++ {
			arrList[i] = strings.Split(arr[i], "\t")
		}
		rsp = arrList[1:]
	}
	return
}
