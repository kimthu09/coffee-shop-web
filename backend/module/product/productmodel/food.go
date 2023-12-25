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
		errors.New("Sản phẩm đã tồn tại"),
	)
	ErrFoodNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của sản phẩm đã tồn tại"),
	)
	ErrFoodProductInfoEmpty = common.NewCustomError(
		errors.New("product info of food is empty"),
		"Thông tin sản phẩm đang thiếu",
		"ErrFoodProductInfoEmpty",
	)
	ErrFoodCategoryEmpty = common.NewCustomError(
		errors.New("category of food is empty"),
		"Danh mục của sản phẩm đang trống",
		"ErrFoodCategoryEmpty",
	)
	ErrFoodExistDuplicateCategory = common.NewCustomError(
		errors.New("exist duplicate category"),
		"Danh mục của sản phẩm đang bị trùng",
		"ErrFoodExistDuplicateCategory",
	)
	ErrFoodSizeEmpty = common.NewCustomError(
		errors.New("list size of food is empty"),
		"Danh sách kích cỡ của sản phẩm đang trống",
		"ErrFoodSizeEmpty",
	)
	ErrFoodExistDuplicateSize = common.NewCustomError(
		errors.New("exist duplicate size"),
		"Kích cỡ của sản phẩm đang bị trùng",
		"ErrFoodExistDuplicateSize",
	)
	ErrFoodSizeIdInvalid = common.NewCustomError(
		errors.New("size id is invalid"),
		"Kích cỡ không hợp lệ",
		"ErrFoodSizeIdInvalid",
	)
	ErrFoodCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo sản phẩm mới"),
	)
	ErrFoodUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin sản phẩm"),
	)
	ErrFoodChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa trạng thái sản phẩm"),
	)
	ErrFoodViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem sản phẩm"),
	)
)
