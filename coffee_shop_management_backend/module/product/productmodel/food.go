package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"errors"
)

type Food struct {
	*Product       `json:",inline"`
	FoodCategories Categories `json:"categories" gorm:"foreignkey:foodId;association_foreignkey:id"`
	FoodSizes      Sizes      `json:"sizes" gorm:"foreignkey:foodId;association_foreignkey:id"`
}

func (*Food) TableName() string {
	return common.TableFood
}

type Categories []categoryfoodmodel.CategoryFood

func (*Categories) TableName() string {
	return common.TableCategoryFood
}

type Sizes []sizefoodmodel.SizeFood

func (*Sizes) TableName() string {
	return common.TableSizeFood
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
	ErrFoodCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create food"),
	)
	ErrFoodUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info food"),
	)
	ErrFoodChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to change status food"),
	)
	ErrFoodViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view food"),
	)
)
