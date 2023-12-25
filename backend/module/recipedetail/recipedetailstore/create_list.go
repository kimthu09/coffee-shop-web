package recipedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
)

func (s *sqlStore) CreateListRecipeDetail(
	ctx context.Context,
	data []recipedetailmodel.RecipeDetailCreate) error {
	db := s.db
	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
