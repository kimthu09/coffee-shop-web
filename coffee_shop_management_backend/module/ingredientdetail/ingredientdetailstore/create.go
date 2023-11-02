package ingredientdetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

func (s *sqlStore) CreateIngredientDetail(
	ctx context.Context,
	data *ingredientdetailmodel.IngredientDetailCreate) error {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
