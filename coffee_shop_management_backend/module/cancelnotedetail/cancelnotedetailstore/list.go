package cancelnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
)

func (s *sqlStore) ListCancelNote(
	ctx context.Context,
	condition map[string]interface{},
	filter *cancelnotedetailmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]cancelnotedetailmodel.CancelNoteDetail, error) {
	var result []cancelnotedetailmodel.CancelNoteDetail

	db := s.db.Table(common.TableIngredientDetail).Where(condition)

	//if filterValue := filter; filterValue != nil {
	//	if filterValue.IsGetEmptyIngredientDetails {
	//		db = db.Where("amount > 0")
	//	}
	//}

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
		Order("ingredientId desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
