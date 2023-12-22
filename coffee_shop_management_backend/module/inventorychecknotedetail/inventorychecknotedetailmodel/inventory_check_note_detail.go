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
	Final                int                              `json:"final" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"Mã của sách không hợp lệ",
		"ErrInventoryCheckDetailBookIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"Số lượng chỉnh sửa không hợp lệ",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
