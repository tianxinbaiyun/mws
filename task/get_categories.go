package task

import (
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"github.com/tianxinbaiyun/mws/service/product_service"
	"time"
)

func GetCategories() {
	config := setting.TaskSetting
	i := 1
	for {
		//定时操作
		ti := time.NewTimer(time.Second * config.GetCategoriesTimeDuration)
		<-ti.C
		logging.Info("GetCategories 次数：", i)
		i++
		// 获取服务
		productService := product_service.ProductService{}

		startDate := now.BeginningOfMonth()
		endDate := time.Now()
		// 获取报表
		_, err := productService.GetProductCategoriesByReport(startDate, endDate)
		if err != nil {
			logging.Error(err)
		}
	}

}
