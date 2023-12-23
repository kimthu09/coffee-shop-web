package productmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Product struct {
	Id           string `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Description  string `json:"description" gorm:"column:description;"`
	CookingGuide string `json:"cookingGuide" gorm:"column:cookingGuide"`
	IsActive     bool   `json:"isActive" gorm:"column:isActive;"`
}

var (
	ErrProductIdInvalid = common.NewCustomError(
		errors.New("id of product is invalid"),
		"Mã không hợp lệ",
		"ErrProductIdInvalid",
	)
	ErrProductNameEmpty = common.NewCustomError(
		errors.New("name of product is empty"),
		"Tên đang trống",
		"ErrProductNameEmpty",
	)
	ErrProductIsActiveEmpty = common.NewCustomError(
		errors.New("status of product is empty"),
		"Trạng thái đang trống",
		"ErrProductIsActiveEmpty",
	)
	ErrProductInactive = common.NewCustomError(
		errors.New("product has been inactive"),
		"Đã ngừng bán",
		"ErrProductInactive",
	)
)
