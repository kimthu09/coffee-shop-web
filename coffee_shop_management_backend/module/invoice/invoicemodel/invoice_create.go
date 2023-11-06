package invoicemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
)

type InvoiceCreate struct {
	Id             string                                   `json:"-" gorm:"column:id;"`
	CustomerId     *string                                  `json:"customerId" gorm:"column:customerId;"`
	TotalPrice     float32                                  `json:"-" gorm:"column:totalPrice;"`
	AmountReceived float32                                  `json:"amountReceived" gorm:"column:amountReceived"`
	AmountDebt     float32                                  `json:"-" gorm:"column:amountDebt"`
	CreateBy       string                                   `json:"-" gorm:"column:createBy;"`
	InvoiceDetails []invoicedetailmodel.InvoiceDetailCreate `json:"details" gorm:"-"`
}

func (*InvoiceCreate) TableName() string {
	return common.TableInvoice
}

func (data *InvoiceCreate) Validate() *common.AppError {
	if !common.ValidateId(data.CustomerId) {
		return ErrInvoiceCustomerIdInvalid
	}
	if common.ValidateNegativeNumber(&data.AmountReceived) {
		return ErrInvoiceAmountReceivedIsNegativeNumber
	}
	if data.InvoiceDetails == nil || len(data.InvoiceDetails) == 0 {
		return ErrInvoiceDetailsEmpty
	}

	for _, invoiceDetail := range data.InvoiceDetails {
		if err := invoiceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
