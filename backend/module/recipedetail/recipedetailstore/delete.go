package recipedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
)

func (s *sqlStore) DeleteRecipeDetail(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Where(conditions).
		Delete(recipedetailmodel.RecipeDetail{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
