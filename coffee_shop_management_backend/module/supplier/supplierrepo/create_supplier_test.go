package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateSupplierStore struct {
	mock.Mock
}

func (m *mockCreateSupplierStore) CreateSupplier(
	ctx context.Context, data *suppliermodel.SupplierCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateSupplierRepo(t *testing.T) {
	type args struct {
		store CreateSupplierStore
	}

	store := new(mockCreateSupplierStore)

	tests := []struct {
		name string
		args args
		want *createSupplierRepo
	}{
		{
			name: "Create object has type CreateSupplierRepo",
			args: args{
				store: store,
			},
			want: &createSupplierRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateSupplierRepo(
				tt.args.store,
			)

			assert.Equal(t, tt.want, got, "NewCreateSupplierRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createSupplierRepo_CreateSupplier(t *testing.T) {
	type fields struct {
		store CreateSupplierStore
	}
	type args struct {
		ctx  context.Context
		data *suppliermodel.SupplierCreate
	}

	store := new(mockCreateSupplierStore)

	id := mock.Anything
	supplierCreate := suppliermodel.SupplierCreate{
		Id:    &id,
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
		wantErr bool
	}{
		{
			name: "Create supplier failed because can not save to database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				store.
					On("CreateSupplier",
						context.Background(),
						&supplierCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				store.
					On("CreateSupplier",
						context.Background(),
						&supplierCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createSupplierRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.CreateSupplier(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
