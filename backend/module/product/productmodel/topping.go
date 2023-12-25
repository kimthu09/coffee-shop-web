package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"errors"
)

type Topping struct {
	*Product `json:",inline"`
	Cost     int                 `json:"cost" gorm:"column:cost;"`
	Price    int                 `json:"price" gorm:"column:price;"`
	RecipeId string              `json:"-" gorm:"column:recipeId;"`
	Recipe   *recipemodel.Recipe `json:"recipe" gorm:"foreignkey:recipeId"`
}

func (*Topping) TableName() string {
	return common.TableTopping
}

var (
	ErrToppingProductInfoEmpty = common.NewCustomError(
		errors.New("product info of topping is empty"),
		"Thông tin topping đang thiếu",
		"ErrToppingProductInfoEmpty",
	)
	ErrToppingCostIsNegativeNumber = common.NewCustomError(
		errors.New("cost is negative number"),
		"Giá bán của topping đang là số âm",
		"ErrSizeFoodCostIsNegativeNumber",
	)
	ErrToppingPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price is negative number"),
		"Giá gốc của topping đang là số âm",
		"ErrSizeFoodPriceIsNegativeNumber",
	)
	ErrToppingIdDuplicate = common.ErrDuplicateKey(
		errors.New("Topping đã tồn tại"),
	)
	ErrToppingNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của topping đã tồn tại"),
	)
	ErrToppingRecipeEmpty = common.NewCustomError(
		errors.New("recipe is empty"),
		"Công thức nấu ăn đang trống",
		"ErrToppingRecipeEmpty",
	)
	ErrToppingCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo topping mới"),
	)
	ErrToppingUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin topping"),
	)
	ErrToppingChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa trạng thái topping"),
	)
	ErrToppingViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem topping"),
	)
)
