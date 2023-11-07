package ingredientdetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

type ListIngredientDetailStore interface {
	ListIngredientDetail(
		ctx context.Context,
		ingredientId string,
		filter *ingredientdetailmodel.Filter,
		paging *common.Paging,
	) ([]ingredientdetailmodel.IngredientDetail, error)
}

type listIngredientDetail struct {
	store     ListIngredientDetailStore
	requester middleware.Requester
}

func NewListIngredientDetailByIdBiz(
	store ListIngredientDetailStore,
	requester middleware.Requester) *listIngredientDetail {
	return &listIngredientDetail{store: store, requester: requester}
}

func (biz *listIngredientDetail) ListIngredientDetail(
	ctx context.Context,
	ingredientId string,
	filter *ingredientdetailmodel.Filter,
	paging *common.Paging) ([]ingredientdetailmodel.IngredientDetail, error) {
	if !biz.requester.IsHasFeature(common.IngredientViewFeatureCode) {
		return nil, ingredientdetailmodel.ErrIngredientDetailViewNoPermission
	}

	result, err := biz.store.ListIngredientDetail(
		ctx, ingredientId, filter, paging,
	)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.TableCategory, err)
	}

	return result, nil
}
