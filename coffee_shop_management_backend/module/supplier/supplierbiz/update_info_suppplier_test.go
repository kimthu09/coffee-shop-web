package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoSupplier struct {
	mock.Mock
}

func (m *mockUpdateInfoSupplier) CheckExist(
	ctx context.Context,
	supplierId string) error {
	args := m.Called(ctx, supplierId)
	return args.Error(0)
}
func (m *mockUpdateInfoSupplier) UpdateSupplierInfo(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateInfo) error {
	args := m.Called(ctx, supplierId, data)
	return args.Error(0)
}

func TestNewUpdateInfoSupplierBiz(t *testing.T) {
	type args struct {
		repo      UpdateInfoSupplierRepo
		requester middleware.Requester
	}

	mockRepo := new(mockUpdateInfoSupplier)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateInfoSupplierBiz
	}{
		{
			name: "Create object has type UpdateInfoSupplierBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &updateInfoSupplierBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoSupplierBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewUpdateInfoSupplierBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoSupplierBiz_UpdateInfoSupplier(t *testing.T) {
	type fields struct {
		repo      UpdateInfoSupplierRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *suppliermodel.SupplierUpdateInfo
	}

	mockRepo := new(mockUpdateInfoSupplier)
	mockRequest := new(mockRequester)
	supplierId := mock.Anything

	name := mock.Anything
	email := "a@gmail.com"
	phone := "0123456789"
	empty := ""
	supplierUpdateInfo := suppliermodel.SupplierUpdateInfo{
		Name:  &name,
		Email: &email,
		Phone: &phone,
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
			name: "Update supplier info failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierId,
				data: &supplierUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier info failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierId,
				data: &suppliermodel.SupplierUpdateInfo{Name: &empty},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier info failed because supplier is not exist",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierId,
				data: &supplierUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckExist",
						context.Background(),
						supplierId).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier info failed because can not save to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierId,
				data: &supplierUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckExist",
						context.Background(),
						supplierId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateSupplierInfo",
						context.Background(),
						supplierId,
						&supplierUpdateInfo).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier info successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierId,
				data: &supplierUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"CheckExist",
						context.Background(),
						supplierId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateSupplierInfo",
						context.Background(),
						supplierId,
						&supplierUpdateInfo).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateInfoSupplierBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateInfoSupplier(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateInfoSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateInfoSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
