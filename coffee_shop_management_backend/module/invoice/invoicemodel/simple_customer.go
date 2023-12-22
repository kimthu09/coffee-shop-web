package invoicemodel

import "coffee_shop_management_backend/common"

type SimpleCustomer struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Phone string `json:"phone" gorm:"column:phone;"`
}

func (*SimpleCustomer) TableName() string {
	return common.TableCustomer
}
