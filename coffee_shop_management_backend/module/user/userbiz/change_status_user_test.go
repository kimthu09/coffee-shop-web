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

type mockChangeStatusUserRepo struct {
	mock.Mock
}

func (m *mockChangeStatusUserRepo) UpdateStatusUser(
	ctx context.Context,
	data *usermodel.UserUpdateStatus) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewChangeStatusUserBiz(t *testing.T) {
	type args struct {
		repo      ChangeStatusUserRepo
		requester middleware.Requester
	}

	mockRepo := new(mockChangeStatusUserRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *changeStatusUserBiz
	}{
		{
			name: "Create object has type ChangeStatusUserBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &changeStatusUserBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusUserBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewChangeStatusUserBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusUserBiz_ChangeStatusUser(t *testing.T) {
	type fields struct {
		repo      ChangeStatusUserRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data []usermodel.UserUpdateStatus
	}

	mockRepo := new(mockChangeStatusUserRepo)
	mockRequest := new(mockRequester)

	mockErr := errors.New(mock.Anything)

	validStatus := true
	userStatus := []usermodel.UserUpdateStatus{
		{
			UserId:   "User001",
			IsActive: &validStatus,
		},
	}
	invalidUserStatus := []usermodel.UserUpdateStatus{
		{
			UserId:   "",
			IsActive: &validStatus,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Change status user failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: userStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateStatusFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status user failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: invalidUserStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateStatusFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status user failed because can not save to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: userStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"UpdateStatusUser",
						context.Background(),
						&userStatus[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status user successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: userStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"UpdateStatusUser",
						context.Background(),
						&userStatus[0]).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &changeStatusUserBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ChangeStatusUser(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "ChangeStatusUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeStatusUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
