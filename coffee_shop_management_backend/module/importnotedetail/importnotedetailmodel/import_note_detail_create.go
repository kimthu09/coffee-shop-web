package importnotedetailmodel

import (
	"coffee_shop_management_backend/common"
)

type ImportNoteDetailCreate struct {
	ImportNoteId   string  `json:"-" gorm:"column:importNoteId;"`
	IngredientId   string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate     string  `json:"expiryDate" gorm:"column:expiryDate;"`
	Price          float32 `json:"price" gorm:"column:price"`
	IsReplacePrice bool    `json:"isReplacePrice" gorm:"-"`
	AmountImport   float32 `json:"amountImport" json:"amountImport" gorm:"column:amountImport"`
}

func (*ImportNoteDetailCreate) TableName() string {
	return common.TableImportNoteDetail
}

func (data *ImportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrImportDetailIngredientIdInvalid
	}
	if !common.ValidateDateString(data.ExpiryDate) {
		return ErrImportDetailExpiryDateInvalid
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrImportDetailPriceIsNegativeNumber
	}
	if common.ValidateNotPositiveNumber(data.AmountImport) {
		return ErrImportDetailAmountImportIsNotPositiveNumber
	}
	return nil
}
