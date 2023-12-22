package usermodel

import "coffee_shop_management_backend/common"

type SimpleUser struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Email string `json:"email" gorm:"column:email;"`
}

func (*SimpleUser) TableName() string {
	return common.TableUser
}
