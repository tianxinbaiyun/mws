package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Reports 报告表
type Reports struct {
	Model
	SellerID               string    `json:"seller_id" sql:"index" `
	MarketplaceID          string    `json:"marketplace_id" sql:"index" `
	SellerSKU              string    `json:"seller_sku" sql:"index" `
	Asin                   string    `json:"asin"`
	ReportRequestID        string    `json:"report_request_id"`
	ReportType             string    `json:"report_type"`
	StartDate              time.Time `json:"start_date"`
	EndDate                time.Time `json:"end_date"`
	Scheduled              string    `json:"scheduled"`
	SubmittedDate          time.Time `json:"submitted_date"`
	ReportProcessingStatus string    `json:"report_processing_status"`
	GeneratedReportID      string    `json:"generated_report_id"`
	StartedProcessingDate  time.Time `json:"started_processing_date"`
	CompletedDate          time.Time `json:"completed_date"`
}

// ExistReports 检查报告是否存在
func ExistReports(par *Reports) (bool, error) {
	var data Reports
	find := db.Select("uuid,report_request_id").
		Where("seller_id = ? ", par.SellerID).
		Where("marketplace_id = ? ", par.MarketplaceID)

	if par.ReportRequestID != "" {
		find = find.Where("report_request_id = ?", par.ReportRequestID)
	}
	if par.GeneratedReportID != "" {
		find = find.Where("generated_report_id = ?", par.GeneratedReportID)
	}
	err := find.First(&data).Error
	if err != nil {
		return false, err
	}

	if data.ReportRequestID != "" {
		return true, nil
	}

	return false, nil
}

// AddProductCategories 保存报表
func AddReports(par *Reports) error {
	newUUID, _ := uuid.NewV4()
	data := Reports{
		Model: Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
			UUID:      newUUID,
		},
		SellerID:               par.SellerID,
		MarketplaceID:          par.MarketplaceID,
		SellerSKU:              par.SellerSKU,
		Asin:                   par.Asin,
		ReportRequestID:        par.ReportRequestID,
		ReportType:             par.ReportType,
		StartDate:              par.StartDate,
		EndDate:                par.EndDate,
		Scheduled:              par.Scheduled,
		SubmittedDate:          par.SubmittedDate,
		ReportProcessingStatus: par.ReportProcessingStatus,
		GeneratedReportID:      par.GeneratedReportID,
		StartedProcessingDate:  par.StartedProcessingDate,
		CompletedDate:          par.CompletedDate,
	}
	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
