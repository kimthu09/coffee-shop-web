package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) ListTopping(
	ctx context.Context,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]productmodel.Topping, error) {
	var result []productmodel.Topping
	db := s.db

	db = db.Table(common.TableTopping)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("name").
		Preload("Recipe.Details.Ingredient").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
