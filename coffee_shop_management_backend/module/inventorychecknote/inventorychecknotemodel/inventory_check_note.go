package inventorychecknotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"errors"
	"time"
)

type InventoryCheckNote struct {
	Id                string                                                   `json:"id" gorm:"column:id;"`
	AmountDifferent   int                                                      `json:"amountDifferent" gorm:"column:amountDifferent;"`
	AmountAfterAdjust int                                                      `json:"amountAfterAdjust" gorm:"column:amountAfterAdjust;" `
	CreatedBy         string                                                   `json:"-" gorm:"column:createdBy;"`
	CreatedByUser     usermodel.SimpleUser                                     `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt         *time.Time                                               `json:"createdAt" gorm:"column:createdAt;"`
	Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetail `json:"details,omitempty"`
}

func (*InventoryCheckNote) TableName() string {
	return common.TableInventoryCheckNote
}

var (
	ErrInventoryCheckNoteIdInvalid = common.NewCustomError(
		errors.New("id of inventory check note is invalid"),
		"Mã của phiếu kiểm kho không hợp lệ",
		"ErrInventoryCheckNoteIdInvalid",
	)
	ErrInventoryCheckNoteDetailEmpty = common.NewCustomError(
		errors.New("exist duplicate book"),
		"Danh sách nguyên vật lệu kiểm kho đang trống",
		"ErrInventoryCheckNoteExistDuplicateBook",
	)
	ErrInventoryCheckNoteExistDuplicateBook = common.NewCustomError(
		errors.New("exist duplicate book"),
		"Trong phiếu nhập đang có 2 sách giống nhau",
		"ErrInventoryCheckNoteExistDuplicateBook",
	)
	ErrInventoryCheckNoteModifyQuantityIsInvalid = common.NewCustomError(
		errors.New("the qty after modification is invalid"),
		"Số lượng sau khi điều chỉnh không hợp lệ",
		"ErrInventoryCheckNoteModifyQuantityIsInvalid",
	)
	ErrInventoryCheckNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu kiểm kho mới"),
	)
	ErrInventoryCheckNoteViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem phiếu kiểm kho"),
	)
)
