package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateRoleRepo struct {
	mock.Mock
}

func (m *mockUpdateRoleRepo) CheckFeatureExist(
	ctx context.Context,
	data *rolemodel.RoleUpdate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
func (m *mockUpdateRoleRepo) GetListRoleFeatures(
	ctx context.Context,
	roleId string) ([]string, error) {
	args := m.Called(ctx, roleId)
	return args.Get(0).([]string), args.Error(1)
}
func (m *mockUpdateRoleRepo) UpdateRole(
	ctx context.Context,
	roleId string,
	data *rolemodel.RoleUpdate) error {
	args := m.Called(ctx, roleId, data)
	return args.Error(0)
}
func (m *mockUpdateRoleRepo) UpdateRoleFeatures(
	ctx context.Context,
	deletedRoleFeatures []rolefeaturemodel.RoleFeature,
	createdRoleFeatures []rolefeaturemodel.RoleFeature) error {
	args := m.Called(ctx, deletedRoleFeatures, createdRoleFeatures)
	return args.Error(0)
}

func TestNewUpdateRoleBiz(t *testing.T) {
	type args struct {
		repo      UpdateRoleRepo
		requester middleware.Requester
	}

	mockRepo := new(mockUpdateRoleRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateRoleBiz
	}{
		{
			name: "Create object has type CreateRoleBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &updateRoleBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateRoleBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewUpdateRoleBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateRoleBiz_UpdateRole(t *testing.T) {
	type fields struct {
		repo      UpdateRoleRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		roleId string
		data   *rolemodel.RoleUpdate
	}

	mockRepo := new(mockUpdateRoleRepo)
	mockRequest := new(mockRequester)

	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	noAdminRole := rolemodel.Role{Id: mock.Anything}

	roleId := mock.Anything
	name := mock.Anything
	invalidName := ""
	validId := "012345678901"
	duplicateId := "012345678902"
	features := []string{
		validId,
		duplicateId,
	}
	currentFeatures := []string{
		mock.Anything,
		duplicateId,
	}
	createdRoleFeature := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: features[0],
		},
	}
	deletedRoleFeature := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: currentFeatures[0],
		},
	}
	var emptyFeatures []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update role failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     &name,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(noAdminRole).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     &invalidName,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role failed because can not save role to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     &name,
					Features: nil,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"UpdateRole",
						context.Background(),
						roleId,
						&rolemodel.RoleUpdate{
							Name:     &name,
							Features: nil,
						},
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role successfully case don't need update features",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     &name,
					Features: nil,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"UpdateRole",
						context.Background(),
						roleId,
						&rolemodel.RoleUpdate{
							Name:     &name,
							Features: nil,
						},
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "Update role failed because can not check features is exist",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     nil,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&rolemodel.RoleUpdate{
							Name:     nil,
							Features: &features,
						},
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role failed because can not get current list features",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     nil,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&rolemodel.RoleUpdate{
							Name:     nil,
							Features: &features,
						},
					).
					Return(nil).
					Times(len(features))

				mockRepo.
					On(
						"GetListRoleFeatures",
						context.Background(),
						roleId,
					).
					Return(emptyFeatures, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role failed because can not save role features to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     nil,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						roleId,
					).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetListRoleFeatures",
						context.Background(),
						roleId,
					).
					Return(currentFeatures, nil).
					Times(len(features))

				mockRepo.
					On(
						"UpdateRoleFeatures",
						context.Background(),
						deletedRoleFeature,
						createdRoleFeature,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update role successfully case don't need update name",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
				data: &rolemodel.RoleUpdate{
					Name:     nil,
					Features: &features,
				},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						roleId,
					).
					Return(nil).
					Times(len(features))

				mockRepo.
					On(
						"GetListRoleFeatures",
						context.Background(),
						roleId,
					).
					Return(currentFeatures, nil).
					Once()

				mockRepo.
					On(
						"UpdateRoleFeatures",
						context.Background(),
						deletedRoleFeature,
						createdRoleFeature,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateRoleBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateRole(tt.args.ctx, tt.args.roleId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
