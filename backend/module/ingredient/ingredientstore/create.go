package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) CreateIngredient(
	ctx context.Context,
	data *ingredientmodel.IngredientCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "name"); key {
			case "PRIMARY":
				return ingredientmodel.ErrIngredientIdDuplicate
			case "name":
				return ingredientmodel.ErrIngredientNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
