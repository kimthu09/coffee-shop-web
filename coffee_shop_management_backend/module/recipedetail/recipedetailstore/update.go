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
	db := s.db

	if err := db.
		Where("recipeId = ? and ingredientId = ?", idRecipe, idIngredient).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
