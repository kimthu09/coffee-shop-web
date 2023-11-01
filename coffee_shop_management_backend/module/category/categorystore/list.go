package categorystore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"strings"
)

func (s *sqlStore) ListCategory(ctx context.Context,
	searchKey string,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]categorymodel.Category, error) {
	var result []categorymodel.Category
	db := s.db

	db = db.Table(common.TableCategory)

	conditions := make([]string, len(propertiesContainSearchKey))
	args := make([]string, len(propertiesContainSearchKey))

	for i, prop := range propertiesContainSearchKey {
		conditions[i] = prop + " LIKE ?"
		args[i] = "%" + searchKey + "%"
	}

	whereClause := strings.Join(conditions, " OR ")

	db.Where(whereClause, args)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	if err := db.
		Limit(int(paging.Limit)).
		Order("create_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
