package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoSupplierStore struct {
	mock.Mock
}

func (m *mockUpdateInfoSupplierStore) UpdateSupplierInfo(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateInfoSupplierRepo(t *testing.T) {
	type args struct {
		store UpdateInfoSupplierStore
	}

	store := new(mockUpdateInfoSupplierStore)

	tests := []struct {
		name string
		args args
		want *updateInfoSupplierRepo
	}{
		{
			name: "Create object has type NewUpdateInfoSupplierRepo",
			args: args{
				store: store,
			},
			want: &updateInfoSupplierRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoSupplierRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewUpdateInfoSupplierRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoSupplierRepo_UpdateSupplierInfo(t *testing.T) {
	type fields struct {
		store UpdateInfoSupplierStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
		updateData *suppliermodel.SupplierUpdateInfo
	}

	store := new(mockUpdateInfoSupplierStore)

	supplierId := "supplier123"
	supplierUpdateInfo := new(suppliermodel.SupplierUpdateInfo)

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update supplier info failed because can not save to database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				updateData: supplierUpdateInfo,
			},
			mock: func() {
				store.
					On("UpdateSupplierInfo",
						context.Background(),
						supplierId,
						supplierUpdateInfo,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier info successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				updateData: supplierUpdateInfo,
			},
			mock: func() {
				store.
					On("UpdateSupplierInfo",
						context.Background(),
						supplierId,
						supplierUpdateInfo,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateInfoSupplierRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.UpdateSupplierInfo(tt.args.ctx, tt.args.supplierId, tt.args.updateData)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateSupplierInfo() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateSupplierInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
