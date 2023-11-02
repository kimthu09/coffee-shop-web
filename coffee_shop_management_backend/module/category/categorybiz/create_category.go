package categorybiz

import (
	"coffee_shop_management_backend/common"
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
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(
	ctx context.Context,
	data *categorymodel.CategoryCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	errStorage := biz.store.CreateCategory(ctx, data)

	return errStorage
}
