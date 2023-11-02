package usermodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type User struct {
	Id       string `json:"id" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"-" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt;"`
	Role     string `json:"role" gorm:"column:role;"`
	IsActive bool   `json:"isActive" gorm:"column:isActive;"`
}

func (u *User) GetUserId() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (*User) TableName() string {
	return common.TableUser
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
