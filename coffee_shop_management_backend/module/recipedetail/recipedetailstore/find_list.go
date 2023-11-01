package recipedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindListRecipeDetail(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*[]recipedetailmodel.RecipeDetail, error) {
	var data []recipedetailmodel.RecipeDetail
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Table(common.TableRecipeDetail).
		Where(conditions).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
