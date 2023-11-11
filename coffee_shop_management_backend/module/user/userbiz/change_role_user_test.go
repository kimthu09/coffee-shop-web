package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeRoleUserRepo struct {
	mock.Mock
}

func (m *mockChangeRoleUserRepo) CheckRoleExist(
	ctx context.Context,
	roleId string) error {
	args := m.Called(ctx, roleId)
	return args.Error(0)
}

func (m *mockChangeRoleUserRepo) CheckUserStatusPermission(
	ctx context.Context,
	userId string) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *mockChangeRoleUserRepo) UpdateRoleUser(
	ctx context.Context,
	userId string,
	data *usermodel.UserUpdateRole) error {
	args := m.Called(ctx, userId, data)
	return args.Error(0)
}

type mockRequester struct {
	mock.Mock
}

func (m *mockRequester) GetUserId() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetEmail() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetRole() rolemodel.Role {
	args := m.Called()
	return args.Get(0).(rolemodel.Role)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewChangeRoleUserBiz(t *testing.T) {
	type args struct {
		repo      UpdateRoleUserRepo
		requester middleware.Requester
	}

	mockRepo := new(mockChangeRoleUserRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *changeRoleUserBiz
	}{
		{
			name: "Create object has type ChangeRoleUserBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &changeRoleUserBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeRoleUserBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewChangeRoleUserBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeRoleUserBiz_ChangeRoleUser(t *testing.T) {
	type fields struct {
		repo      UpdateRoleUserRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *usermodel.UserUpdateRole
	}

	mockRepo := new(mockChangeRoleUserRepo)
	mockRequest := new(mockRequester)

	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	noAdminRole := rolemodel.Role{Id: mock.Anything}

	userId := mock.Anything
	validId := "012345678901"
	invalidId := "This is invalid id"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Change role failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: validId},
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
			name: "Change role failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: invalidId},
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
			name: "Change role failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: validId},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change role failed because role need to update is not exist",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: validId},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CheckRoleExist",
						context.Background(),
						validId).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change role failed because can not save to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: validId},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CheckRoleExist",
						context.Background(),
						validId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateRoleUser",
						context.Background(),
						userId,
						&usermodel.UserUpdateRole{RoleId: validId}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change role successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &usermodel.UserUpdateRole{RoleId: validId},
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CheckRoleExist",
						context.Background(),
						validId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateRoleUser",
						context.Background(),
						userId,
						&usermodel.UserUpdateRole{RoleId: validId}).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &changeRoleUserBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ChangeRoleUser(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "ChangeRoleUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeRoleUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
