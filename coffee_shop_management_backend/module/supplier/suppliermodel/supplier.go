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
	ErrIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrIdInvalid",
	)
	ErrNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"name of supplier is empty",
		"ErrNameEmpty",
	)
	ErrPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"phone of supplier is invalid",
		"ErrPhoneInvalid",
	)
	ErrEmailInvalid = common.NewCustomError(
		errors.New("email of supplier is invalid"),
		"email of supplier is invalid",
		"ErrEmailInvalid",
	)
	ErrDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"debt pay is not exist",
		"ErrDebtPayNotExist",
	)
)
