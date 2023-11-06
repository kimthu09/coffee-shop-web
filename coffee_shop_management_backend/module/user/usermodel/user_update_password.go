package usermodel

import "coffee_shop_management_backend/common"

type UserUpdatePassword struct {
	Password string `json:"password" gorm:"-"`
}

func (*UserUpdatePassword) TableName() string {
	return common.TableUser
}

func (data *UserUpdatePassword) Validate() error {
	if !common.ValidatePassword(&data.Password) {
		return ErrUserUpdatedPassInvalid
	}
	return nil
}
