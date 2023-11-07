package cancelnotemodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type CancelNote struct {
	Id         string     `json:"id" gorm:"column:id;"`
	TotalPrice float32    `json:"totalPrice" gorm:"column:totalPrice;"`
	CreateAt   *time.Time `json:"createAt" gorm:"column:createAt;"`
	CreateBy   string     `json:"createBy" gorm:"column:createBy;"`
}

func (*CancelNote) TableName() string {
	return common.TableCancelNote
}

var (
	ErrCancelNoteIdInvalid = common.NewCustomError(
		errors.New("id of cancel note is invalid"),
		"id of cancel note is invalid",
		"ErrCancelNoteIdInvalid",
	)
	ErrCancelNoteDetailsEmpty = common.NewCustomError(
		errors.New("the list cancel note details are empty"),
		"the list cancel note details are empty",
		"ErrCancelNoteDetailsEmpty",
	)
	ErrCancelNoteAmountCancelIsOverTheStock = common.NewCustomError(
		errors.New("amount cancel is over stock"),
		"amount cancel is over stock",
		"ErrCancelNoteAmountCancelIsOverTheStock",
	)
	ErrCancelNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create cancel note"),
	)
	ErrCancelNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view cancel note"),
	)
)
