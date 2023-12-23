package rolerepo

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
)

type ListRoleStore interface {
	ListRole(
		ctx context.Context,
	) ([]rolemodel.SimpleRole, error)
}

type listRoleRepo struct {
	store ListRoleStore
}

func NewListRoleRepo(
	store ListRoleStore) *listRoleRepo {
	return &listRoleRepo{store: store}
}

func (biz *listRoleRepo) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	result, err := biz.store.ListRole(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
