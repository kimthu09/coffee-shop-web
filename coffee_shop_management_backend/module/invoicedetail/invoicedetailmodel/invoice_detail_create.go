package invoicedetailmodel

import (
	"coffee_shop_management_backend/common"
)

type InvoiceDetailCreate struct {
	InvoiceId   string                 `json:"-" gorm:"column:invoiceId;"`
	FoodId      string                 `json:"foodId" gorm:"column:foodId;"`
	SizeId      string                 `json:"sizeId" gorm:"-"`
	SizeName    string                 `json:"-" gorm:"column:sizeName"`
	Toppings    *InvoiceDetailToppings `json:"toppings" gorm:"column:toppings;"`
	Amount      float32                `json:"amount" gorm:"column:amount;"`
	UnitPrice   float32                `json:"-" gorm:"column:unitPrice"`
	Description string                 `json:"description" gorm:"column:description;"`
}

func (*InvoiceDetailCreate) TableName() string {
	return common.TableInvoiceDetail
}

func (data *InvoiceDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.FoodId) {
		return ErrInvoiceDetailIngredientIdInvalid
	}
	if !common.ValidateNotNilId(&data.SizeId) {
		return ErrInvoiceDetailSizeIdInvalid
	}
	if common.ValidateNotPositiveNumber(data.Amount) {
		return ErrInvoiceDetailAmountIsNotPositiveNumber
	}
	if data.Toppings == nil {
		return ErrInvoiceDetailToppingsInvalid
	}
	return nil
}
