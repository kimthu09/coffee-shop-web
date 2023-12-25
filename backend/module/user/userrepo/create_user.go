package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserRepo struct {
	userStore CreateUserStore
}

func NewCreateUserRepo(
	userStore CreateUserStore) *createUserRepo {
	return &createUserRepo{
		userStore: userStore,
	}
}

func (repo *createUserRepo) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := repo.userStore.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
