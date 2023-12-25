package invoicemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
)

type InvoiceCreate struct {
	Id                  string                                   `json:"-" gorm:"column:id;"`
	CustomerId          *string                                  `json:"customerId" gorm:"column:customerId;"`
	Customer            SimpleCustomer                           `json:"customer" gorm:"-"`
	TotalPrice          int                                      `json:"-" gorm:"column:totalPrice;"`
	IsUsePoint          bool                                     `json:"isUsePoint" gorm:"-"`
	AmountReceived      int                                      `json:"-" gorm:"column:amountReceived"`
	AmountPriceUsePoint int                                      `json:"-" gorm:"column:amountPriceUsePoint"`
	CreatedBy           string                                   `json:"-" gorm:"column:createdBy;"`
	InvoiceDetails      []invoicedetailmodel.InvoiceDetailCreate `json:"details" gorm:"-"`
	MapIngredient       map[string]int                           `json:"-" gorm:"-"`
	ShopName            string                                   `json:"-" gorm:"-"`
	ShopPhone           string                                   `json:"-" gorm:"-"`
	ShopAddress         string                                   `json:"-" gorm:"-"`
	ShopPassWifi        string                                   `json:"-" gorm:"-"`
}

func (*InvoiceCreate) TableName() string {
	return common.TableInvoice
}

func (data *InvoiceCreate) Validate() *common.AppError {
	if !common.ValidateId(data.CustomerId) {
		return ErrInvoiceCustomerIdInvalid
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
