package recipedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
)

func (s *sqlStore) UpdateRecipeDetail(
	ctx context.Context,
	idRecipe string,
	idIngredient string,
	data *recipedetailmodel.RecipeDetailUpdate) error {
	db := s.db.Begin()

	if err := db.
		Where("idRecipe = ? and idIngredient = ?", idRecipe, idIngredient).
		Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
