package ingredientdetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateIngredientDetail(
	ctx context.Context,
	ingredientId string,
	expiryDate string,
	data *ingredientdetailmodel.IngredientDetailUpdate) error {
	db := s.db

	if err := db.
		Table(common.TableIngredientDetail).
		Where("ingredientId = ? and expiryDate = ?", ingredientId, expiryDate).
		Update("amount", gorm.Expr("amount + ?", data.Amount)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrRecordNotFound()
		}
		return common.ErrDB(err)
	}

	return nil
}
