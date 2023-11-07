package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

type CreateCategoryStorage interface {
	CreateCategory(
		ctx context.Context,
		data *categorymodel.CategoryCreate,
	) error
}

type createCategoryBiz struct {
	gen       generator.IdGenerator
	store     CreateCategoryStorage
	requester middleware.Requester
}

func NewCreateCategoryBiz(
	generator generator.IdGenerator,
	store CreateCategoryStorage,
	requester middleware.Requester) *createCategoryBiz {
	return &createCategoryBiz{gen: generator, store: store, requester: requester}
}

func (biz *createCategoryBiz) CreateCategory(
	ctx context.Context,
	data *categorymodel.CategoryCreate) error {
	if !biz.requester.IsHasFeature(common.CategoryCreateFeatureCode) {
		return categorymodel.ErrCategoryCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := biz.gen.GenerateId()
	if err != nil {
		return err
	}

	data.Id = idAddress

	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}
