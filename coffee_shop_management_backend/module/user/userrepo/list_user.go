package userrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type ListUserStore interface {
	ListUser(
		ctx context.Context,
		filter *usermodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]usermodel.User, error)
}

type listUserRepo struct {
	store ListUserStore
}

func NewListUserRepo(store ListUserStore) *listUserRepo {
	return &listUserRepo{store: store}
}

func (repo *listUserRepo) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging) ([]usermodel.User, error) {
	result, err := repo.store.ListUser(
		ctx,
		filter,
		[]string{"id", "name", "email", "phone", "address"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
