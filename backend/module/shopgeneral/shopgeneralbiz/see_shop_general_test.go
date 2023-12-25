package shopgeneralbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

type mockSeeShopGeneralStore struct {
	mock.Mock
}

func (m *mockSeeShopGeneralStore) FindShopGeneral(ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*shopgeneralmodel.ShopGeneral), args.Error(1)
}

func TestNewSeeShopGeneralBiz(t *testing.T) {
	type args struct {
		store     SeeShopGeneralStore
		requester middleware.Requester
	}

	mockStore := new(mockSeeShopGeneralStore)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeShopGeneralBiz
	}{
		{
			name: "Create object has type SeeShopGeneralBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &seeShopGeneralBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeShopGeneralBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeShopGeneralBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeShopGeneralBiz_SeeShopGeneral(t *testing.T) {
	type fields struct {
		store     SeeShopGeneralStore
		requester middleware.Requester
	}
	type args struct {
		ctx context.Context
	}

	mockStore := new(mockSeeShopGeneralStore)
	mockRequest := new(mockRequester)

	shopGeneral := shopgeneralmodel.ShopGeneral{
		Name:                   mock.Anything,
		Email:                  mock.Anything,
		Phone:                  mock.Anything,
		Address:                mock.Anything,
		WifiPass:               mock.Anything,
		AccumulatePointPercent: 0.001,
		UsePointPercent:        1,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *shopgeneralmodel.ShopGeneral
		wantErr bool
	}{
		{
			name: "See shop general failed because user is not an admin",
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
					Return("some-non-admin-role-id").
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See shop general failed because store returns an error",
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
					Return(common.RoleAdminId).
					Once()

				mockStore.
					On("FindShopGeneral", context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See shop general successfully",
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
					Return(common.RoleAdminId).
					Once()

				mockStore.
					On("FindShopGeneral", context.Background()).
					Return(&shopGeneral, nil).
					Once()
			},
			want:    &shopGeneral,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeShopGeneralBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeShopGeneral(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeShopGeneral() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeShopGeneral() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeShopGeneral() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
