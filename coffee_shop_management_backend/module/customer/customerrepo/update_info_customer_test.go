package customerrepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoCustomerStore struct {
	mock.Mock
}

func (m *mockUpdateInfoCustomerStore) UpdateCustomerInfo(
	ctx context.Context,
	id string,
	data *customermodel.CustomerUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateInfoCustomerRepo(t *testing.T) {
	type args struct {
		store UpdateInfoCustomerStore
	}

	store := new(mockUpdateInfoCustomerStore)

	tests := []struct {
		name string
		args args
		want *updateInfoCustomerRepo
	}{
		{
			name: "Create object has type NewUpdateInfoCustomerRepo",
			args: args{
				store: store,
			},
			want: &updateInfoCustomerRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoCustomerRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewUpdateInfoCustomerRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoCustomerRepo_UpdateCustomerInfo(t *testing.T) {
	type fields struct {
		store UpdateInfoCustomerStore
	}
	type args struct {
		ctx        context.Context
		customerId string
		updateData *customermodel.CustomerUpdateInfo
	}

	store := new(mockUpdateInfoCustomerStore)

	customerId := "customer123"
	updateData := new(customermodel.CustomerUpdateInfo)

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update customer info failed because can not save to database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				updateData: updateData,
			},
			mock: func() {
				store.
					On("UpdateCustomerInfo",
						context.Background(),
						customerId,
						updateData,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update customer info successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				updateData: updateData,
			},
			mock: func() {
				store.
					On("UpdateCustomerInfo",
						context.Background(),
						customerId,
						updateData,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateInfoCustomerRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.UpdateCustomerInfo(tt.args.ctx, tt.args.customerId, tt.args.updateData)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
