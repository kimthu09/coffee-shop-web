package customerdebtbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCustomerDebtStore struct {
	mock.Mock
}

func (m *mockListCustomerDebtStore) ListCustomerDebt(
	ctx context.Context,
	customerId string,
	paging *common.Paging) ([]customerdebtmodel.CustomerDebt, error) {
	args := m.Called(ctx, customerId, paging)
	return args.Get(0).([]customerdebtmodel.CustomerDebt),
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

func TestNewListCustomerDebtBiz(t *testing.T) {
	type args struct {
		store     ListCustomerDebtStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListCustomerDebtStore)

	tests := []struct {
		name string
		args args
		want *listCustomerDebtBiz
	}{
		{
			name: "Create object has type ListIngredientDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listCustomerDebtBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCustomerDebtBiz(tt.args.store, tt.args.requester)

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

func Test_listCustomerDebtBiz_ListCustomerDebt(t *testing.T) {
	type fields struct {
		store     ListCustomerDebtStore
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		customerId string
		paging     *common.Paging
	}

	mockStore := new(mockListCustomerDebtStore)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	customerId := mock.Anything
	listCustomerDebts := make([]customerdebtmodel.CustomerDebt, 0)
	var emptyListCustomerDebts []customerdebtmodel.CustomerDebt
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []customerdebtmodel.CustomerDebt
		wantErr bool
	}{
		{
			name: "List customer debt failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listCustomerDebts,
			wantErr: true,
		},
		{
			name: "List customer debt failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCustomerDebt",
						context.Background(),
						customerId,
						&paging,
					).
					Return(emptyListCustomerDebts, mockErr).
					Once()
			},
			want:    listCustomerDebts,
			wantErr: true,
		},
		{
			name: "List customer debt successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCustomerDebt",
						context.Background(),
						customerId,
						&paging,
					).
					Return(listCustomerDebts, nil).
					Once()
			},
			want:    listCustomerDebts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listCustomerDebtBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListCustomerDebt(tt.args.ctx, tt.args.customerId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCustomerDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListCustomerDebt() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCustomerDebt() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
