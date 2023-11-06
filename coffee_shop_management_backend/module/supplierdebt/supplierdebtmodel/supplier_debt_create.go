package supplierdebtmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
)

type SupplierDebtCreate struct {
	Id         string         `json:"-" gorm:"column:id;"`
	IdSupplier string         `json:"idSupplier" gorm:"column:idSupplier;"`
	Amount     float32        `json:"amount" gorm:"column:amount;"`
	AmountLeft float32        `json:"-" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy   string         `json:"-" gorm:"column:createBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return common.TableSupplierDebt
}

func (data *SupplierDebtCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IdSupplier) {
		return ErrIdSupplierInvalid
	}
	if common.ValidateNotNegativeNumber(data.Amount) {
		return ErrAmountIsNotNegativeNumber
	}
	if data.DebtType == nil {
		return ErrDebtTypeEmpty
	}
	return nil
}
