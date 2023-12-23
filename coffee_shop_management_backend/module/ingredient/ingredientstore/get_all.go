package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.Ingredient, error) {
	var result []ingredientmodel.Ingredient
	db := s.db

	db = db.Table(common.TableIngredient)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
