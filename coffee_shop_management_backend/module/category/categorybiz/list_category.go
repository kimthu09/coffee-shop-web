package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

type ListCategoryStore interface {
	ListCategory(
		ctx context.Context,
		filter *categorymodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store     ListCategoryStore
	requester middleware.Requester
}

func NewListCategoryBiz(
	store ListCategoryStore,
	requester middleware.Requester) *listCategoryBiz {
	return &listCategoryBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *listCategoryBiz) ListCategory(
	ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging) ([]categorymodel.Category, error) {
	if !biz.requester.IsHasFeature(common.CategoryViewFeatureCode) {
		return nil, categorymodel.ErrCategoryViewNoPermission
	}

	result, err := biz.store.ListCategory(
		ctx,
		filter,
		[]string{"id", "name"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
