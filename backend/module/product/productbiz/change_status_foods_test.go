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

type mockChangeStatusFoodsRepo struct {
	mock.Mock
}

func (m *mockChangeStatusFoodsRepo) ChangeStatusFoods(ctx context.Context, data []productmodel.FoodUpdateStatus) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewChangeStatusFoodsBiz(t *testing.T) {
	type args struct {
		repo      ChangeStatusFoodsRepo
		requester middleware.Requester
	}
	mockRepo := new(mockChangeStatusFoodsRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *changeStatusFoodsBiz
	}{
		{
			name: "Create object has type ChangeStatusFoodsBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &changeStatusFoodsBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusFoodsBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewChangeStatusFoodsBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusFoodsBiz_ChangeStatusFoods(t *testing.T) {
	type fields struct {
		repo      ChangeStatusFoodsRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data []productmodel.FoodUpdateStatus
	}

	mockRepo := new(mockChangeStatusFoodsRepo)
	mockRequest := new(mockRequester)

	active := true
	validData := []productmodel.FoodUpdateStatus{
		{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: "Product001",
				IsActive:  &active,
			},
		},
	}
	invalidData := []productmodel.FoodUpdateStatus{
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
			name: "Change status of foods failed because user is not allowed",
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
					On("IsHasFeature", common.FoodUpdateStatusFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of foods failed because data is invalid",
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
					On("IsHasFeature", common.FoodUpdateStatusFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of foods failed because repository returns an error",
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
					On("IsHasFeature", common.FoodUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("ChangeStatusFoods", context.Background(), validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status of foods successfully",
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
					On("IsHasFeature", common.FoodUpdateStatusFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("ChangeStatusFoods", context.Background(), validData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &changeStatusFoodsBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ChangeStatusFoods(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "ChangeStatusFoods() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeStatusFoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
