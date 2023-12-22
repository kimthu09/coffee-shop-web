package customermodel

import "coffee_shop_management_backend/common"

type CustomerUpdatePoint struct {
	Amount *float32 `json:"amount" gorm:"-"`
}

func (*CustomerUpdatePoint) TableName() string {
	return common.TableCustomer
}
