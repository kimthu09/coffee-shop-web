package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type ImportNote struct {
	Id         string            `json:"id" gorm:"column:id;"`
	SupplierId string            `json:"supplierId" gorm:"column:supplierId;"`
	TotalPrice float32           `json:"totalPrice" gorm:"column:totalPrice;"`
	Status     *ImportNoteStatus `json:"status" gorm:"column:status;"`
	CreateBy   string            `json:"createBy" gorm:"column:createBy;"`
	CloseBy    *string           `json:"closeBy" gorm:"column:closeBy;"`
	CreateAt   *time.Time        `json:"createAt" gorm:"column:createAt;"`
	CloseAt    *time.Time        `json:"closeAt" gorm:"column:closeAt;"`
}

func (*ImportNote) TableName() string {
	return common.TableImportNote
}

var (
	ErrImportNoteIdInvalid = common.NewCustomError(
		errors.New("id of import note is invalid"),
		"id of import note is invalid",
		"ErrImportNoteIdInvalid",
	)
	ErrImportNoteSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrImportNoteSupplierIdInvalid",
	)
	ErrImportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"list import note details are empty",
		"ErrImportNoteDetailsEmpty",
	)
	ErrImportNoteStatusEmpty = common.NewCustomError(
		errors.New("import's status is empty"),
		"import's status is empty",
		"ErrImportNoteStatusEmpty",
	)
	ErrImportNoteStatusInvalid = common.NewCustomError(
		errors.New("import's status is invalid"),
		"import's status is invalid",
		"ErrImportNoteStatusInvalid",
	)
	ErrImportNoteHasSameIngredientBothUpdatePrice = common.NewCustomError(
		errors.New("exist one ingredient need to update price twice"),
		"exist one ingredient need to update price twice",
		"ErrImportNoteHasSameIngredientBothUpdatePrice",
	)
	ErrImportNoteClosed = common.NewCustomError(
		errors.New("import note has been closed"),
		"import note has been closed",
		"ErrImportNoteClosed",
	)
	ErrImportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create import note"),
	)
	ErrImportNoteChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to change status import note"),
	)
)
