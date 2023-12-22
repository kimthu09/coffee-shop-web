package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.SimpleIngredient, error) {
	var result []ingredientmodel.SimpleIngredient
	db := s.db

	db = db.Table(common.TableIngredient)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
