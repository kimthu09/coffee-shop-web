package userbiz

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAllUserRepo struct {
	mock.Mock
}

func (m *mockGetAllUserRepo) GetAllUser(
	ctx context.Context) ([]usermodel.SimpleUser, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]usermodel.SimpleUser), args.Error(1)
}

func TestNewGetAllUserBiz(t *testing.T) {
	type args struct {
		repo GetAllUserRepo
	}

	mockRepo := new(mockGetAllUserRepo)

	tests := []struct {
		name string
		args args
		want *getAllUserBiz
	}{
		{
			name: "Create object has type GetAllUserBiz",
			args: args{
				repo: mockRepo,
			},
			want: &getAllUserBiz{
				repo: mockRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllUserBiz(tt.args.repo)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewGetAllUserBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_getAllUserBiz_GetAllUser(t *testing.T) {
	type fields struct {
		repo GetAllUserRepo
	}
	type args struct {
		ctx context.Context
	}

	mockRepo := new(mockGetAllUserRepo)
	mockErr := errors.New(mock.Anything)
	listUser := make([]usermodel.SimpleUser, 0)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []usermodel.SimpleUser
		wantErr bool
	}{
		{
			name: "Get all user failed",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllUser", context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get all user successfully",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllUser", context.Background()).
					Return(listUser, nil).
					Once()
			},
			want:    listUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &getAllUserBiz{
				repo: tt.fields.repo,
			}
			tt.mock()

			got, err := biz.GetAllUser(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"GetAllUser() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"GetAllUser() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"GetAllUser() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
