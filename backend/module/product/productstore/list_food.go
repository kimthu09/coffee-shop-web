package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListFood(
	ctx context.Context,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]productmodel.Food, error) {
	var result []productmodel.Food
	db := s.db

	db = db.Table(common.TableFood)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.IsActive != nil {
			if *filter.IsActive {
				db = db.Where("isActive = ?", true)
			} else {
				db = db.Where("isActive = ?", false)
			}
		}
	}
}
