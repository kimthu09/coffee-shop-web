package userbiz

import (
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeProfileRepo struct {
	mock.Mock
}

func (m *mockSeeProfileRepo) SeeUserDetail(
	ctx context.Context,
	userId string) (*usermodel.User, error) {
	args := m.Called(ctx, userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func TestNewSeeProfileBiz(t *testing.T) {
	type args struct {
		repo      SeeUserDetailRepo
		requester middleware.Requester
	}

	mockRepo := new(mockSeeProfileRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeProfileBiz
	}{
		{
			name: "Create object has type SeeProfileBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeProfileBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeProfileBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewSeeProfileBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeProfileBiz_SeeProfile(t *testing.T) {
	type fields struct {
		repo      SeeProfileRepo
		requester middleware.Requester
	}
	type args struct {
		ctx context.Context
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
			name: "See profile failed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetUserId").
					Return(userId).
					Once()

				mockRepo.
					On("SeeUserDetail", context.Background(), userId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See profile successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetUserId").
					Return(userId).
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
			biz := &seeProfileBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}
			tt.mock()

			got, err := biz.SeeProfile(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeProfile() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeProfile() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "SeeProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
