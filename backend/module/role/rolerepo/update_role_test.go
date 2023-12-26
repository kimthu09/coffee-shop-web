package rolerepo

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateRoleStore struct {
	mock.Mock
}

func (m *mockUpdateRoleStore) UpdateRole(
	ctx context.Context,
	id string,
	data *rolemodel.RoleUpdate) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockUpdateRoleFeatureStore struct {
	mock.Mock
}

func (m *mockUpdateRoleFeatureStore) CreateListRoleFeatureDetail(
	ctx context.Context,
	data []rolefeaturemodel.RoleFeature) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockUpdateRoleFeatureStore) DeleteRoleFeature(
	ctx context.Context,
	conditions map[string]interface{}) error {
	args := m.Called(ctx, conditions)
	return args.Error(0)
}

func (m *mockUpdateRoleFeatureStore) FindListFeatures(
	ctx context.Context,
	roleId string) ([]rolefeaturemodel.RoleFeature, error) {
	args := m.Called(ctx, roleId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]rolefeaturemodel.RoleFeature), args.Error(1)
}

func TestNewUpdateRoleRepo(t *testing.T) {
	type args struct {
		roleStore        UpdateRoleStore
		roleFeatureStore UpdateRoleFeature
	}

	roleFeatureStore := new(mockUpdateRoleFeatureStore)
	roleStore := new(mockUpdateRoleStore)

	tests := []struct {
		name string
		args args
		want *updateRoleRepo
	}{
		{
			name: "Create object has type UpdateRoleRepo",
			args: args{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			want: &updateRoleRepo{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateRoleRepo(
				tt.args.roleStore, tt.args.roleFeatureStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewUpdateRoleRepo() = %v, want = %v",
				got,
				tt.want)

			assert.Equalf(t, tt.want, NewUpdateRoleRepo(tt.args.roleStore, tt.args.roleFeatureStore), "NewUpdateRoleRepo(%v, %v)", tt.args.roleStore, tt.args.roleFeatureStore)
		})
	}
}

func Test_updateRoleRepo_GetListRoleFeatures(t *testing.T) {
	type fields struct {
		roleStore        UpdateRoleStore
		roleFeatureStore UpdateRoleFeature
	}
	type args struct {
		ctx    context.Context
		roleId string
	}

	roleFeatureStore := new(mockUpdateRoleFeatureStore)
	roleStore := new(mockUpdateRoleStore)

	roleId := "role1"
	featureId1 := "feature1"
	featureId2 := "feature2"
	roleFeature1 := rolefeaturemodel.RoleFeature{
		RoleId:    roleId,
		FeatureId: featureId1,
	}
	roleFeature2 := rolefeaturemodel.RoleFeature{
		RoleId:    roleId,
		FeatureId: featureId2,
	}
	roleFeatures := []rolefeaturemodel.RoleFeature{roleFeature1, roleFeature2}
	featureStr := []string{featureId1, featureId2}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []string
		wantErr bool
	}{
		{
			name: "Get list role failed because can not get data from database",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleFeatureStore.
					On(
						"FindListFeatures",
						context.Background(),
						roleId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get list role successfully",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				roleFeatureStore.
					On(
						"FindListFeatures",
						context.Background(),
						roleId).
					Return(roleFeatures, nil).
					Once()
			},
			want:    featureStr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateRoleRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
			}

			tt.mock()

			got, err := repo.GetListRoleFeatures(tt.args.ctx, tt.args.roleId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"GetListRoleFeatures() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"GetListRoleFeatures() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"GetListRoleFeatures() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}

func Test_updateRoleRepo_UpdateRole(t *testing.T) {
	type fields struct {
		roleStore        UpdateRoleStore
		roleFeatureStore UpdateRoleFeature
	}
	type args struct {
		ctx    context.Context
		roleId string
		data   *rolemodel.RoleUpdate
	}

	roleFeatureStore := new(mockUpdateRoleFeatureStore)
	roleStore := new(mockUpdateRoleStore)

	roleId := mock.Anything
	roleName := mock.Anything
	featureIds := []string{mock.Anything, mock.Anything}
	roleUpdate := rolemodel.RoleUpdate{
		Name:     &roleName,
		Features: &featureIds,
	}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update role failed because can not save data to database",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data:   &roleUpdate,
			},
			mock: func() {
				roleStore.
					On(
						"UpdateRole",
						context.Background(),
						roleId,
						&roleUpdate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role successfully",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data:   &roleUpdate,
			},
			mock: func() {
				roleStore.
					On(
						"UpdateRole",
						context.Background(),
						roleId,
						&roleUpdate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateRoleRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
			}
			tt.mock()

			err := repo.UpdateRole(tt.args.ctx, tt.args.roleId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateRoleRepo_UpdateRoleFeatures(t *testing.T) {
	type fields struct {
		roleStore        UpdateRoleStore
		roleFeatureStore UpdateRoleFeature
	}
	type args struct {
		ctx                 context.Context
		deletedRoleFeatures []rolefeaturemodel.RoleFeature
		createdRoleFeatures []rolefeaturemodel.RoleFeature
	}

	roleFeatureStore := new(mockUpdateRoleFeatureStore)
	roleStore := new(mockUpdateRoleStore)

	roleId := "role1"
	feature1 := "feature1"
	feature2 := "feature2"
	feature3 := "feature3"
	deletedRoleFeatures := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: feature1,
		},
		{
			RoleId:    roleId,
			FeatureId: feature2,
		},
	}
	createdRoleFeatures := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: feature3,
		},
	}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update role feature failed because can not delete role feature from database",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:                 context.Background(),
				deletedRoleFeatures: deletedRoleFeatures,
				createdRoleFeatures: createdRoleFeatures,
			},
			mock: func() {
				roleFeatureStore.
					On(
						"DeleteRoleFeature",
						context.Background(),
						map[string]interface{}{
							"roleId":    roleId,
							"featureId": feature1,
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role feature failed because can not create role feature",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:                 context.Background(),
				deletedRoleFeatures: deletedRoleFeatures,
				createdRoleFeatures: createdRoleFeatures,
			},
			mock: func() {
				roleFeatureStore.
					On(
						"DeleteRoleFeature",
						context.Background(),
						map[string]interface{}{
							"roleId":    roleId,
							"featureId": feature1,
						}).
					Return(nil).
					Once()

				roleFeatureStore.
					On(
						"DeleteRoleFeature",
						context.Background(),
						map[string]interface{}{
							"roleId":    roleId,
							"featureId": feature2,
						}).
					Return(nil).
					Once()

				roleFeatureStore.
					On(
						"CreateListRoleFeatureDetail",
						context.Background(),
						createdRoleFeatures).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role feature successfully",
			fields: fields{
				roleStore:        roleStore,
				roleFeatureStore: roleFeatureStore,
			},
			args: args{
				ctx:                 context.Background(),
				deletedRoleFeatures: deletedRoleFeatures,
				createdRoleFeatures: createdRoleFeatures,
			},
			mock: func() {
				roleFeatureStore.
					On(
						"DeleteRoleFeature",
						context.Background(),
						map[string]interface{}{
							"roleId":    roleId,
							"featureId": feature1,
						}).
					Return(nil).
					Once()

				roleFeatureStore.
					On(
						"DeleteRoleFeature",
						context.Background(),
						map[string]interface{}{
							"roleId":    roleId,
							"featureId": feature2,
						}).
					Return(nil).
					Once()

				roleFeatureStore.
					On(
						"CreateListRoleFeatureDetail",
						context.Background(),
						createdRoleFeatures).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateRoleRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
			}

			tt.mock()

			err := repo.UpdateRoleFeatures(tt.args.ctx, tt.args.deletedRoleFeatures, tt.args.createdRoleFeatures)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
