package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListRoleRepo struct {
	mock.Mock
}

func (m *mockListRoleRepo) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]rolemodel.SimpleRole), args.Error(1)
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
func (m *mockRequester) GetRoleId() string {
	args := m.Called()
	return args.Get(0).(string)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewListRoleBiz(t *testing.T) {
	type args struct {
		repo      ListRoleRepo
		requester middleware.Requester
	}

	mockRepo := new(mockListRoleRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *listRoleBiz
	}{
		{
			name: "Create object has type ListRoleBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listRoleBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewListRoleBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListRoleBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listRoleBiz_ListRole(t *testing.T) {
	type fields struct {
		repo      ListRoleRepo
		requester middleware.Requester
	}
	type args struct {
		ctx context.Context
	}

	mockRepo := new(mockListRoleRepo)
	mockRequest := new(mockRequester)

	adminRole := rolemodel.SimpleRole{Id: common.RoleAdminId}
	noAdminRole := rolemodel.SimpleRole{Id: mock.Anything}
	listRole := make([]rolemodel.SimpleRole, 0)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []rolemodel.SimpleRole
		wantErr bool
	}{
		{
			name: "List role failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(noAdminRole.Id).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List role failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockRepo.
					On(
						"ListRole",
						context.Background(),
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List role successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockRepo.
					On(
						"ListRole",
						context.Background(),
					).
					Return(listRole, nil).
					Once()
			},
			want:    listRole,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listRoleBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListRole(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListRole() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
