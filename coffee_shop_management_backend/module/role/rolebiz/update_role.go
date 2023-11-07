package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
)

type UpdateRoleRepo interface {
	CheckFeatureExist(
		ctx context.Context,
		data *rolemodel.RoleUpdate,
	) error
	GetListRoleFeatures(
		ctx context.Context,
		roleId string,
	) ([]string, error)
	UpdateRole(
		ctx context.Context,
		roleId string,
		data *rolemodel.RoleUpdate,
	) error
	UpdateRoleFeatures(
		ctx context.Context,
		deletedRoleFeatures []rolefeaturemodel.RoleFeature,
		createdRoleFeatures []rolefeaturemodel.RoleFeature,
	) error
}

type updateRoleBiz struct {
	repo      UpdateRoleRepo
	requester middleware.Requester
}

func NewUpdateRoleBiz(
	repo UpdateRoleRepo,
	requester middleware.Requester) *updateRoleBiz {
	return &updateRoleBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *updateRoleBiz) UpdateRole(
	ctx context.Context,
	roleId string,
	data *rolemodel.RoleUpdate) error {
	if biz.requester.GetRole().Id != common.RoleAdminId {
		return rolemodel.ErrRoleUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if data.Name != nil {
		if err := biz.repo.UpdateRole(ctx, roleId, data); err != nil {
			return err
		}
	}

	if data.Features != nil {
		if err := biz.repo.CheckFeatureExist(ctx, data); err != nil {
			return err
		}

		currentFeatures, errGetCurrentFeatures := biz.repo.GetListRoleFeatures(ctx, roleId)
		if errGetCurrentFeatures != nil {
			return errGetCurrentFeatures
		}

		mapFeatureExist := getMapFeatureExist(currentFeatures, *data.Features)
		deletedFeatures, createdFeatures := getDeletedAndCreatedFeaturesFromMapExist(
			roleId,
			mapFeatureExist,
		)

		if err := biz.repo.UpdateRoleFeatures(
			ctx,
			deletedFeatures,
			createdFeatures); err != nil {
			return err
		}
	}

	return nil
}

func getMapFeatureExist(
	currentFeatures []string,
	updatedFeatures []string) map[string]int {
	mapFeatureExist := make(map[string]int)
	for _, v := range currentFeatures {
		mapFeatureExist[v]--
	}
	for _, v := range updatedFeatures {
		mapFeatureExist[v]++
	}
	return mapFeatureExist
}

func getDeletedAndCreatedFeaturesFromMapExist(
	roleId string,
	mapFeatureExist map[string]int,
) ([]rolefeaturemodel.RoleFeature, []rolefeaturemodel.RoleFeature) {
	var deletedFeatures []rolefeaturemodel.RoleFeature
	var createdFeatures []rolefeaturemodel.RoleFeature
	for key, value := range mapFeatureExist {
		if value == 0 {
			continue
		} else {
			feature := rolefeaturemodel.RoleFeature{
				RoleId:    roleId,
				FeatureId: key,
			}
			if value == 1 {
				createdFeatures = append(createdFeatures, feature)
			} else if value == -1 {
				deletedFeatures = append(deletedFeatures, feature)
			}
		}
	}
	return deletedFeatures, createdFeatures
}
