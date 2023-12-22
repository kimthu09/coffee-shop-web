package supplierdebtmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
)

type SupplierDebtCreate struct {
	Id         string         `json:"-" gorm:"column:id;"`
	SupplierId string         `json:"supplierId" gorm:"column:supplierId;"`
	Amount     int            `json:"amount" gorm:"column:amount;"`
	AmountLeft int            `json:"-" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreatedBy  string         `json:"-" gorm:"column:createdBy;"`
}

func (*SupplierDebtCreate) TableName() string {
	return common.TableSupplierDebt
}
