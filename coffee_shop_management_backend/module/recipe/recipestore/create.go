package recipestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"context"
)

func (s *sqlStore) CreateRecipe(
	ctx context.Context,
	data *recipemodel.RecipeCreate) error {
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
