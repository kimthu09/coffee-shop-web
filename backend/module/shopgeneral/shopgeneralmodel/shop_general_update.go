package shopgeneralmodel

import "coffee_shop_management_backend/common"

type ShopGeneralUpdate struct {
	Name                   *string  `json:"name" gorm:"column:name;"`
	Email                  *string  `json:"email" gorm:"column:email;"`
	Phone                  *string  `json:"phone" gorm:"column:phone;"`
	Address                *string  `json:"address" gorm:"column:address;"`
	WifiPass               *string  `json:"wifiPass" gorm:"column:wifiPass;"`
	AccumulatePointPercent *float32 `json:"accumulatePointPercent" gorm:"column:accumulatePointPercent"`
	UsePointPercent        *float32 `json:"usePointPercent" gorm:"column:usePointPercent"`
}

func (*ShopGeneralUpdate) TableName() string {
	return common.TableShopGeneral
}

func (data *ShopGeneralUpdate) Validate() *common.AppError {
	if data.Email != nil && *data.Email != "" && !common.ValidateEmail(*data.Email) {
		return ErrEmailInvalid
	}
	if data.Phone != nil && *data.Phone != "" && !common.ValidatePhone(*data.Phone) {
		return ErrPhoneInvalid
	}
	if data.AccumulatePointPercent != nil && common.ValidateNegativeNumberFloat(*data.AccumulatePointPercent) {
		return ErrAccumulatePointPercentInvalid
	}
	if data.UsePointPercent != nil && common.ValidateNegativeNumberFloat(*data.UsePointPercent) {
		return ErrUsePointPercentInvalid
	}
	return nil
}
