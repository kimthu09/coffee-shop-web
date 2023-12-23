package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"errors"
	"time"
)

type ImportNote struct {
	Id            string                                   `json:"id" gorm:"column:id;"`
	SupplierId    string                                   `json:"-" gorm:"column:supplierId;"`
	Supplier      suppliermodel.SimpleSupplier             `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice    int                                      `json:"totalPrice" gorm:"column:totalPrice;"`
	Status        *ImportNoteStatus                        `json:"status" gorm:"column:status;"`
	CreatedBy     string                                   `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser                     `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	ClosedBy      *string                                  `json:"-" gorm:"column:closedBy;"`
	ClosedByUser  usermodel.SimpleUser                     `json:"closedBy" gorm:"foreignKey:ClosedBy;references:Id"`
	CreatedAt     *time.Time                               `json:"createdAt" gorm:"column:createdAt;"`
	ClosedAt      *time.Time                               `json:"closedAt" gorm:"column:closedAt;"`
	Details       []importnotedetailmodel.ImportNoteDetail `json:"details,omitempty" gorm:"-"`
}

func (*ImportNote) TableName() string {
	return common.TableImportNote
}

var (
	ErrImportNoteIdInvalid = common.NewCustomError(
		errors.New("id of import note is invalid"),
		"Mã phiếu nhập không hợp lệ",
		"ErrImportNoteIdInvalid",
	)
	ErrImportNoteSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"Nhà cung cấp của phiếu nhập không hợp lệ",
		"ErrImportNoteSupplierIdInvalid",
	)
	ErrImportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"Danh sách các nguyên vật liệu muốn nhập đang trống",
		"ErrImportNoteDetailsEmpty",
	)
	ErrImportNoteStatusEmpty = common.NewCustomError(
		errors.New("import's status is empty"),
		"Trạng thái muốn thay đổi đang trống",
		"ErrImportNoteStatusEmpty",
	)
	ErrImportNoteStatusInvalid = common.NewCustomError(
		errors.New("import's status is invalid"),
		"Trạng thái muốn thay đổi không hợp lệ",
		"ErrImportNoteStatusInvalid",
	)
	ErrImportNoteExistDuplicateIngredient = common.NewCustomError(
		errors.New("exist duplicate ingredient"),
		"Tồn tại 2 nguyên vật liệu trùng nhau",
		"ErrImportNoteExistDuplicateIngredient",
	)
	ErrImportNoteClosed = common.NewCustomError(
		errors.New("import note has been closed"),
		"Phiếu nhập đã đóng",
		"ErrImportNoteClosed",
	)
	ErrImportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu nhập mới"),
	)
	ErrImportNoteChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa trạng thái phiếu nhập"),
	)
	ErrImportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem phiếu nhập"),
	)
)
