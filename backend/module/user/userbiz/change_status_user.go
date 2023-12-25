package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type ChangeStatusUserRepo interface {
	UpdateStatusUser(
		ctx context.Context,
		data *usermodel.UserUpdateStatus,
	) error
}

type changeStatusUserBiz struct {
	repo      ChangeStatusUserRepo
	requester middleware.Requester
}

func NewChangeStatusUserBiz(
	repo ChangeStatusUserRepo,
	requester middleware.Requester) *changeStatusUserBiz {
	return &changeStatusUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusUserBiz) ChangeStatusUser(
	ctx context.Context,
	data []usermodel.UserUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.UserUpdateStatusFeatureCode) {
		return usermodel.ErrUserUpdateStatusNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	for _, v := range data {
		if err := biz.repo.UpdateStatusUser(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
