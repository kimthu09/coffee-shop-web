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
	ErrArrCancelNoteDetailsEmpty = common.NewCustomError(
		errors.New("the list cancel note details are empty"),
		"the list cancel note details are empty",
		"ErrArrCancelNoteDetailsEmpty",
	)
)
