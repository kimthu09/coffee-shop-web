package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
)

type ListCustomerRepo interface {
	ListCustomer(
		ctx context.Context,
		filter *customermodel.Filter,
		paging *common.Paging,
	) ([]customermodel.Customer, error)
}

type listCustomerBiz struct {
	repo      ListCustomerRepo
	requester middleware.Requester
}

func NewListCustomerBiz(
	repo ListCustomerRepo,
	requester middleware.Requester) *listCustomerBiz {
	return &listCustomerBiz{repo: repo, requester: requester}
}

func (biz *listCustomerBiz) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	paging *common.Paging) ([]customermodel.Customer, error) {
	if !biz.requester.IsHasFeature(common.CustomerViewFeatureCode) {
		return nil, customermodel.ErrCustomerViewNoPermission
	}

	result, err := biz.repo.ListCustomer(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
