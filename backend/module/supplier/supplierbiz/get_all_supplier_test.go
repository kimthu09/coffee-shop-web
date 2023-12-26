package supplierbiz

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAllSupplierRepo struct {
	mock.Mock
}

func (m *mockGetAllSupplierRepo) GetAllSupplier(
	ctx context.Context) ([]suppliermodel.SimpleSupplier, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]suppliermodel.SimpleSupplier), args.Error(1)
}

func TestNewGetAllSupplierBiz(t *testing.T) {
	type args struct {
		repo GetAllSupplierRepo
	}

	mockRepo := new(mockGetAllSupplierRepo)

	tests := []struct {
		name string
		args args
		want *getAllSupplierBiz
	}{
		{
			name: "Create object has type GetAllSupplierBiz",
			args: args{
				repo: mockRepo,
			},
			want: &getAllSupplierBiz{
				repo: mockRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllSupplierBiz(tt.args.repo)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewGetAllSupplierBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_getAllSupplierBiz_GetAllUser(t *testing.T) {
	type fields struct {
		repo GetAllSupplierRepo
	}
	type args struct {
		ctx context.Context
	}

	mockRepo := new(mockGetAllSupplierRepo)
	listSuppliers := []suppliermodel.SimpleSupplier{
		{
			Id:   mock.Anything,
			Name: mock.Anything,
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
			name: "Get all supplier failed because can not get data from database",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllSupplier", context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get all supplier successfully",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllSupplier", context.Background()).
					Return(listSuppliers, nil).
					Once()
			},
			want:    listSuppliers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &getAllSupplierBiz{
				repo: tt.fields.repo,
			}

			tt.mock()

			got, err := biz.GetAllSupplier(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"GetAllSupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"GetAllSupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"GetAllSupplier() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
