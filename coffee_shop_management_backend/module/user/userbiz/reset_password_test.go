package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/hasher"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockResetPasswordRepo struct {
	mock.Mock
}

func (m *mockResetPasswordRepo) GetUser(
	ctx context.Context,
	id string) (*usermodel.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}
func (m *mockResetPasswordRepo) CheckUserStatusPermission(
	ctx context.Context,
	userId string) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}
func (m *mockResetPasswordRepo) UpdateUserPassword(
	ctx context.Context,
	id string,
	pass string) error {
	args := m.Called(ctx, id, pass)
	return args.Error(0)
}

func TestNewResetPasswordBiz(t *testing.T) {
	type args struct {
		repo      ResetPasswordRepo
		hasher    hasher.Hasher
		requester middleware.Requester
	}

	mockRepo := new(mockResetPasswordRepo)
	mockHash := new(mockHasher)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *resetPasswordBiz
	}{
		{
			name: "Create object has type ResetPasswordRepo",
			args: args{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			want: &resetPasswordBiz{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewResetPasswordBiz(
				tt.args.repo,
				tt.args.hasher,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewResetPasswordBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_resetPasswordBiz_ResetPassword(t *testing.T) {
	type fields struct {
		repo      ResetPasswordRepo
		hasher    hasher.Hasher
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *usermodel.UserResetPassword
	}

	mockRepo := new(mockResetPasswordRepo)
	mockHash := new(mockHasher)
	mockRequest := new(mockRequester)
	userIdNeedToReset := mock.Anything

	pass := mock.Anything
	userResetPass := usermodel.UserResetPassword{
		UserSenderPass: pass,
	}

	userId := mock.Anything
	hashedPassword := mock.Anything
	userRequest := usermodel.User{
		Id:       userId,
		Password: hashedPassword,
		Salt:     mock.Anything,
	}
	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	notAdminRole := rolemodel.Role{Id: mock.Anything}

	newHashedPassword := mock.Anything
	userNeedToReset := usermodel.User{
		Id:   userIdNeedToReset,
		Salt: mock.Anything,
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
			name: "Reset password failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(notAdminRole).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password failed because data is not valid",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &usermodel.UserResetPassword{UserSenderPass: ""},
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
			name: "Reset password failed because can note get user make request",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password failed because wrong password user request",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(&userRequest, nil).
					Once()

				mockHash.
					On(
						"Hash",
						userResetPass.UserSenderPass+userRequest.Salt).
					Return("").
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password failed because user need to update is inactive",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(&userRequest, nil).
					Once()

				mockHash.
					On(
						"Hash",
						userResetPass.UserSenderPass+userRequest.Salt).
					Return(hashedPassword).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userIdNeedToReset).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password failed because can not get user need to reset",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(&userRequest, nil).
					Once()

				mockHash.
					On(
						"Hash",
						userResetPass.UserSenderPass+userRequest.Salt).
					Return(hashedPassword).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userIdNeedToReset).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userIdNeedToReset).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password failed because can not update user",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(&userRequest, nil).
					Once()

				mockHash.
					On(
						"Hash",
						userResetPass.UserSenderPass+userRequest.Salt).
					Return(hashedPassword).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userIdNeedToReset).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userIdNeedToReset).
					Return(&userNeedToReset, nil).
					Once()

				mockHash.
					On(
						"Hash",
						common.DefaultPass+userNeedToReset.Salt).
					Return(newHashedPassword).
					Once()

				mockRepo.
					On(
						"UpdateUserPassword",
						context.Background(),
						userIdNeedToReset,
						newHashedPassword).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Reset password successfully",
			fields: fields{
				repo:      mockRepo,
				hasher:    mockHash,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userIdNeedToReset,
				data: &userResetPass,
			},
			mock: func() {
				mockRequest.
					On("GetRole").
					Return(adminRole).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userId,
					).
					Return(&userRequest, nil).
					Once()

				mockHash.
					On(
						"Hash",
						userResetPass.UserSenderPass+userRequest.Salt).
					Return(hashedPassword).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userIdNeedToReset).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetUser",
						context.Background(),
						userIdNeedToReset).
					Return(&userNeedToReset, nil).
					Once()

				mockHash.
					On(
						"Hash",
						common.DefaultPass+userNeedToReset.Salt).
					Return(newHashedPassword).
					Once()

				mockRepo.
					On(
						"UpdateUserPassword",
						context.Background(),
						userIdNeedToReset,
						newHashedPassword).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &resetPasswordBiz{
				repo:      tt.fields.repo,
				hasher:    tt.fields.hasher,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ResetPassword(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
