package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	db := s.db

	if err := db.Table(common.TableIngredient).
		Where("id = ?", id).
		Update("amount", gorm.Expr("amount + ?", data.Amount)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
