package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
)

type ImportNoteCreate struct {
	Id                string                                         `json:"-" gorm:"column:id;"`
	TotalPrice        float32                                        `json:"-" gorm:"column:totalPrice;"`
	SupplierId        string                                         `json:"supplierId" gorm:"column:supplierId"`
	CreateBy          string                                         `json:"-" gorm:"column:createBy;"`
	ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate `json:"importNoteDetails" gorm:"-"`
}

func (*ImportNoteCreate) TableName() string {
	return common.TableImportNote
}

func (data *ImportNoteCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrSupplierIdInvalid
	}
	if data.ImportNoteDetails == nil || len(data.ImportNoteDetails) == 0 {
		return ErrImportNoteDetailsEmpty
	}

	for _, importNoteDetail := range data.ImportNoteDetails {
		if err := importNoteDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
