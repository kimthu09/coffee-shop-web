package inventorychecknotedetailmodel

import "coffee_shop_management_backend/common"

type InventoryCheckNoteDetailCreate struct {
	InventoryCheckNoteId string `json:"-" gorm:"column:inventoryCheckNoteId;"`
	IngredientId         string `json:"ingredientId" gorm:"column:ingredientId;"`
	Initial              int    `json:"-" gorm:"column:initial;"`
	Difference           int    `json:"difference" gorm:"column:difference;"`
	Final                int    `json:"-" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetailCreate) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

func (data *InventoryCheckNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrInventoryCheckDetailIngredientIdInvalid
	}
	if data.Difference == 0 {
		return ErrInventoryCheckDifferenceIsInvalid
	}
	return nil
}
