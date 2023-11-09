package invoicemodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type SimpleCustomer struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleCustomer) TableName() string {
	return common.TableCustomer
}

type Invoice struct {
	Id             string         `json:"id" gorm:"column:id;"`
	CustomerId     string         `json:"-" gorm:"column:customerId;"`
	Customer       SimpleCustomer `json:"customer" gorm:"foreignKey:CustomerId;references:Id"`
	TotalPrice     float32        `json:"totalPrice" gorm:"column:totalPrice;"`
	AmountReceived float32        `json:"amountReceived" gorm:"column:amountReceived"`
	AmountDebt     float32        `json:"amountDebt" gorm:"column:amountDebt"`
	CreateBy       string         `json:"createBy" gorm:"column:createBy;"`
	CreateAt       *time.Time     `json:"createAt" gorm:"column:createAt;"`
}

func (*Invoice) TableName() string {
	return common.TableInvoice
}

var (
	ErrInvoiceCustomerIdInvalid = common.NewCustomError(
		errors.New("id of customer is invalid"),
		"id of customer is invalid",
		"ErrInvoiceCustomerIdInvalid",
	)
	ErrInvoiceAmountReceivedIsNegativeNumber = common.NewCustomError(
		errors.New("amount received is negative number"),
		"amount received is negative number",
		"ErrInvoiceAmountReceivedIsNegativeNumber",
	)
	ErrInvoiceAmountDebtIsNegativeNumber = common.NewCustomError(
		errors.New("amount debt is negative number"),
		"amount debt is negative number",
		"ErrInvoiceAmountDebtIsNegativeNumber",
	)
	ErrInvoiceNotHaveCustomerForDebt = common.NewCustomError(
		errors.New("customer is required for this invoice"),
		"customer is required for this invoice",
		"ErrInvoiceNotHaveCustomerForDebt",
	)
	ErrInvoiceDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"list import note details are empty",
		"ErrInvoiceDetailsEmpty",
	)
	ErrInvoiceCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create invoice"),
	)
	ErrInvoiceViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view invoice"),
	)
)