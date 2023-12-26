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

type mockListUserRepo struct {
	mock.Mock
}

func (m *mockListUserRepo) ListUser(
	ctx context.Context,
	userSearch string,
	filter *usermodel.Filter,
	paging *common.Paging) ([]usermodel.User, error) {
	args := m.Called(ctx, userSearch, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]usermodel.User), args.Error(1)
}

func TestNewListUserBiz(t *testing.T) {
	type args struct {
		repo      ListUserRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListUserRepo)

	tests := []struct {
		name string
		args args
		want *listUserBiz
	}{
		{
			name: "Create object has type ListUserBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listUserBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListUserBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListUserBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listUserBiz_ListUser(t *testing.T) {
	type fields struct {
		repo      ListUserRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *usermodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListUserRepo)
	mockRequest := new(mockRequester)
	filter := usermodel.Filter{
		SearchKey: "",
		IsActive:  nil,
		Role:      "",
	}
	paging := common.Paging{
		Page: 1,
	}
	userSearch := "User001"
	listUser := make([]usermodel.User, 0)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []usermodel.User
		wantErr bool
	}{
		{
			name: "List user failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List user failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userSearch).
					Once()

				mockRepo.
					On(
						"ListUser",
						context.Background(),
						userSearch,
						&filter,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List user successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.UserViewFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(userSearch).
					Once()

				mockRepo.
					On(
						"ListUser",
						context.Background(),
						userSearch,
						&filter,
						&paging,
					).
					Return(listUser, nil).
					Once()
			},
			want:    listUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listUserBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}
			tt.mock()

			got, err := biz.ListUser(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListUser() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListUser() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListUser() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
