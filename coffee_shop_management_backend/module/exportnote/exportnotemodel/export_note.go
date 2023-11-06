package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type ExportNote struct {
	Id         string     `json:"id" gorm:"column:id;"`
	TotalPrice float32    `json:"totalPrice" gorm:"column:totalPrice;"`
	CreateAt   *time.Time `json:"createAt" gorm:"column:createAt;"`
	CreateBy   string     `json:"createBy" gorm:"column:createBy;"`
}

func (*ExportNote) TableName() string {
	return common.TableExportNote
}

var (
	ErrExportNoteIdInvalid = common.NewCustomError(
		errors.New("id of export note is invalid"),
		"id of export note is invalid",
		"ErrExportNoteIdInvalid",
	)
	ErrExportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list export note details are empty"),
		"list export note details are empty",
		"ErrExportNoteDetailsEmpty",
	)
	ErrExportNoteAmountCancelIsOverTheStock = common.NewCustomError(
		errors.New("amount export is over stock"),
		"amount export is over stock",
		"ErrExportNoteAmountCancelIsOverTheStock",
	)
	ErrExportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create export note"),
	)
)
