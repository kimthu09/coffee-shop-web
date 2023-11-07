package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
)

type ListRoleRepo interface {
	ListRole(
		ctx context.Context,
	) ([]rolemodel.Role, error)
}

type listRoleBiz struct {
	repo      ListRoleRepo
	requester middleware.Requester
}

func NewListRoleBiz(
	repo ListRoleRepo,
	requester middleware.Requester) *listRoleBiz {
	return &listRoleBiz{repo: repo, requester: requester}
}

func (biz *listRoleBiz) ListRole(
	ctx context.Context) ([]rolemodel.Role, error) {
	if biz.requester.GetRole().Id != common.RoleAdminId {
		return nil, rolemodel.ErrRoleViewNoPermission
	}

	result, err := biz.repo.ListRole(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
