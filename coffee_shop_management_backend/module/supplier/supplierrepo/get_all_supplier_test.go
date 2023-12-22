package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAllSupplierStore struct {
	mock.Mock
}

func (m *mockGetAllSupplierStore) GetAllSupplier(
	ctx context.Context,
	moreKeys ...string) ([]suppliermodel.SimpleSupplier, error) {
	args := m.Called(ctx, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]suppliermodel.SimpleSupplier), args.Error(1)
}

func TestNewGetAllSupplierRepo(t *testing.T) {
	type args struct {
		store GetAllSupplierStore
	}

	store := new(mockGetAllSupplierStore)

	tests := []struct {
		name string
		args args
		want *getAllSupplierRepo
	}{
		{
			name: "Create object has type ListSupplierRepo",
			args: args{
				store: store,
			},
			want: &getAllSupplierRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllSupplierRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewGetAllSupplierRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_getAllSupplierRepo_GetAllSupplier(t *testing.T) {
	type fields struct {
		store GetAllSupplierStore
	}
	type args struct {
		ctx context.Context
	}

	store := new(mockGetAllSupplierStore)
	var moreKeys []string
	simpleSuppliers := []suppliermodel.SimpleSupplier{
		{
			Id:   "",
			Name: "",
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []suppliermodel.SimpleSupplier
		wantErr bool
	}{
		{
			name: "Get all suppliers failed because can not get data from database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				store.
					On("GetAllSupplier",
						context.Background(),
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    simpleSuppliers,
			wantErr: true,
		},
		{
			name: "Get all suppliers successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				store.
					On("GetAllSupplier",
						context.Background(),
						moreKeys).
					Return(simpleSuppliers, nil).
					Once()
			},
			want:    simpleSuppliers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &getAllSupplierRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.GetAllSupplier(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetAllSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
