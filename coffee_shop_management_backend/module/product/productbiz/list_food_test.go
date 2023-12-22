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

type mockListFoodRepo struct {
	mock.Mock
}

func (m *mockListFoodRepo) ListFood(ctx context.Context, filter *productmodel.Filter, paging *common.Paging) ([]productmodel.Food, error) {
	args := m.Called(ctx, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]productmodel.Food), args.Error(1)
}

func TestNewListFoodBiz(t *testing.T) {
	type args struct {
		repo      ListFoodRepo
		requester middleware.Requester
	}
	mockRepo := new(mockListFoodRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *listFoodBiz
	}{
		{
			name: "Create object has type ListFoodBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listFoodBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListFoodBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewListFoodBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listFoodBiz_ListFood(t *testing.T) {
	type fields struct {
		repo      ListFoodRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *productmodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListFoodRepo)
	mockRequest := new(mockRequester)

	foods := make([]productmodel.Food, 0)
	mockErr := errors.New(mock.Anything)
	filter := productmodel.Filter{
		SearchKey: "",
		IsActive:  nil,
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []productmodel.Food
		wantErr bool
	}{
		{
			name: "List foods failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(false).
					Once()
			},
			want:    foods,
			wantErr: true,
		},
		{
			name: "List foods failed because repository returns an error",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListFood",
						context.Background(),
						&filter,
						&paging).
					Return(nil, mockErr).
					Once()
			},
			want:    foods,
			wantErr: true,
		},
		{
			name: "List foods successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListFood",
						context.Background(),
						&filter,
						&paging).
					Return(foods, nil).
					Once()
			},
			want:    foods,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listFoodBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListFood(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
