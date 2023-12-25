package importnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"errors"
)

type ImportNoteDetail struct {
	ImportNoteId string                           `json:"importNoteId" gorm:"column:importNoteId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient" gorm:"foreignKey:IngredientId;references:Id"`
	Price        float32                          `json:"price" gorm:"column:price"`
	TotalUnit    float32                          `json:"totalUnit" gorm:"column:totalUnit"`
	AmountImport int                              `json:"amountImport" gorm:"column:amountImport"`
}

func (*ImportNoteDetail) TableName() string {
	return common.TableImportNoteDetail
}

var (
	ErrImportDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrImportDetailIngredientIdInvalid",
	)
	ErrImportDetailPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"Giá của nguyên vật liệu đang là số âm",
		"ErrImportDetailPriceIsNegativeNumber",
	)
	ErrImportDetailAmountImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"Lượng muốn nhập đang không phải số dương",
		"ErrImportDetailAmountImportIsNotPositiveNumber",
	)
)
