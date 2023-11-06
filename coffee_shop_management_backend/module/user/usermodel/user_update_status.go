package usermodel

import "coffee_shop_management_backend/common"

type UserUpdateStatus struct {
	IsActive *bool `json:"isActive" gorm:"column:isActive;"`
}

func (*UserUpdateStatus) TableName() string {
	return common.TableUser
}

func (data *UserUpdateStatus) Validate() error {
	if data.IsActive == nil {
		return ErrUserStatusEmpty
	}
	return nil
}
