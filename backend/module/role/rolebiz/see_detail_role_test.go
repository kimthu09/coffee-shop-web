package rolebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/feature/featuremodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeDetailRoleRepo struct {
	mock.Mock
}

func (m *mockSeeDetailRoleRepo) SeeRoleDetail(
	ctx context.Context,
	roleId string) (*rolemodel.RoleDetail, error) {
	args := m.Called(ctx, roleId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rolemodel.RoleDetail), args.Error(1)
}

func TestNewSeeDetailRoleBiz(t *testing.T) {
	type args struct {
		repo      SeeDetailRoleRepo
		requester middleware.Requester
	}

	mockRepo := new(mockSeeDetailRoleRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeDetailRoleBiz
	}{
		{
			name: "Create object has type SeeDetailRoleBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeDetailRoleBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailRoleBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeDetailRoleBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeDetailRoleBiz_SeeDetailRole(t *testing.T) {
	type fields struct {
		repo      SeeDetailRoleRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		roleId string
	}

	mockRepo := new(mockSeeDetailRoleRepo)
	mockRequest := new(mockRequester)

	adminRole := rolemodel.Role{Id: common.RoleAdminId}
	noAdminRole := rolemodel.Role{Id: mock.Anything}
	roleId := mock.Anything
	role := rolemodel.RoleDetail{
		Id:   roleId,
		Name: mock.Anything,
		Features: []featuremodel.FeatureDetail{
			{
				Id:          mock.Anything,
				Description: mock.Anything,
				GroupName:   mock.Anything,
				IsHas:       true,
			},
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *rolemodel.RoleDetail
		wantErr bool
	}{
		{
			name: "See role detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
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
			name: "See role detail failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockRepo.
					On("SeeRoleDetail", context.Background(), roleId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See role detail successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(adminRole.Id).
					Once()

				mockRepo.
					On("SeeRoleDetail", context.Background(), roleId).
					Return(&role, nil).
					Once()
			},
			want:    &role,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeDetailRoleBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}
			tt.mock()

			got, err := biz.SeeDetailRole(tt.args.ctx, tt.args.roleId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeDetailRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeDetailRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeDetailRole() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
