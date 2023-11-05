package importnotemodel

import (
	"coffee_shop_management_backend/common"
)

type ImportNoteUpdate struct {
	CloseBy    string            `json:"-" gorm:"column:closeBy;"`
	Id         string            `json:"-" gorm:"-"`
	SupplierId string            `json:"-" gorm:"-"`
	TotalPrice float32           `json:"-" gorm:"-"`
	Status     *ImportNoteStatus `json:"status" gorm:"column:status;"`
}

func (*ImportNoteUpdate) TableName() string {
	return common.TableImportNote
}

func (data *ImportNoteUpdate) Validate() *common.AppError {
	if data.Status == nil {
		return ErrImportNoteStatusEmpty
	}
	if *data.Status == InProgress {
		return ErrImportNoteStatusInvalid
	}
	return nil
}
