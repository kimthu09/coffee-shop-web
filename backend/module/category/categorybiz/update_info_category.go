package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

type UpdateInfoCategoryStore interface {
	FindCategory(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*categorymodel.Category, error)
	UpdateInfoCategory(
		ctx context.Context,
		id string,
		data *categorymodel.CategoryUpdateInfo) error
}

type updateInfoCategoryBiz struct {
	store     UpdateInfoCategoryStore
	requester middleware.Requester
}

func NewUpdateInfoCategoryBiz(
	store UpdateInfoCategoryStore,
	requester middleware.Requester) *updateInfoCategoryBiz {
	return &updateInfoCategoryBiz{store: store, requester: requester}
}

func (biz *updateInfoCategoryBiz) UpdateInfoCategory(
	ctx context.Context,
	id string,
	data *categorymodel.CategoryUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.CategoryUpdateInfoFeatureCode) {
		return categorymodel.ErrCategoryUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	_, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdateInfoCategory(ctx, id, data); err != nil {
		return err
	}

	return nil
}
