package inventorychecknotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"errors"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string                           `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;"`
	IngredientId         string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient           ingredientmodel.SimpleIngredient `json:"ingredient"`
	Initial              int                              `json:"initial" gorm:"column:initial;"`
	Difference           int                              `json:"difference" gorm:"column:difference;"`
	Final                int                              `json:"difference" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrInventoryCheckDetailIngredientIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"Số lượng chỉnh sửa không hợp lệ",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
