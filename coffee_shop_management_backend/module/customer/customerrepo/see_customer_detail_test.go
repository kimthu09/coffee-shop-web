package customerrepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockFindCustomerStore struct {
	mock.Mock
}

func (m *MockFindCustomerStore) FindCustomer(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*customermodel.Customer, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*customermodel.Customer), args.Error(1)
}

func TestNewSeeCustomerDetailRepo(t *testing.T) {
	type args struct {
		customerStore FindCustomerStore
	}

	mockStore := new(MockFindCustomerStore)

	tests := []struct {
		name string
		args args
		want *seeCustomerDetailRepo
	}{
		{
			name: "Create object has type NewSeeCustomerDetailRepo",
			args: args{
				customerStore: mockStore,
			},
			want: &seeCustomerDetailRepo{
				customerStore: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeCustomerDetailRepo(tt.args.customerStore)
			assert.Equal(t, tt.want, got, "NewSeeCustomerDetailRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeCustomerDetailRepo_SeeCustomerDetail(t *testing.T) {
	type fields struct {
		customerStore FindCustomerStore
	}
	type args struct {
		ctx        context.Context
		customerId string
	}

	mockStore := new(MockFindCustomerStore)

	customerId := "123"
	mockCustomer := &customermodel.Customer{
		Id:    customerId,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "1234567890",
		Point: 100.5,
	}

	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *customermodel.Customer
		wantErr bool
	}{
		{
			name: "See customer detail failed because can not get data from database",
			fields: fields{
				customerStore: mockStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				mockStore.
					On("FindCustomer",
						context.Background(),
						map[string]interface{}{"id": customerId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See customer detail successfully",
			fields: fields{
				customerStore: mockStore,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				mockStore.
					On("FindCustomer",
						context.Background(),
						map[string]interface{}{"id": customerId},
						moreKeys,
					).
					Return(mockCustomer, nil).
					Once()
			},
			want:    mockCustomer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeCustomerDetailRepo{
				customerStore: tt.fields.customerStore,
			}

			tt.mock()

			got, err := repo.SeeCustomerDetail(tt.args.ctx, tt.args.customerId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeCustomerDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeCustomerDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeCustomerDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
