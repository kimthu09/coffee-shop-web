package usermodel

import "coffee_shop_management_backend/common"

type UserUpdatePassword struct {
	OldPassword string `json:"oldPassword" gorm:"-"`
	NewPassword string `json:"newPassword" gorm:"-"`
}

func (*UserUpdatePassword) TableName() string {
	return common.TableUser
}

func (data *UserUpdatePassword) Validate() error {
	if !common.ValidatePassword(&data.OldPassword) {
		return ErrUserUpdatedPassInvalid
	}
	if !common.ValidatePassword(&data.NewPassword) {
		return ErrUserUpdatedPassInvalid
	}
	return nil
}
