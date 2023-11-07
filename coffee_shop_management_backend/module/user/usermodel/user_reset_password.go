package usermodel

import "coffee_shop_management_backend/common"

type UserResetPassword struct {
	UserSenderPass string `json:"userSenderPass" gorm:"-"`
}

func (*UserResetPassword) TableName() string {
	return common.TableUser
}

func (data *UserResetPassword) Validate() error {
	if !common.ValidatePassword(&data.UserSenderPass) {
		return ErrUserSenderPassInvalid
	}
	return nil
}
