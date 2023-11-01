package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type ExportNote struct {
	Id         string     `json:"id" gorm:"column:id;"`
	TotalPrice float32    `json:"totalPrice" gorm:"column:totalPrice;"`
	CreateBy   string     `json:"createBy" gorm:"column:createBy;"`
	CreateAt   *time.Time `json:"createAt" gorm:"column:createAt;"`
}

func (*ExportNote) TableName() string {
	return common.TableExportNote
}

var (
	ErrSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrSupplierIdInvalid",
	)
	ErrImportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"list import note details are empty",
		"ErrImportNoteDetailsEmpty",
	)
	ErrImportStatusEmpty = common.NewCustomError(
		errors.New("import's status is empty"),
		"import's status is empty",
		"ErrImportStatusEmpty",
	)
	ErrImportStatusInvalid = common.NewCustomError(
		errors.New("import's status is invalid"),
		"import's status is invalid",
		"ErrImportStatusInvalid",
	)
)
