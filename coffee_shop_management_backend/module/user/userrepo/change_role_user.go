package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type ChangeRoleUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
	UpdateRoleUser(
		ctx context.Context,
		id string,
		data *usermodel.UserUpdateRole) error
}

type changeRoleUserRepo struct {
	userStore ChangeRoleUserStore
	roleStore CheckRoleStore
}

func NewChangeRoleUserRepo(
	userStore ChangeRoleUserStore,
	roleStore CheckRoleStore) *changeRoleUserRepo {
	return &changeRoleUserRepo{
		userStore: userStore,
		roleStore: roleStore,
	}
}

func (repo *changeRoleUserRepo) CheckRoleExist(ctx context.Context, roleId string) error {
	if _, err := repo.roleStore.FindRole(
		ctx,
		map[string]interface{}{
			"id": roleId,
		},
	); err != nil {
		return err
	}

	return nil
}

func (repo *changeRoleUserRepo) CheckUserStatusPermission(
	ctx context.Context,
	userId string) error {
	currentUser, err := repo.userStore.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return err
	}

	if !currentUser.IsActive {
		return usermodel.ErrUserInactive
	}
	return nil
}

func (repo *changeRoleUserRepo) UpdateRoleUser(
	ctx context.Context,
	userId string,
	data *usermodel.UserUpdateRole) error {
	if err := repo.userStore.UpdateRoleUser(ctx, userId, data); err != nil {
		return err
	}
	return nil
}
