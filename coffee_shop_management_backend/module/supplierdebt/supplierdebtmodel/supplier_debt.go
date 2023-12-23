package supplierdebtmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"time"
)

type SupplierDebt struct {
	Id         string         `json:"id" gorm:"column:id;"`
	SupplierId string         `json:"supplierId" gorm:"column:supplierId;"`
	Amount     int            `json:"amount" gorm:"column:amount;"`
	AmountLeft int            `json:"amountLeft" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreatedBy  string         `json:"createdBy" gorm:"column:createdBy;"`
	CreatedAt  *time.Time     `json:"createdAt" gorm:"column:createdAt;"`
}

func (*SupplierDebt) TableName() string {
	return common.TableSupplierDebt
}
