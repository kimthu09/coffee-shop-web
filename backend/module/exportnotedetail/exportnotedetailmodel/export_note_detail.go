package exportnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"errors"
)

type ExportNoteDetail struct {
	ExportNoteId string                           `json:"-" gorm:"column:exportNoteId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient" gorm:"foreignKey:IngredientId;references:Id"`
	AmountExport int                              `json:"amountExport" gorm:"column:amountExport;"`
}

func (*ExportNoteDetail) TableName() string {
	return common.TableExportNoteDetail
}

var (
	ErrExportDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrExportDetailIngredientIdInvalid",
	)
	ErrExportDetailAmountExportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount export is not positive number"),
		"Lượng muốn hủy đang không phải số dương",
		"ErrExportDetailAmountExportIsNotPositiveNumber",
	)
)
