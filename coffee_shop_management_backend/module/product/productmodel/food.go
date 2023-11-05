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
	ErrFoodIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of food is duplicate"),
	)
	ErrFoodNameDuplicate = common.ErrDuplicateKey(
		errors.New("name of food is duplicate"),
	)
	ErrFoodCategoryEmpty = common.NewCustomError(
		errors.New("category of food is empty"),
		"category of food is empty",
		"ErrFoodCategoryEmpty",
	)
	ErrFoodExistDuplicateCategory = common.NewCustomError(
		errors.New("exist duplicate category"),
		"exist duplicate category",
		"ErrFoodExistDuplicateCategory",
	)
	ErrFoodSizeEmpty = common.NewCustomError(
		errors.New("list size of food is empty"),
		"list size of food is empty",
		"ErrFoodSizeEmpty",
	)
	ErrFoodExistDuplicateSize = common.NewCustomError(
		errors.New("exist duplicate size"),
		"exist duplicate size",
		"ErrFoodExistDuplicateSize",
	)
	ErrFoodSizeIdInvalid = common.NewCustomError(
		errors.New("size id is invalid"),
		"size id is invalid",
		"ErrFoodSizeIdInvalid",
	)
)
