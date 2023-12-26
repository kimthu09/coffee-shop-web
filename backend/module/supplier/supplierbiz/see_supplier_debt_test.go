package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockSeeSupplierDebtRepo struct {
	mock.Mock
}

func (m *mockSeeSupplierDebtRepo) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	args := m.Called(ctx, supplierId, filterSupplierDebt, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]supplierdebtmodel.SupplierDebt), args.Error(1)
}

func TestNewSeeSupplierDebtBiz(t *testing.T) {
	type args struct {
		repo      SeeSupplierDebtRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierDebtRepo)

	tests := []struct {
		name string
		args args
		want *seeSupplierDebtBiz
	}{
		{
			name: "Create object has type SeeSupplierDebtBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeSupplierDebtBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierDebtBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeSupplierDebtBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeSupplierDebtBiz_SeeSupplierDebt(t *testing.T) {
	type fields struct {
		repo      SeeSupplierDebtRepo
		requester middleware.Requester
	}
	type args struct {
		ctx                context.Context
		supplierId         string
		filterSupplierDebt *filter.SupplierDebtFilter
		paging             *common.Paging
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierDebtRepo)
	supplierId := mock.Anything
	date := int64(123)
	filterSupplierDebt := filter.SupplierDebtFilter{
		DateFrom: &date,
		DateTo:   &date,
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
		Total: 12,
	}
	debtType := enum.Debt
	timeDebt := time.Date(2023, 12, 20, 0, 0, 0, 0, time.UTC)
	supplierDebts := []supplierdebtmodel.SupplierDebt{
		{
			Id:         mock.Anything,
			SupplierId: supplierId,
			Amount:     0,
			AmountLeft: 0,
			DebtType:   &debtType,
			CreatedBy:  mock.Anything,
			CreatedAt:  &timeDebt,
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []supplierdebtmodel.SupplierDebt
		wantErr bool
	}{
		{
			name: "See supplier debt failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                context.Background(),
				supplierId:         supplierId,
				filterSupplierDebt: &filterSupplierDebt,
				paging:             &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See supplier debt failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                context.Background(),
				supplierId:         supplierId,
				filterSupplierDebt: &filterSupplierDebt,
				paging:             &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierDebt",
						context.Background(),
						supplierId,
						&filterSupplierDebt,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "See supplier debt successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                context.Background(),
				supplierId:         supplierId,
				filterSupplierDebt: &filterSupplierDebt,
				paging:             &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierDebt",
						context.Background(),
						supplierId,
						&filterSupplierDebt,
						&paging,
					).
					Return(supplierDebts, nil).
					Once()
			},
			want:    supplierDebts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeSupplierDebtBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeSupplierDebt(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filterSupplierDebt,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeSupplierDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeSupplierDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeSupplierDebt() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
