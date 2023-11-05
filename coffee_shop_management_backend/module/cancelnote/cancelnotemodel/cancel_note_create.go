package cancelnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
)

type CancelNoteCreate struct {
	Id                      *string                                        `json:"id" gorm:"column:id;"`
	TotalPrice              float32                                        `json:"-" gorm:"column:totalPrice;"`
	CreateBy                string                                         `json:"-" gorm:"column:createBy;"`
	CancelNoteCreateDetails []cancelnotedetailmodel.CancelNoteDetailCreate `json:"details" gorm:"-"`
}

func (*CancelNoteCreate) TableName() string {
	return common.TableCancelNote
}

func (data *CancelNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrCancelNoteIdInvalid
	}
	if data.CancelNoteCreateDetails == nil || len(data.CancelNoteCreateDetails) == 0 {
		return ErrCancelNoteDetailsEmpty
	}
	for _, v := range data.CancelNoteCreateDetails {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
