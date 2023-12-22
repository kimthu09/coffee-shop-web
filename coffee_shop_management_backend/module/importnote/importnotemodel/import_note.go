package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"errors"
	"time"
)

type ImportNote struct {
	Id            string                       `json:"id" gorm:"column:id;"`
	SupplierId    string                       `json:"-" gorm:"column:supplierId;"`
	Supplier      suppliermodel.SimpleSupplier `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice    int                          `json:"totalPrice" gorm:"column:totalPrice;"`
	Status        *ImportNoteStatus            `json:"status" gorm:"column:status;"`
	CreatedBy     string                       `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser         `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	ClosedBy      *string                      `json:"-" gorm:"column:closedBy;"`
	ClosedByUser  usermodel.SimpleUser         `json:"closedBy" gorm:"foreignKey:ClosedBy;references:Id"`
	CreatedAt     *time.Time                   `json:"createdAt" gorm:"column:createdAt;"`
	ClosedAt      *time.Time                   `json:"closedAt" gorm:"column:closedAt;"`
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
	ErrImportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view import note"),
	)
)
