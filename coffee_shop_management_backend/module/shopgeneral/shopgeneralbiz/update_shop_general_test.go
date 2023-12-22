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

type mockUpdateGeneralShopStore struct {
	mock.Mock
}

func (m *mockUpdateGeneralShopStore) UpdateGeneralShop(ctx context.Context, data *shopgeneralmodel.ShopGeneralUpdate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewUpdateGeneralShopBiz(t *testing.T) {
	type args struct {
		store     UpdateGeneralShopStore
		requester middleware.Requester
	}

	mockStore := new(mockUpdateGeneralShopStore)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateGeneralShopBiz
	}{
		{
			name: "Create object has type UpdateGeneralShopBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &updateGeneralShopBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateGeneralShopBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewUpdateGeneralShopBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_updateGeneralShopBiz_UpdateGeneralShop(t *testing.T) {
	type fields struct {
		store     UpdateGeneralShopStore
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *shopgeneralmodel.ShopGeneralUpdate
	}
	mockStore := new(mockUpdateGeneralShopStore)
	mockRequest := new(mockRequester)

	str := mock.Anything
	email := "john@gmail.com"
	phone := "0123456789"
	floatNumber := float32(0)
	negativeFloatNumber := float32(-1)
	generalShopUpdate := shopgeneralmodel.ShopGeneralUpdate{
		Name:                   &str,
		Email:                  &email,
		Phone:                  &phone,
		Address:                &str,
		WifiPass:               &str,
		AccumulatePointPercent: &floatNumber,
		UsePointPercent:        &floatNumber,
	}
	invalidGeneralShopUpdate := shopgeneralmodel.ShopGeneralUpdate{
		Name:                   &str,
		Email:                  &email,
		Phone:                  &phone,
		Address:                &str,
		WifiPass:               &str,
		AccumulatePointPercent: &negativeFloatNumber,
		UsePointPercent:        &floatNumber,
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
			name: "Update shop general failed because user is not an admin",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &generalShopUpdate,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return("some-non-admin-role-id").
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update shop general failed because data validation fails",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidGeneralShopUpdate,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(common.RoleAdminId).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update shop general failed because can not save data to database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &generalShopUpdate,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(common.RoleAdminId).
					Once()

				mockStore.
					On(
						"UpdateGeneralShop",
						context.Background(),
						&generalShopUpdate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update shop general successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &generalShopUpdate,
			},
			mock: func() {
				mockRequest.
					On("GetRoleId").
					Return(common.RoleAdminId).
					Once()

				mockStore.
					On(
						"UpdateGeneralShop",
						context.Background(),
						&generalShopUpdate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateGeneralShopBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateGeneralShop(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"UpdateGeneralShop() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"UpdateGeneralShop() error = %v, wantErr %v",
					err,
					tt.wantErr)
			}
		})
	}
}
