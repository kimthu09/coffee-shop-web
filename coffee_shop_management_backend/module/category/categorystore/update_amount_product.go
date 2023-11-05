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
	db := s.db

	if err := db.Table(common.TableCategory).
		Where("id = ?", id).
		Update("amountProduct", gorm.Expr(
			"amountProduct + ?", data.AmountProduct,
		)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
