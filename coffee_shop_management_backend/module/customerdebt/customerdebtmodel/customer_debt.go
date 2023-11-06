package customerdebtmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"errors"
	"time"
)

type CustomerDebt struct {
	Id         string         `json:"id" gorm:"column:id;"`
	IdCustomer string         `json:"idCustomer" gorm:"column:idCustomer;"`
	Amount     float32        `json:"amount" gorm:"column:amount;"`
	AmountLeft float32        `json:"amountLeft" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy   string         `json:"createBy" gorm:"column:createBy;"`
	CreateAt   *time.Time     `json:"createAt" gorm:"column:createAt;"`
}

func (*CustomerDebt) TableName() string {
	return common.TableCustomerDebt
}

var (
	ErrIdSupplierInvalid = common.NewCustomError(
		errors.New("id of customer is invalid"),
		"id of customer is invalid",
		"ErrIdSupplierInvalid",
	)
	ErrAmountIsNotNegativeNumber = common.NewCustomError(
		errors.New("amount is not negative number"),
		"amount is not negative number",
		"ErrAmountIsNotNegativeNumber",
	)
	ErrDebtTypeEmpty = common.NewCustomError(
		errors.New("debt type is empty"),
		"debt type is empty",
		"ErrDebtTypeEmpty",
	)
)
