package categorystore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateAmountProductCategory(
	ctx context.Context,
	id string,
	data *categorymodel.CategoryUpdateAmountProduct) error {
	db := s.db.Begin()

	if err := db.Table(common.TableCategory).
		Where("id = ?", id).
		Update("amountProduct", gorm.Expr("amountProduct + ?", data.AmountProduct)).
		Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
