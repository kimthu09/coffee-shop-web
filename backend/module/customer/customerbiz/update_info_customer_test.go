package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoCustomer struct {
	mock.Mock
}

func (m *mockUpdateInfoCustomer) UpdateCustomerInfo(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateInfo) error {
	args := m.Called(ctx, customerId, data)
	return args.Error(0)
}

func TestNewUpdateInfoCustomerBiz(t *testing.T) {
	type args struct {
		repo      UpdateInfoCustomerRepo
		requester middleware.Requester
	}

	mockRepo := new(mockUpdateInfoCustomer)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateInfoCustomerBiz
	}{
		{
			name: "Create object has type UpdateInfoCustomerBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &updateInfoCustomerBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoCustomerBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewUpdateInfoCustomerBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoCustomerBiz_UpdateInfoCustomer(t *testing.T) {
	type fields struct {
		repo      UpdateInfoCustomerRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *customermodel.CustomerUpdateInfo
	}

	mockRepo := new(mockUpdateInfoCustomer)
	mockRequest := new(mockRequester)
	customerId := mock.Anything

	name := mock.Anything
	email := "a@gmail.com"
	phone := "0123456789"
	empty := ""
	customerUpdateInfo := customermodel.CustomerUpdateInfo{
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
			name: "Update customer info failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: &customerUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update customer info failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: &customermodel.CustomerUpdateInfo{Name: &empty},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update customer info failed because can not save to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: &customerUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"UpdateCustomerInfo",
						context.Background(),
						customerId,
						&customerUpdateInfo).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update customer info successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: &customerUpdateInfo,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"UpdateCustomerInfo",
						context.Background(),
						customerId,
						&customerUpdateInfo).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateInfoCustomerBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateInfoCustomer(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateInfoCustomer() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateInfoCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
