package featurebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/feature/featuremodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type mockListFeatureStore struct {
	mock.Mock
}

func (m *mockListFeatureStore) ListFeature(
	ctx context.Context) ([]featuremodel.Feature, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]featuremodel.Feature), args.Error(1)
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

func TestNewListFeatureBiz(t *testing.T) {
	type args struct {
		store     ListFeatureStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListFeatureStore)

	tests := []struct {
		name string
		args args
		want *listFeatureBiz
	}{
		{
			name: "Create object has type ListFeatureBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listFeatureBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListFeatureBiz(tt.args.store, tt.args.requester); !reflect.DeepEqual(got, tt.want) {
				got := NewListFeatureBiz(tt.args.store, tt.args.requester)

				assert.Equal(
					t,
					tt.want,
					got,
					"NewListFeatureBiz() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}

func Test_listFeatureBiz_ListFeature(t *testing.T) {
	type fields struct {
		store     ListFeatureStore
		requester middleware.Requester
	}
	type args struct {
		ctx context.Context
	}

	mockStore := new(mockListFeatureStore)
	mockRequest := new(mockRequester)
	listFeature := make([]featuremodel.Feature, 0)
	mockErr := errors.New(mock.Anything)
	roleAdmin := rolemodel.Role{
		Id: common.RoleAdminId,
	}
	roleNoAdmin := rolemodel.Role{
		Id: mock.Anything,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []featuremodel.Feature
		wantErr bool
	}{
		{
			name: "List features failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(roleNoAdmin.Id).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List feature failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(roleAdmin.Id).
					Once()

				mockStore.
					On(
						"ListFeature",
						context.Background(),
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List feature successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(roleAdmin.Id).
					Once()

				mockStore.
					On(
						"ListFeature",
						context.Background(),
					).
					Return(listFeature, nil).
					Once()
			},
			want:    listFeature,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listFeatureBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListFeature(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListFeature() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListFeature() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListFeature() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
