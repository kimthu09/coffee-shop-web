package exportnotedetailmodel

import "coffee_shop_management_backend/common"

type ExportNoteDetailCreate struct {
	ExportNoteId string `json:"-" gorm:"column:exportNoteId;"`
	IngredientId string `json:"ingredientId" gorm:"column:ingredientId;"`
	AmountExport int    `json:"amountExport" gorm:"column:amountExport"`
}

func (*ExportNoteDetailCreate) TableName() string {
	return common.TableExportNoteDetail
}

func (data *ExportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrExportDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumber(data.AmountExport) {
		return ErrExportDetailAmountExportIsNotPositiveNumber
	}
	return nil
}
