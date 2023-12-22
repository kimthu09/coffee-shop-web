package customerrepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCreateCustomerStore struct {
	mock.Mock
}

func (m *MockCreateCustomerStore) CreateCustomer(
	ctx context.Context,
	data *customermodel.CustomerCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateCustomerRepo(t *testing.T) {
	type args struct {
		store CreateCustomerStore
	}

	store := new(MockCreateCustomerStore)

	tests := []struct {
		name string
		args args
		want *createCustomerRepo
	}{
		{
			name: "Create object has type CreateCustomerRepo",
			args: args{
				store: store,
			},
			want: &createCustomerRepo{
				store: store,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateCustomerRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewCreateCustomerRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createCustomerRepo_CreateCustomer(t *testing.T) {
	type fields struct {
		store CreateCustomerStore
	}
	type args struct {
		ctx  context.Context
		data *customermodel.CustomerCreate
	}

	mockStore := new(MockCreateCustomerStore)

	id := "123"
	customerCreate := customermodel.CustomerCreate{
		Id:    &id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "1234567890",
	}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create customer failed because can not save to the database",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockStore.
					On("CreateCustomer",
						context.Background(),
						&customerCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create customer successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockStore.
					On("CreateCustomer",
						context.Background(),
						&customerCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCustomerRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.CreateCustomer(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
