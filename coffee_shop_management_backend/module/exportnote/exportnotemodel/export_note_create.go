package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
)

type ExportNoteCreate struct {
	Id                *string                                        `json:"id" gorm:"column:id;"`
	CreatedBy         string                                         `json:"-" gorm:"column:createdBy;"`
	Reason            *ExportReason                                  `json:"reason" gorm:"column:reason;"`
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
	if data.Reason == nil {
		return ErrExportNoteReasonEmpty
	}
	mapExist := make(map[string]int)
	for _, v := range data.ExportNoteDetails {
		if err := v.Validate(); err != nil {
			return err
		}
		mapExist[v.IngredientId]++
		if mapExist[v.IngredientId] >= 2 {
			return ErrExportNoteExistDuplicateIngredient
		}
	}
	return nil
}
