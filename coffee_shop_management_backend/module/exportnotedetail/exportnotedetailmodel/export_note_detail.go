package exportnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type ExportNoteDetail struct {
	ExportNoteId string  `json:"exportNoteId" gorm:"column:exportNoteId;"`
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string  `json:"expiryDate" gorm:"column:expiryDate;"`
	AmountExport float32 `json:"amountExport" gorm:"column:amountExport"`
}

func (*ExportNoteDetail) TableName() string {
	return common.TableExportNoteDetail
}

var (
	ErrExportDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrExportDetailIngredientIdInvalid",
	)
	ErrExportDetailExpiryDateInvalid = common.NewCustomError(
		errors.New("expiry date is invalid"),
		"expiry date is invalid",
		"ErrExportDetailExpiryDateInvalid",
	)
	ErrExportDetailAmountExportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount export is not positive number"),
		"amount export is not positive number",
		"ErrExportDetailAmountExportIsNotPositiveNumber",
	)
)
