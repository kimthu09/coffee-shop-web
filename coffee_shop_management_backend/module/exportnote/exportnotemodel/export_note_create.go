package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
)

type ExportNoteCreate struct {
	Id                *string                                        `json:"id" gorm:"column:id;"`
	TotalPrice        float32                                        `json:"-" gorm:"column:totalPrice;"`
	CreateBy          string                                         `json:"-" gorm:"column:createBy;"`
	ExportNoteDetails []exportnotedetailmodel.ExportNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ExportNoteCreate) TableName() string {
	return common.TableExportNote
}

func (data *ExportNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrExportNoteIdInvalid
	}
	if data.ExportNoteDetails == nil || len(data.ExportNoteDetails) == 0 {
		return ErrExportNoteDetailsEmpty
	}

	for _, v := range data.ExportNoteDetails {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
