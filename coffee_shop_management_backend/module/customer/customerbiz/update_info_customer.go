package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
)

type UpdateInfoCustomerRepo interface {
	CheckExist(
		ctx context.Context,
		customerId string,
	) error
	UpdateCustomerInfo(
		ctx context.Context,
		customerId string,
		data *customermodel.CustomerUpdateInfo,
	) error
}

type updateInfoCustomerBiz struct {
	repo      UpdateInfoCustomerRepo
	requester middleware.Requester
}

func NewUpdateInfoCustomerBiz(
	repo UpdateInfoCustomerRepo,
	requester middleware.Requester) *updateInfoCustomerBiz {
	return &updateInfoCustomerBiz{repo: repo, requester: requester}
}

func (biz *updateInfoCustomerBiz) UpdateInfoCustomer(
	ctx context.Context,
	id string,
	data *customermodel.CustomerUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.CustomerUpdateInfoFeatureCode) {
		return customermodel.ErrCustomerUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckExist(ctx, id); err != nil {
		return err
	}

	if err := biz.repo.UpdateCustomerInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
