package usermodel

import "coffee_shop_management_backend/common"

type UserUpdateInfo struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Phone   *string `json:"phone" gorm:"column:phone;"`
	Address *string `json:"address" gorm:"column:address;"`
}

func (*UserUpdateInfo) TableName() string {
	return common.TableUser
}

func (data *UserUpdateInfo) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrUserNameEmpty
	}
	if data.Phone != nil && len(*data.Phone) != 0 && !common.ValidatePhone(*data.Phone) {
		return ErrUserPhoneInvalid
	}
	return nil
}
