package product_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/tianxinbaiyun/mws/models"
	"github.com/tianxinbaiyun/mws/pkg/e"
	"github.com/tianxinbaiyun/mws/pkg/logging"
	"github.com/tianxinbaiyun/mws/pkg/mwsapi"
	"github.com/tianxinbaiyun/mws/pkg/setting"
	"github.com/tianxinbaiyun/mws/service/report_service"
	"time"
)

// GetProductCategoriesForSKU 获取分类
func (handle *ProductService) GetProductCategoriesForSKU(sellerSKU string) (rsp mwsapi.GetProductCategoriesForSKUResponse, err error) {
	product := mwsapi.Products()
	ctx := context.Background()
	credential := setting.MwsSetting
	// 查询sku是否存在
	exist, err := models.ExistProductCategories(&models.ProductCategories{
		SellerID:      credential.SellerID,
		MarketplaceID: credential.MarketplaceID,
		SellerSKU:     sellerSKU,
	})
	// 如果存在退出
	if exist {
		err = errors.New(e.MsgFlags[e.ERROR_EXIST_SKU])
		return
	}
	_, rsp, err = product.GetProductCategoriesForSKU(ctx, credential, sellerSKU)
	if err != nil {
		logging.Error(err)
		return
	}
	if len(rsp.GetProductCategoriesForSKUResult) > 0 {
		for _, item := range rsp.GetProductCategoriesForSKUResult {
			logging.Info(fmt.Printf("GetProductCategoriesForSKUResult.item:%v", item))
			productCategories := models.ProductCategories{
				Model:                     models.Model{},
				SellerID:                  setting.MwsSetting.SellerID,
				MarketplaceID:             setting.MwsSetting.MarketplaceID,
				SellerSKU:                 sellerSKU,
				Asin:                      "",
				ProductCategoryID:         item.ProductCategoryId,
				ProductCategoryName:       item.ProductCategoryName,
				ParentProductCategoryID:   item.ParentProductCategoryId,
				ParentProductCategoryName: item.ParentProductCategoryName,
				SecondProductCategoryID:   item.SecondProductCategoryId,
				SecondProductCategoryName: item.SecondProductCategoryName,
			}
			err = models.AddProductCategories(&productCategories)
			if err != nil {
				logging.Error(err)
				return
			}
		}
	}

	return
}

// GetProductCategoriesByReport 通过报表获取分类
func (handle *ProductService) GetProductCategoriesByReport(startDate, endDate time.Time) (rsp int, err error) {
	reportService := report_service.ReportService{}
	req := map[string]string{}
	req["RequestedFromDate"] = now.With(startDate).Format(time.RFC3339)
	req["RequestedToDate"] = now.With(endDate).Format(time.RFC3339)
	req["ReportTypeList.Type.1"] = "_GET_FLAT_FILE_OPEN_LISTINGS_DATA_"
	// 获取报表
	reportRsp, err := reportService.GetReportRequestList(req)
	if err != nil {
		logging.Error(err)
		return
	}
	if len(reportRsp.ReportRequestInfo) > 0 {
		for _, item := range reportRsp.ReportRequestInfo {
			if item.ReportProcessingStatus == "_DONE_" {
				// 获取get_report
				report, err := reportService.GetReport(item.GeneratedReportID)
				if err != nil {
					logging.Error(err)
				}
				for _, reportItem := range report {
					_, err = handle.GetProductCategoriesForSKU(reportItem[0])
					if err == nil {
						rsp++
					}
				}
			}
		}
	}
	return
}
