package usermodel

import "coffee_shop_management_backend/common"

type UserUpdateRole struct {
	RoleId string `json:"roleId" gorm:"column:roleId;"`
}

func (*UserUpdateRole) TableName() string {
	return common.TableUser
}

func (data *UserUpdateRole) Validate() error {
	if !common.ValidateNotNilId(&data.RoleId) {
		return ErrUserRoleInvalid
	}
	return nil
}
