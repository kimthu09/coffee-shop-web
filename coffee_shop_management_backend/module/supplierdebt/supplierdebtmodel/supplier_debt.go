package supplierdebtmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type SupplierDebt struct {
	Id               string            `json:"id" gorm:"column:id;"`
	IdSupplier       string            `json:"idSupplier" gorm:"column:idSupplier;"`
	Amount           float32           `json:"amount" gorm:"column:amount;"`
	AmountLeft       float32           `json:"amountLeft" gorm:"column:amountLeft;"`
	SupplierDebtType *SupplierDebtType `json:"type" gorm:"column:type;"`
	CreateBy         string            `json:"createBy" gorm:"column:createBy;"`
	CreateAt         *time.Time        `json:"createAt" gorm:"column:createAt;"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}

var (
	ErrIdSupplierInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrIdSupplierInvalid",
	)
	ErrAmountIsNotNegativeNumber = common.NewCustomError(
		errors.New("amount is not negative number"),
		"amount is not negative number",
		"ErrAmountIsNotNegativeNumber",
	)
	ErrSupplierDebtTypeEmpty = common.NewCustomError(
		errors.New("supplier debt type is empty"),
		"supplier debt type is empty",
		"ErrSupplierDebtTypeEmpty",
	)
)
