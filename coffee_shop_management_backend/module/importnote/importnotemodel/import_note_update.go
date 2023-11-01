package importnotemodel

import (
	"coffee_shop_management_backend/common"
)

type ImportNoteUpdate struct {
	CloseBy string            `json:"-" gorm:"column:closeBy;"`
	Status  *ImportNoteStatus `json:"status" gorm:"column:status;"`
}

func (*ImportNoteUpdate) TableName() string {
	return common.TableImportNote
}

func (data *ImportNoteUpdate) Validate() *common.AppError {
	if data.Status == nil {
		return ErrImportStatusEmpty
	}
	if *data.Status == InProgress {
		return ErrImportStatusInvalid
	}
	return nil
}
