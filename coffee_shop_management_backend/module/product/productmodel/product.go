package productmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Product struct {
	Id           string `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Description  string `json:"description" gorm:"column:description;"`
	ProductSizes *Sizes `json:"sizes" gorm:"column:sizes;"`
	IsActive     bool   `json:"isActive" gorm:"column:isActive;"`
}

var (
	ErrIdInvalid = common.NewCustomError(
		errors.New("id of product is empty"),
		"id of product is empty",
		"ErrIdInvalid",
	)
	ErrNameEmpty = common.NewCustomError(
		errors.New("name of product is empty"),
		"name of product is empty",
		"ErrNameEmpty",
	)
	ErrSizeNotExist = common.NewCustomError(
		errors.New("size of product has not been included"),
		"size of product has not been included",
		"ErrSizeNotExist",
	)
	ErrProductInactive = common.NewCustomError(
		errors.New("product has been inactive"),
		"product has been inactive",
		"ErrProductInactive",
	)
)
