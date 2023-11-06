package exportnotedetailmodel

import "coffee_shop_management_backend/common"

type ExportNoteDetailCreate struct {
	ExportNoteId string  `json:"-" gorm:"column:exportNoteId;"`
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string  `json:"expiryDate" gorm:"column:expiryDate;"`
	AmountExport float32 `json:"amountExport" gorm:"column:amountExport"`
}

func (*ExportNoteDetailCreate) TableName() string {
	return common.TableExportNoteDetail
}

func (data *ExportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrExportDetailIngredientIdInvalid
	}
	if !common.ValidateDateString(data.ExpiryDate) {
		return ErrExportDetailExpiryDateInvalid
	}
	if common.ValidateNotPositiveNumber(data.AmountExport) {
		return ErrExportDetailAmountExportIsNotPositiveNumber
	}
	return nil
}
