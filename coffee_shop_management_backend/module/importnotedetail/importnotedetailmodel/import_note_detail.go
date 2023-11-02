package importnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type ImportNoteDetail struct {
	ImportNoteId string  `json:"importNoteId" gorm:"column:importNoteId;"`
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string  `json:"expiryDate" gorm:"column:expiryDate;"`
	AmountImport float32 `json:"amountImport" gorm:"column:amountImport"`
}

func (*ImportNoteDetail) TableName() string {
	return common.TableImportNoteDetail
}

var (
	ErrIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrIngredientIdInvalid",
	)
	ErrExpiryDateInvalid = common.NewCustomError(
		errors.New("expiry date is invalid"),
		"expiry date is invalid",
		"ErrExpiryDateInvalid",
	)
	ErrAmountImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"amount import is not positive number",
		"ErrAmountImportIsNotPositiveNumber",
	)
)
