package invoicedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type SimpleFood struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleFood) TableName() string {
	return common.TableFood
}

type InvoiceDetail struct {
	InvoiceId   string     `json:"invoiceId" gorm:"column:invoiceId;"`
	FoodId      string     `json:"-" gorm:"column:foodId;"`
	Food        SimpleFood `json:"food" gorm:"foreignKey:FoodId;references:Id"`
	SizeName    string     `json:"sizeName" gorm:"column:sizeName"`
	Amount      int        `json:"amount" gorm:"column:amount;"`
	UnitPrice   int        `json:"unitPrice" gorm:"column:unitPrice"`
	Description string     `json:"description" gorm:"column:description;"`
}

func (*InvoiceDetail) TableName() string {
	return common.TableInvoiceDetail
}

var (
	ErrInvoiceDetailFoodIdInvalid = common.NewCustomError(
		errors.New("id of food is invalid"),
		"Mã của sản phẩm không hợp lệ",
		"ErrInvoiceDetailFoodIdInvalid",
	)
	ErrInvoiceDetailSizeIdInvalid = common.NewCustomError(
		errors.New("id of size is invalid"),
		"Kích cỡ không hợp lệ",
		"ErrInvoiceDetailSizeIdInvalid",
	)
	ErrInvoiceDetailAmountIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"Số lượng sản phẩm cần thanh toán đang không phải số dương",
		"ErrInvoiceDetailAmountIsNotPositiveNumber",
	)
	ErrInvoiceDetailToppingsInvalid = common.NewCustomError(
		errors.New("list of topping is invalid"),
		"Danh sách các topping không hợp lệ",
		"ErrInvoiceDetailToppingsInvalid",
	)
	ErrInvoiceDetailFoodIsInactive = common.NewCustomError(
		errors.New("food is inactive"),
		"Sản phẩm đã ngừng bán",
		"ErrInvoiceDetailFoodIsInactive",
	)
	ErrInvoiceDetailExistToppingIsInactive = common.NewCustomError(
		errors.New("topping is inactive"),
		"Topping đã ngừng bán",
		"ErrInvoiceDetailExistToppingIsInactive",
	)
)
