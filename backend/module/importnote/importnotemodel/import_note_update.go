package importnotemodel

import (
	"coffee_shop_management_backend/common"
)

type ImportNoteUpdate struct {
	ClosedBy   string            `json:"-" gorm:"column:closedBy;"`
	Id         string            `json:"-" gorm:"-"`
	SupplierId string            `json:"-" gorm:"-"`
	TotalPrice int               `json:"-" gorm:"-"`
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
