package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

type ListCategoryStorage interface {
	ListCategory(ctx context.Context,
		searchKey string,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListCategoryBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{
		store: store,
	}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging) ([]categorymodel.Category, error) {
	result, err := biz.store.ListCategory(
		ctx,
		filter.SearchKey,
		[]string{"id", "name"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
