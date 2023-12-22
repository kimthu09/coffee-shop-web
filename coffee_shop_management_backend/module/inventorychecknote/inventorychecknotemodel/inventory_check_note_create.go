package inventorychecknotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

type InventoryCheckNoteCreate struct {
	Id                *string                                                        `json:"id" gorm:"column:id;" example:""`
	AmountDifferent   int                                                            `json:"-" gorm:"column:amountDifferent;"`
	AmountAfterAdjust int                                                            `json:"-" gorm:"column:amountAfterAdjust;"`
	CreatedBy         string                                                         `json:"-" gorm:"column:createdBy;"`
	Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate `json:"details" gorm:"-"`
}

func (*InventoryCheckNoteCreate) TableName() string {
	return common.TableInventoryCheckNote
}

func (data *InventoryCheckNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrInventoryCheckNoteIdInvalid
	}
	if data.Details == nil || len(data.Details) == 0 {
		return ErrInventoryCheckNoteDetailEmpty
	}
	mapExits := make(map[string]int)
	for _, detail := range data.Details {
		if err := detail.Validate(); err != nil {
			return err
		}
		mapExits[detail.IngredientId]++
		if mapExits[detail.IngredientId] >= 2 {
			return ErrInventoryCheckNoteExistDuplicateBook
		}
	}
	return nil
}
