package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type SeeUserDetailRepo interface {
	SeeUserDetail(
		ctx context.Context,
		userId string) (*usermodel.User, error)
}

type seeUserDetailBiz struct {
	repo      SeeUserDetailRepo
	requester middleware.Requester
}

func NewSeeUserDetailBiz(
	repo SeeUserDetailRepo,
	requester middleware.Requester) *seeUserDetailBiz {
	return &seeUserDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeUserDetailBiz) SeeUserDetail(
	ctx context.Context,
	userId string) (*usermodel.User, error) {
	if !biz.requester.IsHasFeature(common.UserViewFeatureCode) {
		return nil, usermodel.ErrUserViewNoPermission
	}

	user, err := biz.repo.SeeUserDetail(
		ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
