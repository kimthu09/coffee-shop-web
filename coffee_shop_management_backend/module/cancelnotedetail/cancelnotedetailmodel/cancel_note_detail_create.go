package cancelnotedetailmodel

import (
	"coffee_shop_management_backend/common"
)

type CancelNoteDetailCreate struct {
	CancelNoteId string        `json:"-" gorm:"column:cancelNoteId;"`
	IngredientId string        `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string        `json:"expiryDate" gorm:"column:expiryDate;"`
	Reason       *CancelReason `json:"reason" gorm:"column:reason;"`
	AmountCancel float32       `json:"amountCancel" gorm:"column:amountCancel;"`
}

func (*CancelNoteDetailCreate) TableName() string {
	return common.TableCancelNoteDetail
}

func (data *CancelNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrCancelDetailIngredientIdInvalid
	}
	if !common.ValidateDateString(data.ExpiryDate) {
		return ErrCancelDetailExpiryDateInvalid
	}
	if data.Reason == nil {
		return ErrCancelDetailCancelReasonEmpty
	}
	if common.ValidateNotPositiveNumber(data.AmountCancel) {
		return ErrCancelDetailAmountCancelIsNotPositiveNumber
	}
	return nil
}
