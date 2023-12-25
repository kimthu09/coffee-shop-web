package recipestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"context"
)

func (s *sqlStore) CreateRecipe(
	ctx context.Context,
	data *recipemodel.RecipeCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
