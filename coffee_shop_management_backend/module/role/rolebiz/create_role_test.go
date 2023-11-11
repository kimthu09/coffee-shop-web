package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) GenerateId() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockIdGenerator) IdProcess(id *string) (*string, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

type mockCreateRoleRepo struct {
	mock.Mock
}

func (m *mockCreateRoleRepo) CheckFeatureExist(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateRoleRepo) CreateRole(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateRoleRepo) CreateRoleFeatures(
	ctx context.Context,
	roleId string,
	featureIds []string) error {
	args := m.Called(ctx, roleId, featureIds)
	return args.Error(0)
}

func TestNewCreateRoleStore(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateRoleRepo
		requester middleware.Requester
	}

	mockRepo := new(mockCreateRoleRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createRoleStore
	}{
		{
			name: "Create object has type CreateRoleBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createRoleStore{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateRoleStore(
				tt.args.gen,
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateRoleStore() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createRoleStore_CreateRole(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateRoleRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *rolemodel.RoleCreate
	}

	mockRepo := new(mockCreateRoleRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	noAdminRole := rolemodel.Role{Id: mock.Anything}

	validId := "123456789012"
	validData := rolemodel.RoleCreate{
		Name:     mock.Anything,
		Features: []string{validId},
	}
	invalidData := rolemodel.RoleCreate{Name: ""}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create role failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
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
			name: "Create role failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidData,
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
			name: "Create role failed because user need to update is inactive",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(mock.Anything, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role failed because feature is not exist",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role failed because can not create role",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&validData).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateRole",
						context.Background(),
						&validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role failed because can not create role features",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&validData).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateRole",
						context.Background(),
						&validData).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateRoleFeatures",
						context.Background(),
						validData.Id,
						validData.Features).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create role successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &validData,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(validId, nil).
					Once()

				mockRepo.
					On(
						"CheckFeatureExist",
						context.Background(),
						&validData).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateRole",
						context.Background(),
						&validData).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateRoleFeatures",
						context.Background(),
						validData.Id,
						validData.Features).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createRoleStore{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateRole(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
