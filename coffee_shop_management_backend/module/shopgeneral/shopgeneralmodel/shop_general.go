package shopgeneralmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type ShopGeneral struct {
	Name                   string  `json:"name" gorm:"column:name;"`
	Email                  string  `json:"email" gorm:"column:email;"`
	Phone                  string  `json:"phone" gorm:"column:phone;"`
	Address                string  `json:"address" gorm:"column:address;"`
	WifiPass               string  `json:"wifiPass" gorm:"column:wifiPass;"`
	AccumulatePointPercent float32 `json:"accumulatePointPercent" gorm:"column:accumulatePointPercent"`
	UsePointPercent        float32 `json:"usePointPercent" gorm:"column:usePointPercent"`
}

func (*ShopGeneral) TableName() string {
	return common.TableShopGeneral
}

var (
	ErrEmailInvalid = common.NewCustomError(
		errors.New("email of shop is invalid"),
		"email of shop is invalid",
		"ErrEmailInvalid",
	)
	ErrPhoneInvalid = common.NewCustomError(
		errors.New("phone of shop is invalid"),
		"phone of shop is invalid",
		"ErrPhoneInvalid",
	)
	ErrAccumulatePointPercentInvalid = common.NewCustomError(
		errors.New("accumulate point percent of shop is invalid"),
		"accumulate point percent of shop is invalid",
		"ErrAccumulatePointPercentInvalid",
	)
	ErrUsePointPercentInvalid = common.NewCustomError(
		errors.New("use point percent of shop is invalid"),
		"use point percent of shop is invalid",
		"ErrUsePointPercentInvalid",
	)
	ErrGeneralShopViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view shop general"),
	)
	ErrGeneralShopUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update shop general"),
	)
)
