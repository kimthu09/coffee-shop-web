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

type mockListToppingRepo struct {
	mock.Mock
}

func (m *mockListToppingRepo) ListTopping(ctx context.Context, filter *productmodel.Filter, paging *common.Paging) ([]productmodel.Topping, error) {
	args := m.Called(ctx, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]productmodel.Topping), args.Error(1)
}

func TestNewListToppingBiz(t *testing.T) {
	type args struct {
		repo      ListToppingRepo
		requester middleware.Requester
	}
	mockRepo := new(mockListToppingRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *listToppingBiz
	}{
		{
			name: "Create object has type ListToppingBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listToppingBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListToppingBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewListToppingBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listToppingBiz_ListTopping(t *testing.T) {
	type fields struct {
		repo      ListToppingRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *productmodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListToppingRepo)
	mockRequest := new(mockRequester)

	toppings := make([]productmodel.Topping, 0)
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
		want    []productmodel.Topping
		wantErr bool
	}{
		{
			name: "List toppings failed because user is not allowed",
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
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(false).
					Once()
			},
			want:    toppings,
			wantErr: true,
		},
		{
			name: "List toppings failed because repository returns an error",
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
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListTopping",
						context.Background(),
						&filter,
						&paging).
					Return(nil, mockErr).
					Once()
			},
			want:    toppings,
			wantErr: true,
		},
		{
			name: "List toppings successfully",
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
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListTopping",
						context.Background(),
						&filter,
						&paging).
					Return(toppings, nil).
					Once()
			},
			want:    toppings,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listToppingBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListTopping(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListTopping() = %v, want %v", got, tt.want)
			}
		})
	}
}
