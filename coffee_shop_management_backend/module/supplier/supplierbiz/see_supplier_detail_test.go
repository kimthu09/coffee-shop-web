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

type mockSeeSupplierDetailRepo struct {
	mock.Mock
}

func (m *mockSeeSupplierDetailRepo) SeeSupplierDetail(
	ctx context.Context,
	supplierId string) (*suppliermodel.Supplier, error) {
	args := m.Called(ctx, supplierId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*suppliermodel.Supplier), args.Error(1)
}

func TestNewSeeSupplierDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeSupplierDetailRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierDetailRepo)

	tests := []struct {
		name string
		args args
		want *seeSupplierDetailBiz
	}{
		{
			name: "Create object has type SeeSupplierDetailBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeSupplierDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierDetailBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeSupplierDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeSupplierDetailBiz_SeeSupplierDetail(t *testing.T) {
	type fields struct {
		repo      SeeSupplierDetailRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		supplierId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierDetailRepo)

	supplierId := mock.Anything
	supplier := suppliermodel.Supplier{
		Id:    supplierId,
		Name:  mock.Anything,
		Email: mock.Anything,
		Phone: mock.Anything,
		Debt:  0,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "See supplier detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &supplier,
			wantErr: true,
		},
		{
			name: "See supplier detail failed because can not get supplier from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierDetail",
						context.Background(),
						supplierId).
					Return(nil, mockErr).
					Once()
			},
			want:    &supplier,
			wantErr: true,
		},
		{
			name: "See supplier detail successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierDetail",
						context.Background(),
						supplierId).
					Return(&supplier, nil).
					Once()
			},
			want:    &supplier,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeSupplierDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeSupplierDetail(tt.args.ctx, tt.args.supplierId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeSupplierDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeSupplierDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeSupplierDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
