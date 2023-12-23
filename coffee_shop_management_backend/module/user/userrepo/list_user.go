package userrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

type ListUserStore interface {
	ListUser(
		ctx context.Context,
		userSearch string,
		filter *usermodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string,
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
	userSearch string,
	filter *usermodel.Filter,
	paging *common.Paging) ([]usermodel.User, error) {
	result, err := repo.store.ListUser(
		ctx,
		userSearch,
		filter,
		[]string{"id", "name", "email"},
		paging,
		"Role")

	if err != nil {
		return nil, err
	}

	return result, nil
}
