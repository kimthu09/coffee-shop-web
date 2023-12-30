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

type mockSeeDetailToppingRepo struct {
	mock.Mock
}

func (m *mockSeeDetailToppingRepo) SeeDetailTopping(ctx context.Context, toppingId string) (*productmodel.Topping, error) {
	args := m.Called(ctx, toppingId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Topping), args.Error(1)
}

func TestNewSeeDetailToppingBiz(t *testing.T) {
	type args struct {
		repo      SeeDetailToppingRepo
		requester middleware.Requester
	}
	mockRepo := new(mockSeeDetailToppingRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *seeDetailToppingBiz
	}{
		{
			name: "Create object has type SeeDetailToppingBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeDetailToppingBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailToppingBiz(tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewSeeDetailToppingBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeDetailToppingBiz_SeeDetailTopping(t *testing.T) {
	type fields struct {
		repo      SeeDetailToppingRepo
		requester middleware.Requester
	}
	type args struct {
		ctx       context.Context
		toppingId string
	}
	mockRepo := new(mockSeeDetailToppingRepo)
	mockRequest := new(mockRequester)

	toppingDetail := &productmodel.Topping{}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Topping
		wantErr bool
	}{
		{
			name: "See detail topping failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				toppingId: "validToppingId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail topping failed because repository returns an error",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				toppingId: "validToppingId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeDetailTopping", context.Background(), "validToppingId").
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail topping successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				toppingId: "validToppingId",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("SeeDetailTopping", context.Background(), "validToppingId").
					Return(toppingDetail, nil).
					Once()
			},
			want:    toppingDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeDetailToppingBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeDetailTopping(tt.args.ctx, tt.args.toppingId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeDetailTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeDetailTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeDetailTopping() = %v, want %v", got, tt.want)
			}
		})
	}
}
