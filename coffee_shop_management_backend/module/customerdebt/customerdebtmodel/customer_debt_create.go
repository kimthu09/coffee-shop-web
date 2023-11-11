package customerdebtmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
)

type CustomerDebtCreate struct {
	Id         string         `json:"-" gorm:"column:id;"`
	CustomerId string         `json:"customerId" gorm:"column:customerId;"`
	Amount     float32        `json:"amount" gorm:"column:amount;"`
	AmountLeft float32        `json:"-" gorm:"column:amountLeft;"`
	DebtType   *enum.DebtType `json:"type" gorm:"column:type;"`
	CreateBy   string         `json:"-" gorm:"column:createBy;"`
}

func (*CustomerDebtCreate) TableName() string {
	return common.TableCustomerDebt
}
