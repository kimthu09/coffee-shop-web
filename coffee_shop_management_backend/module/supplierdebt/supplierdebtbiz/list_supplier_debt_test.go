package supplierdebtbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
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
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	args := m.Called(ctx, supplierId, paging)
	return args.Get(0).([]supplierdebtmodel.SupplierDebt),
		args.Error(1)
}

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
func (m *mockRequester) GetRole() rolemodel.Role {
	args := m.Called()
	return args.Get(0).(rolemodel.Role)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewListSupplierDebtBiz(t *testing.T) {
	type args struct {
		store     ListSupplierDebtStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListSupplierDebtStore)

	tests := []struct {
		name string
		args args
		want *listSupplierDebtBiz
	}{
		{
			name: "Create object has type ListIngredientDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listSupplierDebtBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListSupplierDebtBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListIngredientDetailByIdBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listSupplierDebtBiz_ListSupplierDebt(t *testing.T) {
	type fields struct {
		store     ListSupplierDebtStore
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		supplierId string
		paging     *common.Paging
	}

	mockStore := new(mockListSupplierDebtStore)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	supplierId := mock.Anything
	listSupplierDebts := make([]supplierdebtmodel.SupplierDebt, 0)
	var emptyListSupplierDebts []supplierdebtmodel.SupplierDebt
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
			name: "List supplier debt failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listSupplierDebts,
			wantErr: true,
		},
		{
			name: "List supplier debt failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListSupplierDebt",
						context.Background(),
						supplierId,
						&paging,
					).
					Return(emptyListSupplierDebts, mockErr).
					Once()
			},
			want:    listSupplierDebts,
			wantErr: true,
		},
		{
			name: "List supplier debt successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListSupplierDebt",
						context.Background(),
						supplierId,
						&paging,
					).
					Return(listSupplierDebts, nil).
					Once()
			},
			want:    listSupplierDebts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listSupplierDebtBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListSupplierDebt(tt.args.ctx, tt.args.supplierId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListSupplierDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListSupplierDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListSupplierDebt() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
