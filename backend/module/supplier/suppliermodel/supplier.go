package suppliermodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Supplier struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Email string `json:"email" gorm:"column:email;"`
	Phone string `json:"phone" gorm:"column:phone;"`
	Debt  int    `json:"debt" gorm:"column:debt;"`
}

func (*Supplier) TableName() string {
	return common.TableSupplier
}

var (
	ErrSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"Mã của nhà cung cấp không hợp lệ",
		"ErrSupplierIdInvalid",
	)
	ErrSupplierNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"Tên của nhà cung cấp đang trống",
		"ErrSupplierNameEmpty",
	)
	ErrSupplierPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"Số điện thọai của nhà cung cấp không hợp lệ",
		"ErrSupplierPhoneInvalid",
	)
	ErrSupplierEmailInvalid = common.NewCustomError(
		errors.New("email of supplier is invalid"),
		"Email của nhà cung cấp không hợp lệ",
		"ErrSupplierEmailInvalid",
	)
	ErrSupplierInitDebtInvalid = common.NewCustomError(
		errors.New("init debt of supplier is invalid"),
		"Nợ ban đầu của nhà cung cấp không hợp lệ",
		"ErrSupplierInitDebtInvalid",
	)
	ErrSupplierDebtIdInvalid = common.NewCustomError(
		errors.New("id of supplier debt is invalid"),
		"Mã phiếu chi không hợp lệ",
		"ErrSupplierDebtIdInvalid",
	)
	ErrSupplierDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"Số tiền chi đang trống",
		"ErrSupplierDebtPayNotExist",
	)
	ErrSupplierDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"Số tiền chi không hợp lệ",
		"ErrSupplierDebtPayIsInvalid",
	)
	ErrSupplierIdDuplicate = common.ErrDuplicateKey(
		errors.New("Nhà cung cấp đã tồn tại"),
	)
	ErrSupplierPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("Số điện thoại nhà cung cấp đã tồn tại"),
	)
	ErrSupplierCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo nhà cung cấp mới"),
	)
	ErrSupplierPayNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu chi mới"),
	)
	ErrSupplierUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin nhà cung cấp"),
	)
	ErrSupplierViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem nhà cung cấp"),
	)
)
