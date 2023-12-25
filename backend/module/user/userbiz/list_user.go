package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type ListUserRepo interface {
	ListUser(
		ctx context.Context,
		userSearch string,
		filter *usermodel.Filter,
		paging *common.Paging,
	) ([]usermodel.User, error)
}

type listUserBiz struct {
	repo      ListUserRepo
	requester middleware.Requester
}

func NewListUserBiz(
	repo ListUserRepo,
	requester middleware.Requester) *listUserBiz {
	return &listUserBiz{repo: repo, requester: requester}
}

func (biz *listUserBiz) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging) ([]usermodel.User, error) {
	if !biz.requester.IsHasFeature(common.UserViewFeatureCode) {
		return nil, usermodel.ErrUserViewNoPermission
	}

	result, err := biz.repo.ListUser(ctx, biz.requester.GetUserId(), filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
