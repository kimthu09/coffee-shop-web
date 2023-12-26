package rolerepo

import (
	"coffee_shop_management_backend/module/feature/featuremodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListRoleFeaturesStore struct {
	mock.Mock
}

func (m *mockListRoleFeaturesStore) FindListFeatures(
	ctx context.Context,
	roleId string) ([]rolefeaturemodel.RoleFeature, error) {
	args := m.Called(ctx, roleId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]rolefeaturemodel.RoleFeature), args.Error(1)
}

type mockListAllFeatures struct {
	mock.Mock
}

func (m *mockListAllFeatures) ListFeature(
	ctx context.Context) ([]featuremodel.Feature, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]featuremodel.Feature), args.Error(1)
}

type mockFindRoleStore struct {
	mock.Mock
}

func (m *mockFindRoleStore) FindRole(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*rolemodel.Role, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rolemodel.Role), args.Error(1)
}

func TestNewSeeRoleDetailRepo(t *testing.T) {
	type args struct {
		roleStore        FindRoleStore
		roleFeatureStore ListRoleFeaturesStore
		featureStore     ListAllFeatures
	}

	roleFeatureStore := new(mockListRoleFeaturesStore)
	roleStore := new(mockFindRoleStore)
	featureStore := new(mockListAllFeatures)

	tests := []struct {
		name string
		args args
		want *seeRoleDetailRepo
	}{
		{
			name: "Create object has type SeeRoleDetailRepo",
			args: args{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
			want: &seeRoleDetailRepo{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeRoleDetailRepo(
				tt.args.roleStore, tt.args.roleFeatureStore, tt.args.featureStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewSeeRoleDetailRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_seeRoleDetailRepo_SeeRoleDetail(t *testing.T) {
	type fields struct {
		roleStore        FindRoleStore
		roleFeatureStore ListRoleFeaturesStore
		featureStore     ListAllFeatures
	}
	type args struct {
		ctx    context.Context
		roleId string
	}

	roleFeatureStore := new(mockListRoleFeaturesStore)
	roleStore := new(mockFindRoleStore)
	featureStore := new(mockListAllFeatures)

	roleId := mock.Anything
	roleName := mock.Anything
	featureId1 := "feature1"
	featureId2 := "feature2"
	role := rolemodel.Role{
		Id:   roleId,
		Name: roleName,
	}

	feature1 := featuremodel.Feature{
		Id:          featureId1,
		Description: mock.Anything,
		GroupName:   mock.Anything,
	}
	feature2 := featuremodel.Feature{
		Id:          featureId2,
		Description: mock.Anything,
		GroupName:   mock.Anything,
	}

	roleFeature1 := rolefeaturemodel.RoleFeature{
		RoleId:    roleId,
		FeatureId: featureId1,
	}

	features := []featuremodel.Feature{feature1, feature2}

	roleFeatures := []rolefeaturemodel.RoleFeature{roleFeature1}

	featureDetails := []featuremodel.FeatureDetail{
		{
			Id:          featureId1,
			Description: feature1.Description,
			GroupName:   feature1.GroupName,
			IsHas:       true,
		},
		{
			Id:          featureId2,
			Description: feature2.Description,
			GroupName:   feature2.GroupName,
			IsHas:       false,
		},
	}

	roleDetail := rolemodel.RoleDetail{
		Id:       roleId,
		Name:     roleName,
		Features: featureDetails,
	}

	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *rolemodel.RoleDetail
		wantErr bool
	}{
		{
			name: "See detail role failed because can not find role from database",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleStore.
					On(
						"FindRole",
						context.Background(),
						map[string]interface{}{"id": roleId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail role failed because can not list all features",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleStore.
					On(
						"FindRole",
						context.Background(),
						map[string]interface{}{"id": roleId},
						moreKeys).
					Return(&role, nil).
					Once()

				featureStore.
					On(
						"ListFeature",
						context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail role failed because can not get features that role has",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleStore.
					On(
						"FindRole",
						context.Background(),
						map[string]interface{}{"id": roleId},
						moreKeys).
					Return(&role, nil).
					Once()

				featureStore.
					On(
						"ListFeature",
						context.Background()).
					Return(features, nil).
					Once()

				roleFeatureStore.
					On(
						"FindListFeatures",
						context.Background(), roleId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail role successfully",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
				featureStore:     featureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleStore.
					On(
						"FindRole",
						context.Background(),
						map[string]interface{}{"id": roleId},
						moreKeys).
					Return(&role, nil).
					Once()

				featureStore.
					On(
						"ListFeature",
						context.Background()).
					Return(features, nil).
					Once()

				roleFeatureStore.
					On(
						"FindListFeatures",
						context.Background(), roleId).
					Return(roleFeatures, nil).
					Once()
			},
			want:    &roleDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeRoleDetailRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
				featureStore:     tt.fields.featureStore,
			}

			tt.mock()

			got, err := biz.SeeRoleDetail(tt.args.ctx, tt.args.roleId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeRoleDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeRoleDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeRoleDetail() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
