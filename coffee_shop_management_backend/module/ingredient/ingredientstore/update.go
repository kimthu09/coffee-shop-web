package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdate) error {
	db := s.db.Begin()

	if err := db.Table(common.TableIngredient).
		Where("id = ?", id).
		Update("totalAmount", gorm.Expr("totalAmount + ?", data.Amount)).
		Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
