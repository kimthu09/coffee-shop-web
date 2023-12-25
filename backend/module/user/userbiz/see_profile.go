package userbiz

import (
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type SeeProfileRepo interface {
	SeeUserDetail(
		ctx context.Context,
		userId string) (*usermodel.User, error)
}

type seeProfileBiz struct {
	repo      SeeProfileRepo
	requester middleware.Requester
}

func NewSeeProfileBiz(
	repo SeeUserDetailRepo,
	requester middleware.Requester) *seeProfileBiz {
	return &seeProfileBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeProfileBiz) SeeProfile(
	ctx context.Context) (*usermodel.User, error) {
	user, err := biz.repo.SeeUserDetail(
		ctx, biz.requester.GetUserId())
	if err != nil {
		return nil, err
	}

	return user, nil
}
