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

type mockCreateRoleStore struct {
	mock.Mock
}

func (m *mockCreateRoleStore) CreateRole(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateRoleFeaturesStore struct {
	mock.Mock
}

func (m *mockCreateRoleFeaturesStore) CreateListRoleFeatureDetail(
	ctx context.Context,
	data []rolefeaturemodel.RoleFeature) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateRoleRepo(t *testing.T) {
	type args struct {
		roleStore        CreateRoleStore
		roleFeatureStore CreateListRoleFeatureStore
	}

	mockRole := new(mockCreateRoleStore)
	mockRoleFeature := new(mockCreateRoleFeaturesStore)

	tests := []struct {
		name string
		args args
		want *createRoleRepo
	}{
		{
			name: "Create object has type CreateRoleRepo",
			args: args{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
			want: &createRoleRepo{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateRoleRepo(
				tt.args.roleStore,
				tt.args.roleFeatureStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateRoleRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createRoleRepo_CreateRole(t *testing.T) {
	type fields struct {
		roleStore        CreateRoleStore
		roleFeatureStore CreateListRoleFeatureStore
	}
	type args struct {
		ctx  context.Context
		data *rolemodel.RoleCreate
	}

	mockRole := new(mockCreateRoleStore)
	mockRoleFeature := new(mockCreateRoleFeaturesStore)

	roleId := "role001"
	featureId1 := "role0001"
	featureId2 := "role0002"
	data := rolemodel.RoleCreate{
		Id:       roleId,
		Name:     mock.Anything,
		Features: []string{featureId1, featureId2},
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
			name: "Create role failed because can not save data to database",
			fields: fields{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRole.
					On(
						"CreateRole",
						context.Background(),
						&data).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role successfully",
			fields: fields{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
			args: args{
				ctx:  context.Background(),
				data: &data,
			},
			mock: func() {
				mockRole.
					On(
						"CreateRole",
						context.Background(),
						&data).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createRoleRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
			}

			tt.mock()

			err := repo.CreateRole(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createRoleRepo_CreateRoleFeatures(t *testing.T) {
	type fields struct {
		roleStore        CreateRoleStore
		roleFeatureStore CreateListRoleFeatureStore
	}
	type args struct {
		ctx        context.Context
		roleId     string
		featureIds []string
	}

	mockRole := new(mockCreateRoleStore)
	mockRoleFeature := new(mockCreateRoleFeaturesStore)

	roleId := "role001"
	featureId1 := "role0001"
	featureId2 := "role0002"
	listFeature := []string{featureId1, featureId2}
	listRoleFeatures := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: featureId1,
		},
		{
			RoleId:    roleId,
			FeatureId: featureId2,
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
			name: "Create role feature failed because can not save data to database",
			fields: fields{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
			args: args{
				ctx:        context.Background(),
				roleId:     roleId,
				featureIds: listFeature,
			},
			mock: func() {
				mockRoleFeature.
					On(
						"CreateListRoleFeatureDetail",
						context.Background(),
						listRoleFeatures).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role feature successfully",
			fields: fields{
				roleStore:        mockRole,
				roleFeatureStore: mockRoleFeature,
			},
			args: args{
				ctx:        context.Background(),
				roleId:     roleId,
				featureIds: listFeature,
			},
			mock: func() {
				mockRoleFeature.
					On(
						"CreateListRoleFeatureDetail",
						context.Background(),
						listRoleFeatures).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createRoleRepo{
				roleStore:        tt.fields.roleStore,
				roleFeatureStore: tt.fields.roleFeatureStore,
			}

			tt.mock()

			err := repo.CreateRoleFeatures(tt.args.ctx, tt.args.roleId, tt.args.featureIds)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateRoleFeatures() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateRoleFeatures() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
