package productmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Food struct {
	*Product `json:",inline"`
}

func (*Food) TableName() string {
	return common.TableFood
}

var (
	ErrCategoryEmpty = common.NewCustomError(
		errors.New("category of food is empty"),
		"category of food is empty",
		"ErrCategoryEmpty",
	)
	ErrExistDuplicateCategory = common.NewCustomError(
		errors.New("exist duplicate category"),
		"exist duplicate category",
		"ErrExistDuplicateCategory",
	)
	ErrSizeEmpty = common.NewCustomError(
		errors.New("list size of food is empty"),
		"list size of food is empty",
		"ErrSizeEmpty",
	)
	ErrExistDuplicateSize = common.NewCustomError(
		errors.New("exist duplicate size"),
		"exist duplicate size",
		"ErrExistDuplicateSize",
	)
	ErrSizeIdInvalid = common.NewCustomError(
		errors.New("size id is invalid"),
		"size id is invalid",
		"ErrSizeIdInvalid",
	)
)
