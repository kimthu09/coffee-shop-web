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
		"Id của nhà cung cấp không hợp lệ",
		"ErrSupplierIdInvalid",
	)
	ErrSupplierNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"Tên của nhà cung cấp đang trống",
		"ErrSupplierNameEmpty",
	)
	ErrSupplierPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"Số điện thoại của nhà cung cấp không hợp lệ",
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
	ErrDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"Số tiền trả nợ cho nhà cung cấp đang trống",
		"ErrDebtPayNotExist",
	)
	ErrDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"Số tiền trả nợ cho nhà cung cấp không hợp lệ",
		"ErrDebtPayIsInvalid",
	)
	ErrSupplierIdDuplicate = common.ErrDuplicateKey(
		errors.New("Đã tồn tại nhà cung cấp trong hệ thống"),
	)
	ErrSupplierPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("Đã tồn tại nhà cung cấp có số điện thoại này trong hệ thống"),
	)
	ErrSupplierCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo nhà cung cấp mới"),
	)
	ErrSupplierPayNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền trả nợ nhà cung cấp"),
	)
	ErrSupplierUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền thay đổi thông tin nhà cung cấp"),
	)
	ErrSupplierViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem nhà cung cấp"),
	)
)
