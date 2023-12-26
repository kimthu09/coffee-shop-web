package invoicedetailmodel

import (
	"coffee_shop_management_backend/common"
)

type InvoiceDetailCreate struct {
	InvoiceId   string                 `json:"-" gorm:"column:invoiceId;"`
	FoodId      string                 `json:"foodId" gorm:"column:foodId;"`
	FoodName    string                 `json:"foodName" gorm:"column:foodName;"`
	SizeId      string                 `json:"sizeId" gorm:"-"`
	SizeName    string                 `json:"sizeName" gorm:"column:sizeName"`
	Toppings    *InvoiceDetailToppings `json:"toppings" gorm:"column:toppings;"`
	Amount      int                    `json:"amount" gorm:"column:amount;"`
	UnitPrice   int                    `json:"-" gorm:"column:unitPrice"`
	Description string                 `json:"description" gorm:"column:description;"`
}

func (*InvoiceDetailCreate) TableName() string {
	return common.TableInvoiceDetail
}

func (data *InvoiceDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.FoodId) {
		return ErrInvoiceDetailFoodIdInvalid
	}
	if !common.ValidateNotNilId(&data.SizeId) {
		return ErrInvoiceDetailSizeIdInvalid
	}
	if common.ValidateNotPositiveNumberInt(data.Amount) {
		return ErrInvoiceDetailAmountIsNotPositiveNumber
	}
	if data.Toppings == nil {
		return ErrInvoiceDetailToppingsInvalid
	}
	return nil
}
