package importnotedetailmodel

import (
	"coffee_shop_management_backend/common"
)

type ImportNoteDetailCreate struct {
	ImportNoteId   string  `json:"-" gorm:"column:importNoteId;"`
	IngredientId   string  `json:"ingredientId" gorm:"column:ingredientId;"`
	Price          float32 `json:"price" gorm:"column:price"`
	IsReplacePrice bool    `json:"isReplacePrice" gorm:"-"`
	AmountImport   int     `json:"amountImport" json:"amountImport" gorm:"column:amountImport"`
	TotalUnit      float32 `json:"-" gorm:"column:totalUnit"`
}

func (*ImportNoteDetailCreate) TableName() string {
	return common.TableImportNoteDetail
}

func (data *ImportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrImportDetailIngredientIdInvalid
	}
	if common.ValidateNegativeNumberFloat(data.Price) {
		return ErrImportDetailPriceIsNegativeNumber
	}
	if common.ValidateNotPositiveNumberInt(data.AmountImport) {
		return ErrImportDetailAmountImportIsNotPositiveNumber
	}
	return nil
}

func (data *ImportNoteDetailCreate) Round() {
	common.CustomRound(&data.Price)
}
