package userbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInforUser struct {
	mock.Mock
}

func (m *mockUpdateInforUser) CheckUserStatusPermission(
	ctx context.Context,
	userId string) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}
func (m *mockUpdateInforUser) UpdateInfoUser(
	ctx context.Context,
	userId string,
	data *usermodel.UserUpdateInfo) error {
	args := m.Called(ctx, userId, data)
	return args.Error(0)
}

func TestNewUpdateInfoUserBiz(t *testing.T) {
	type args struct {
		repo      UpdateInfoUserRepo
		requester middleware.Requester
	}

	mockRepo := new(mockUpdateInforUser)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateInfoUserBiz
	}{
		{
			name: "Create object has type UpdateInfoUserBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &updateInfoUserBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoUserBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewUpdateInfoUserBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoUserBiz_UpdateUser(t *testing.T) {
	type fields struct {
		repo      UpdateInfoUserRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *usermodel.UserUpdateInfo
	}

	mockRepo := new(mockUpdateInforUser)
	mockRequest := new(mockRequester)
	userId := mock.Anything

	name := mock.Anything
	emptyName := ""
	phone := "0123456789"
	user := usermodel.UserUpdateInfo{
		Name:    &name,
		Phone:   &phone,
		Address: nil,
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
			name: "Update infor user failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &user,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update infor user failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				id:  userId,
				data: &usermodel.UserUpdateInfo{
					Name: &emptyName,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update infor user failed because user is inactive",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &user,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userId).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update infor user failed because can not save to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &user,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateInfoUser",
						context.Background(),
						userId,
						&user).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update infor user successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   userId,
				data: &user,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckUserStatusPermission",
						context.Background(),
						userId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateInfoUser",
						context.Background(),
						userId,
						&user).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateInfoUserBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateUser(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
