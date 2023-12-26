package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeStatusToppingsRepo struct {
	mock.Mock
}

func (m *mockChangeStatusToppingsRepo) ChangeStatusToppings(ctx context.Context, data []productmodel.ToppingUpdateStatus) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewChangeStatusToppingsBiz(t *testing.T) {
	type args struct {
		repo      ChangeStatusToppingsRepo
		requester middleware.Requester
	}
	mockRepo := new(mockChangeStatusToppingsRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *changeStatusToppingsBiz
	}{
		{
			name: "Create object has type ChangeStatusToppingsBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &changeStatusToppingsBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusToppingsBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewChangeStatusToppingsBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusToppingsBiz_ChangeStatusToppings(t *testing.T) {
	type fields struct {
		repo      ChangeStatusToppingsRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data []productmodel.ToppingUpdateStatus
	}

	mockRepo := new(mockChangeStatusToppingsRepo)
	mockRequest := new(mockRequester)

	active := true
	validData := []productmodel.ToppingUpdateStatus{
		{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: "Product001",
				IsActive:  &active,
			},
		},
	}
	invalidData := []productmodel.ToppingUpdateStatus{
		{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: "",
				IsActive:  &active,
			},
		},
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
			name: "Change status of toppings failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingUpdateStatusFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of toppings failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: invalidData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingUpdateStatusFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of toppings failed because repository returns an error",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("ChangeStatusToppings", context.Background(), validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of toppings successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("ChangeStatusToppings", context.Background(), validData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &changeStatusToppingsBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ChangeStatusToppings(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "ChangeStatusToppings() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeStatusToppings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
