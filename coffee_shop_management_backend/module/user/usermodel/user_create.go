package usermodel

import "coffee_shop_management_backend/common"

type UserCreate struct {
	Id       string `json:"-" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt;"`
	Role     string `json:"role" gorm:"column:role;"`
}

func (*UserCreate) TableName() string {
	return common.TableUser
}
