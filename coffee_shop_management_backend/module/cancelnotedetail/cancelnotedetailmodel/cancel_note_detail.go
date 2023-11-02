package cancelnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type CancelNoteDetail struct {
	CancelNoteId string       `json:"cancelNoteId" gorm:"column:cancelNoteId;"`
	IngredientId string       `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string       `json:"expiryDate" gorm:"column:expiryDate;"`
	Reason       CancelReason `json:"reason" gorm:"column:reason;"`
	AmountCancel float32      `json:"amountCancel" gorm:"column:amountCancel;"`
}

func (*CancelNoteDetail) TableName() string {
	return common.TableCancelNoteDetail
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
	ErrCancelReasonEmpty = common.NewCustomError(
		errors.New("cancel reason is empty"),
		"cancel reason is empty",
		"ErrCancelReasonEmpty",
	)
	ErrAmountCancelIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount cancel is not positive number"),
		"amount cancel is not positive number",
		"ErrAmountCancelIsNotPositiveNumber",
	)
)
