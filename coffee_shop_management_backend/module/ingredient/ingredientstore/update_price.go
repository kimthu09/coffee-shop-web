package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) UpdatePriceIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdatePrice) error {
	db := s.db

	if err := db.Table(common.TableIngredient).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
