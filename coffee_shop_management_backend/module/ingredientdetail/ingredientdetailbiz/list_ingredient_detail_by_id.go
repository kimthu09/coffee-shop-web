package ingredientdetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

type ListIngredientDetailByIdStorage interface {
	ListIngredientDetailById(
		ctx context.Context,
		condition map[string]interface{},
		filter *ingredientdetailmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ingredientdetailmodel.IngredientDetail, error)
}

type listIngredientDetailByIdBiz struct {
	store ListIngredientDetailByIdStorage
}

func NewListIngredientDetailByIdBiz(store ListIngredientDetailByIdStorage) *listIngredientDetailByIdBiz {
	return &listIngredientDetailByIdBiz{store: store}
}

func (biz *listIngredientDetailByIdBiz) ListIngredientDetailById(
	ctx context.Context,
	ingredientId string,
	filter *ingredientdetailmodel.Filter,
	paging *common.Paging) ([]ingredientdetailmodel.IngredientDetail, error) {
	result, err := biz.store.ListIngredientDetailById(
		ctx, map[string]interface{}{"ingredientId": ingredientId}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.TableCategory, err)
	}

	return result, nil
}
