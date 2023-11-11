package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockPaySupplierRepo struct {
	mock.Mock
}

func (m *mockPaySupplierRepo) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*float32, error) {
	args := m.Called(ctx, supplierId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}
func (m *mockPaySupplierRepo) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
func (m *mockPaySupplierRepo) UpdateDebtSupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateDebt) error {
	args := m.Called(ctx, supplierId, data)
	return args.Error(0)
}
func TestNewUpdatePayBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      PaySupplierStoreRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)
	mockRepo := new(mockPaySupplierRepo)

	tests := []struct {
		name string
		args args
		want *paySupplierBiz
	}{
		{
			name: "Create object has type PaySupplierBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &paySupplierBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdatePayBiz(tt.args.gen, tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewUpdatePayBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_paySupplierBiz_PaySupplier(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      PaySupplierStoreRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		supplierId string
		data       *suppliermodel.SupplierUpdateDebt
	}

	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)
	mockRepo := new(mockPaySupplierRepo)
	supplierId := mock.Anything

	validAmount := float32(100)
	amount0 := float32(0)
	negativeAmount := float32(-1)
	createBy := mock.Anything
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount:   &validAmount,
		CreateBy: createBy,
	}
	supplierDebtId := mock.Anything
	currentDebtSupplier := float32(100)
	mockErr := errors.New(mock.Anything)
	debtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         supplierDebtId,
		SupplierId: supplierId,
		Amount:     float32(100),
		AmountLeft: float32(200),
		DebtType:   &debtType,
		CreateBy:   createBy,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *string
		wantErr bool
	}{
		{
			name: "Pay supplier failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(false).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because data is not validate with amount equals 0",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data: &suppliermodel.SupplierUpdateDebt{
					Amount: &amount0,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because data is not validate with amount < 0",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data: &suppliermodel.SupplierUpdateDebt{
					Amount: &negativeAmount,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because can not get debt supplier",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(nil, mockErr).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because can not get debt supplier",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(nil, mockErr).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(&currentDebtSupplier, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return("", mockErr).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because can not update debt of supplier",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(&currentDebtSupplier, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(supplierDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						supplierId,
						&supplierUpdateDebt).
					Return(mockErr).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier failed because can not create SupplierDebt",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(&currentDebtSupplier, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(supplierDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						supplierId,
						&supplierUpdateDebt).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						&supplierDebtCreate).
					Return(mockErr).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: true,
		},
		{
			name: "Pay supplier successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				data:       &supplierUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtSupplier",
						context.Background(),
						supplierId).
					Return(&currentDebtSupplier, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(supplierDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						supplierId,
						&supplierUpdateDebt).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						&supplierDebtCreate).
					Return(nil).
					Once()
			},
			want:    &supplierDebtId,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &paySupplierBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.PaySupplier(tt.args.ctx, tt.args.supplierId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"PaySupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"PaySupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"PaySupplier() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
