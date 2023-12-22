package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
)

type ImportNoteCreate struct {
	Id                *string                                        `json:"id" gorm:"column:id;"`
	TotalPrice        int                                            `json:"-" gorm:"column:totalPrice;"`
	SupplierId        string                                         `json:"supplierId" gorm:"column:supplierId"`
	CreatedBy         string                                         `json:"-" gorm:"column:createdBy;"`
	ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ImportNoteCreate) TableName() string {
	return common.TableImportNote
}

func (data *ImportNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrImportNoteIdInvalid
	}
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrImportNoteSupplierIdInvalid
	}
	if data.ImportNoteDetails == nil || len(data.ImportNoteDetails) == 0 {
		return ErrImportNoteDetailsEmpty
	}

	mapIngredientUpdatePriceTimes := make(map[string]int)
	for _, importNoteDetail := range data.ImportNoteDetails {
		if err := importNoteDetail.Validate(); err != nil {
			return err
		}
		if importNoteDetail.IsReplacePrice {
			mapIngredientUpdatePriceTimes[importNoteDetail.IngredientId]++
			if mapIngredientUpdatePriceTimes[importNoteDetail.IngredientId] > 1 {
				return ErrImportNoteHasSameIngredientBothUpdatePrice
			}
		}
	}
	return nil
}

func (data *ImportNoteCreate) Round() {
	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].Round()
	}
}
