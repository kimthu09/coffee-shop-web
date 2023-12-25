package recipestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"context"
)

func (s *sqlStore) DeleteRecipe(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Where(conditions).
		Delete(recipemodel.Recipe{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
