package featuremodel

import "coffee_shop_management_backend/common"

type Feature struct {
	Id          string `json:"id" gorm:"column:id;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (*Feature) TableName() string {
	return common.TableFeature
}
