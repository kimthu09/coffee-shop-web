package suppliermodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Supplier struct {
	Id    string  `json:"id" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Email string  `json:"email" gorm:"column:email;"`
	Phone string  `json:"phone" gorm:"column:phone;"`
	Debt  float32 `json:"debt" gorm:"column:debt;"`
}

func (*Supplier) TableName() string {
	return common.TableSupplier
}

var (
	ErrSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrSupplierIdInvalid",
	)
	ErrSupplierNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"name of supplier is empty",
		"ErrSupplierNameEmpty",
	)
	ErrSupplierPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"phone of supplier is invalid",
		"ErrSupplierPhoneInvalid",
	)
	ErrSupplierEmailInvalid = common.NewCustomError(
		errors.New("email of supplier is invalid"),
		"email of supplier is invalid",
		"ErrSupplierEmailInvalid",
	)
	ErrDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"debt pay is not exist",
		"ErrDebtPayNotExist",
	)
	ErrDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"debt pay is invalid",
		"ErrDebtPayIsInvalid",
	)
	ErrSupplierIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of supplier is duplicate"),
	)
	ErrSupplierPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("phone of supplier is duplicate"),
	)
	ErrSupplierCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create supplier"),
	)
	ErrSupplierPayNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to pay supplier"),
	)
	ErrSupplierUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info supplier"),
	)
)
