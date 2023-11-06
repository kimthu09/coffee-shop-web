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
		userId string,
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
	id string,
	data *usermodel.UserUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.UserUpdateStatusFeatureCode) {
		return usermodel.ErrUserUpdateStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateStatusUser(ctx, id, data); err != nil {
		return err
	}

	return nil
}
