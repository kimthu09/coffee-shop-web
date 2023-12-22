package rolerepo

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
)

type CreateRoleStore interface {
	CreateRole(
		ctx context.Context,
		data *rolemodel.RoleCreate,
	) error
}

type CreateListRoleFeatureStore interface {
	CreateListRoleFeatureDetail(
		ctx context.Context,
		data []rolefeaturemodel.RoleFeature,
	) error
}

type createRoleRepo struct {
	roleStore        CreateRoleStore
	roleFeatureStore CreateListRoleFeatureStore
}

func NewCreateRoleRepo(
	roleStore CreateRoleStore,
	roleFeatureStore CreateListRoleFeatureStore) *createRoleRepo {
	return &createRoleRepo{
		roleStore:        roleStore,
		roleFeatureStore: roleFeatureStore,
	}
}

func (repo *createRoleRepo) CreateRole(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	if err := repo.roleStore.CreateRole(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *createRoleRepo) CreateRoleFeatures(
	ctx context.Context,
	roleId string,
	featureIds []string) error {
	var featureCreates []rolefeaturemodel.RoleFeature

	for _, v := range featureIds {
		featureCreate := rolefeaturemodel.RoleFeature{
			RoleId:    roleId,
			FeatureId: v,
		}
		featureCreates = append(featureCreates, featureCreate)
	}
	if err := repo.roleFeatureStore.CreateListRoleFeatureDetail(
		ctx, featureCreates,
	); err != nil {
		return err
	}
	return nil
}
