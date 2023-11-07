package categorymodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type SimpleCategory struct {
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}

type Category struct {
	Id            string `json:"id" gorm:"column:id;"`
	Name          string `json:"name" gorm:"column:name;"`
	Description   string `json:"description" gorm:"column:description;"`
	AmountProduct int    `json:"amountProduct" gorm:"column:amountProduct;"`
}

func (*Category) TableName() string {
	return common.TableCategory
}

var (
	ErrCategoryNameEmpty = common.NewCustomError(
		errors.New("name of category is empty"),
		"name of category is empty",
		"ErrCategoryNameEmpty",
	)
	ErrCategoryAmountProductCategoryNotExist = common.NewCustomError(
		errors.New("amount product of category is empty"),
		"amount product of category is empty",
		"ErrCategoryAmountProductCategoryNotExist",
	)
	ErrCategoryIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of category is duplicate"),
	)
	ErrCategoryNameDuplicate = common.ErrDuplicateKey(
		errors.New("name of category is duplicate"),
	)
	ErrCategoryCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create category"),
	)
	ErrCategoryUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info category"),
	)
	ErrCategoryViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view category"),
	)
)
