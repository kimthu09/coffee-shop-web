package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListSupplierDebtStore struct {
	mock.Mock
}

func (m *mockListSupplierDebtStore) ListSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging,
	moreKeys ...string) ([]supplierdebtmodel.SupplierDebt, error) {
	args := m.Called(ctx, supplierId, filterSupplierDebt, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]supplierdebtmodel.SupplierDebt), args.Error(1)
}

func TestNewSeeSupplierDebtRepo(t *testing.T) {
	type args struct {
		debtStore ListSupplierDebtStore
	}

	store := new(mockListSupplierDebtStore)

	tests := []struct {
		name string
		args args
		want *seeSupplierDebtRepo
	}{
		{
			name: "Create object has type NewSeeSupplierDebtRepo",
			args: args{
				debtStore: store,
			},
			want: &seeSupplierDebtRepo{
				debtStore: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierDebtRepo(tt.args.debtStore)
			assert.Equal(t, tt.want, got, "NewSeeSupplierDebtRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeSupplierDebtRepo_SeeSupplierDebt(t *testing.T) {
	type fields struct {
		debtStore ListSupplierDebtStore
	}
	type args struct {
		ctx                context.Context
		supplierId         string
		filterSupplierDebt *filter.SupplierDebtFilter
		paging             *common.Paging
	}

	store := new(mockListSupplierDebtStore)

	supplierId := mock.Anything
	debtType := enum.Debt
	supplierDebts := []supplierdebtmodel.SupplierDebt{
		{
			Id:         mock.Anything,
			SupplierId: supplierId,
			Amount:     0,
			AmountLeft: 0,
			DebtType:   &debtType,
			CreatedBy:  mock.Anything,
		},
	}

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
	moreKeys := []string{"CreatedByUser"}

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
			name: "See supplier debt failed because can not get data from database",
			fields: fields{
				debtStore: store,
			},
			args: args{
				ctx:                context.Background(),
				supplierId:         supplierId,
				filterSupplierDebt: &filterSupplierDebt,
				paging:             &paging,
			},
			mock: func() {
				store.
					On("ListSupplierDebt",
						context.Background(),
						supplierId,
						&filterSupplierDebt,
						&paging,
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    supplierDebts,
			wantErr: true,
		},
		{
			name: "See supplier debt successfully",
			fields: fields{
				debtStore: store,
			},
			args: args{
				ctx:                context.Background(),
				supplierId:         supplierId,
				filterSupplierDebt: &filterSupplierDebt,
				paging:             &paging,
			},
			mock: func() {
				store.
					On("ListSupplierDebt",
						context.Background(),
						supplierId,
						&filterSupplierDebt,
						&paging,
						moreKeys,
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
			biz := &seeSupplierDebtRepo{
				debtStore: tt.fields.debtStore,
			}

			tt.mock()

			got, err := biz.SeeSupplierDebt(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filterSupplierDebt,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeSupplierDebt() = %v, want %v", got, tt.want)
			}
		})
	}
}
