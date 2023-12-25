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

func TestNewSeeUserDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeUserDetailRepo
		requester middleware.Requester
	}

	mockRepo := new(mockSeeProfileRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeUserDetailBiz
	}{
		{
			name: "Create object has type SeeUserDetailBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeUserDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeUserDetailBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewSeeUserDetailBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeUserDetailBiz_SeeUserDetail(t *testing.T) {
	type fields struct {
		repo      SeeUserDetailRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockRepo := new(mockSeeProfileRepo)
	mockRequest := new(mockRequester)

	userId := "User001"
	user := usermodel.User{
		Id: userId,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.User
		wantErr bool
	}{
		{
			name: "See user detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &user,
			wantErr: true,
		},
		{
			name: "See user detail failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeUserDetail", context.Background(), userId).
					Return(nil, mockErr).
					Once()
			},
			want:    &user,
			wantErr: true,
		},
		{
			name: "See user detail successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeUserDetail", context.Background(), userId).
					Return(&user, nil).
					Once()
			},
			want:    &user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeUserDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeUserDetail(tt.args.ctx, tt.args.userId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeUserDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeUserDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "SeeUserDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
