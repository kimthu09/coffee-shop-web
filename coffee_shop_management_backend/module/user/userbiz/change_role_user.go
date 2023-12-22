package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type UpdateRoleUserRepo interface {
	CheckUserStatusPermission(
		ctx context.Context,
		userId string,
	) error
	UpdateRoleUser(
		ctx context.Context,
		userId string,
		data *usermodel.UserUpdateRole) error
}

type changeRoleUserBiz struct {
	repo      UpdateRoleUserRepo
	requester middleware.Requester
}

func NewChangeRoleUserBiz(
	repo UpdateRoleUserRepo,
	requester middleware.Requester) *changeRoleUserBiz {
	return &changeRoleUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeRoleUserBiz) ChangeRoleUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateRole) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserUpdateRoleNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckUserStatusPermission(ctx, id); err != nil {
		return err
	}

	if err := biz.repo.UpdateRoleUser(ctx, id, data); err != nil {
		return err
	}

	return nil
}
