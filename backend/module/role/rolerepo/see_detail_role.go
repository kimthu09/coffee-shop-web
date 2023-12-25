package rolerepo

import (
	"coffee_shop_management_backend/module/feature/featuremodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
)

type ListRoleFeaturesStore interface {
	FindListFeatures(
		ctx context.Context,
		roleId string,
	) ([]rolefeaturemodel.RoleFeature, error)
}

type ListAllFeatures interface {
	ListFeature(
		ctx context.Context,
	) ([]featuremodel.Feature, error)
}

type FindRoleStore interface {
	FindRole(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*rolemodel.Role, error)
}

type seeRoleDetailRepo struct {
	roleStore        FindRoleStore
	roleFeatureStore ListRoleFeaturesStore
	featureStore     ListAllFeatures
}

func NewSeeRoleDetailRepo(
	roleStore FindRoleStore,
	roleFeatureStore ListRoleFeaturesStore,
	featureStore ListAllFeatures) *seeRoleDetailRepo {
	return &seeRoleDetailRepo{
		roleStore:        roleStore,
		roleFeatureStore: roleFeatureStore,
		featureStore:     featureStore,
	}
}

func (biz *seeRoleDetailRepo) SeeRoleDetail(
	ctx context.Context,
	roleId string) (*rolemodel.RoleDetail, error) {
	role, errRole := biz.roleStore.FindRole(
		ctx, map[string]interface{}{"id": roleId})
	if errRole != nil {
		return nil, errRole
	}

	features, errFeature := biz.featureStore.ListFeature(ctx)
	if errFeature != nil {
		return nil, errFeature
	}

	featuresRoleHas, errRoleFeature :=
		biz.roleFeatureStore.FindListFeatures(ctx, roleId)
	if errRoleFeature != nil {
		return nil, errRoleFeature
	}

	mapHasFeature := make(map[string]bool)
	for _, v := range features {
		mapHasFeature[v.Id] = false
	}
	for _, v := range featuresRoleHas {
		mapHasFeature[v.FeatureId] = true
	}

	var featureDetails []featuremodel.FeatureDetail
	for _, v := range features {
		featureDetail := featuremodel.FeatureDetail{
			Id:          v.Id,
			Description: v.Description,
			GroupName:   v.GroupName,
			IsHas:       mapHasFeature[v.Id],
		}
		featureDetails = append(featureDetails, featureDetail)
	}

	var roleDetail rolemodel.RoleDetail
	roleDetail.Id = role.Id
	roleDetail.Name = role.Name
	roleDetail.Features = featureDetails

	return &roleDetail, nil
}
