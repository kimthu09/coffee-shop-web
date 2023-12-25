package usermodel

import "coffee_shop_management_backend/common"

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"-"`
}

func (*UserLogin) TableName() string {
	return common.TableUser
}
