package featurebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/feature/featuremodel"
	"context"
)

type ListFeatureStore interface {
	ListFeature(
		ctx context.Context,
	) ([]featuremodel.Feature, error)
}

type listFeatureBiz struct {
	store     ListFeatureStore
	requester middleware.Requester
}

func NewListFeatureBiz(
	store ListFeatureStore,
	requester middleware.Requester) *listFeatureBiz {
	return &listFeatureBiz{store: store, requester: requester}
}

func (biz *listFeatureBiz) ListFeature(
	ctx context.Context) ([]featuremodel.Feature, error) {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return nil, featuremodel.ErrFeatureViewNoPermission
	}

	result, err := biz.store.ListFeature(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
