package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"errors"
)

type Ingredient struct {
	Id          string            `json:"id" gorm:"column:id;"`
	Name        string            `json:"name" gorm:"column:name;"`
	Amount      int               `json:"amount" gorm:"column:amount;"`
	MeasureType *enum.MeasureType `json:"measureType" gorm:"column:measureType;"`
	Price       float32           `json:"price" gorm:"column:price;"`
}

func (*Ingredient) TableName() string {
	return common.TableIngredient
}

var (
	ErrIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrIngredientIdInvalid",
	)
	ErrIngredientNameEmpty = common.NewCustomError(
		errors.New("name of ingredient is empty"),
		"Tên của nguyên vật liệu đang trống",
		"ErrIngredientNameEmpty",
	)
	ErrIngredientPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"Giá của nguyên vật liệu đang là số âm",
		"ErrIngredientPriceIsNegativeNumber",
	)
	ErrIngredientMeasureTypeEmpty = common.NewCustomError(
		errors.New("measure type of ingredient is empty"),
		"Loại đo lường của nguyên vật liệu đang trống",
		"ErrIngredientMeasureTypeEmpty",
	)
	ErrIngredientIdDuplicate = common.ErrDuplicateKey(
		errors.New("Nguyên vật liệu đã tồn tại"),
	)
	ErrIngredientNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của nguyên vật liệu đã tồn tại"),
	)
	ErrIngredientCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền thêm nguyên vật liệu mới"),
	)
	ErrIngredientViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem nguyên vật liệu"),
	)
)
