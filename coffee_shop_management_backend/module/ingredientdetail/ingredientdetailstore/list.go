package ingredientdetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListIngredientDetail(
	ctx context.Context,
	ingredientId string,
	filter *ingredientdetailmodel.Filter,
	paging *common.Paging) ([]ingredientdetailmodel.IngredientDetail, error) {
	var result []ingredientdetailmodel.IngredientDetail
	db := s.db

	db = db.Table(common.TableIngredientDetail)

	handleFilter(db, filter)

	db = db.Where("ingredientId = ?", ingredientId)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *ingredientdetailmodel.Filter) {
	if filter != nil {
		if !filter.IsGetEmpty {
			db = db.Where("amount != 0")
		}
	}
}

func handlePaging(db *gorm.DB, paging *common.Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	return db, nil
}
