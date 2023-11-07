package cancelnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"errors"
)

type CancelNoteDetail struct {
	CancelNoteId string                           `json:"cancelNoteId" gorm:"column:cancelNoteId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient" gorm:"foreignKey:IngredientId;references:Id"`
	ExpiryDate   string                           `json:"expiryDate" gorm:"column:expiryDate;"`
	Reason       CancelReason                     `json:"reason" gorm:"column:reason;"`
	AmountCancel float32                          `json:"amountCancel" gorm:"column:amountCancel;"`
}

func (*CancelNoteDetail) TableName() string {
	return common.TableCancelNoteDetail
}

var (
	ErrCancelDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrCancelDetailIngredientIdInvalid",
	)
	ErrCancelDetailExpiryDateInvalid = common.NewCustomError(
		errors.New("expiry date is invalid"),
		"expiry date is invalid",
		"ErrCancelDetailExpiryDateInvalid",
	)
	ErrCancelDetailCancelReasonEmpty = common.NewCustomError(
		errors.New("cancel reason is empty"),
		"cancel reason is empty",
		"ErrCancelDetailCancelReasonEmpty",
	)
	ErrCancelDetailAmountCancelIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount cancel is not positive number"),
		"amount cancel is not positive number",
		"ErrCancelDetailAmountCancelIsNotPositiveNumber",
	)
	ErrCancelDetailViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view cancel note"),
	)
)
