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

type mockSeeDetailFoodRepo struct {
	mock.Mock
}

func (m *mockSeeDetailFoodRepo) SeeDetailFood(ctx context.Context, foodId string) (*productmodel.Food, error) {
	args := m.Called(ctx, foodId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Food), args.Error(1)
}

func TestNewSeeDetailFoodBiz(t *testing.T) {
	type args struct {
		repo      SeeDetailFoodRepo
		requester middleware.Requester
	}
	mockRepo := new(mockSeeDetailFoodRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeDetailFoodBiz
	}{
		{
			name: "Create object has type SeeDetailFoodBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeDetailFoodBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailFoodBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewSeeDetailFoodBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeDetailFoodBiz_SeeDetailFood(t *testing.T) {
	type fields struct {
		repo      SeeDetailFoodRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		foodId string
	}
	mockRepo := new(mockSeeDetailFoodRepo)
	mockRequest := new(mockRequester)

	foodDetail := &productmodel.Food{}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Food
		wantErr bool
	}{
		{
			name: "See detail food failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "validFoodId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail food failed because repository returns an error",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "validFoodId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeDetailFood", context.Background(), "validFoodId").
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail food successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "validFoodId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeDetailFood", context.Background(), "validFoodId").
					Return(foodDetail, nil).
					Once()
			},
			want:    foodDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeDetailFoodBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeDetailFood(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeDetailFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeDetailFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeDetailFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
