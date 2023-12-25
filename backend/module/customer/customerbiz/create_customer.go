package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
)

type CreateCustomerRepo interface {
	CreateCustomer(
		ctx context.Context,
		data *customermodel.CustomerCreate,
	) error
}

type createCustomerBiz struct {
	gen       generator.IdGenerator
	repo      CreateCustomerRepo
	requester middleware.Requester
}

func NewCreateCustomerBiz(
	gen generator.IdGenerator,
	repo CreateCustomerRepo,
	requester middleware.Requester) *createCustomerBiz {
	return &createCustomerBiz{gen: gen, repo: repo, requester: requester}
}

func (biz *createCustomerBiz) CreateCustomer(
	ctx context.Context,
	data *customermodel.CustomerCreate) error {
	if !biz.requester.IsHasFeature(common.CustomerCreateFeatureCode) {
		return customermodel.ErrCustomerCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleCustomerId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateCustomer(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleCustomerId(gen generator.IdGenerator, data *customermodel.CustomerCreate) error {
	idAddress, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress
	return nil
}
