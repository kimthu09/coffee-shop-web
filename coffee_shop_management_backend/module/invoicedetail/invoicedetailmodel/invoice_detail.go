package invoicedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type InvoiceDetail struct {
	InvoiceId   string  `json:"invoiceId" gorm:"column:invoiceId;"`
	FoodId      string  `json:"foodId" gorm:"column:foodId;"`
	SizeName    string  `json:"sizeName" gorm:"column:sizeName"`
	Amount      float32 `json:"amount" gorm:"column:amount;"`
	UnitPrice   float32 `json:"unitPrice" gorm:"column:unitPrice"`
	Description string  `json:"description" gorm:"column:description;"`
}

func (*InvoiceDetail) TableName() string {
	return common.TableInvoiceDetail
}

var (
	ErrInvoiceDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrInvoiceDetailIngredientIdInvalid",
	)
	ErrInvoiceDetailSizeIdInvalid = common.NewCustomError(
		errors.New("id of size is invalid"),
		"id of size is invalid",
		"ErrInvoiceDetailSizeIdInvalid",
	)
	ErrInvoiceDetailAmountIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"amount import is not positive number",
		"ErrInvoiceDetailAmountIsNotPositiveNumber",
	)
	ErrInvoiceDetailToppingsInvalid = common.NewCustomError(
		errors.New("list of topping is invalid"),
		"list of topping is invalid",
		"ErrInvoiceDetailToppingsInvalid",
	)
	ErrInvoiceDetailFoodIsInactive = common.NewCustomError(
		errors.New("food is inactive"),
		"food is inactive",
		"ErrInvoiceDetailFoodIsInactive",
	)
	ErrInvoiceDetailExistToppingIsInactive = common.NewCustomError(
		errors.New("topping is inactive"),
		"topping is inactive",
		"ErrInvoiceDetailExistToppingIsInactive",
	)
)
