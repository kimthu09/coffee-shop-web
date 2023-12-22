package categorymodel

import "coffee_shop_management_backend/common"

type CategoryUpdateAmountProduct struct {
	AmountProduct *int `json:"amountProduct" gorm:"column:amountProduct;"`
}

func (*CategoryUpdateAmountProduct) TableName() string {
	return common.TableCategory
}
