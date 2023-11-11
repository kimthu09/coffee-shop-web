package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockPayCustomerRepo struct {
	mock.Mock
}

func (m *mockPayCustomerRepo) GetDebtCustomer(
	ctx context.Context,
	customerId string) (*float32, error) {
	args := m.Called(ctx, customerId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}
func (m *mockPayCustomerRepo) CreateCustomerDebt(
	ctx context.Context,
	data *customerdebtmodel.CustomerDebtCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
func (m *mockPayCustomerRepo) UpdateDebtCustomer(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateDebt) error {
	args := m.Called(ctx, customerId, data)
	return args.Error(0)
}

func TestNewUpdatePayBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      PayCustomerRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)
	mockRepo := new(mockPayCustomerRepo)

	tests := []struct {
		name string
		args args
		want *payCustomerBiz
	}{
		{
			name: "Create object has type PayCustomerBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &payCustomerBiz{
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

func Test_payCustomerBiz_PayCustomer(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      PayCustomerRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		customerId string
		data       *customermodel.CustomerUpdateDebt
	}

	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)
	mockRepo := new(mockPayCustomerRepo)
	customerId := mock.Anything

	validAmount := float32(-100)
	amount0 := float32(0)
	positiveAmount := float32(1)
	createBy := mock.Anything
	customerUpdateDebt := customermodel.CustomerUpdateDebt{
		Amount:   &validAmount,
		CreateBy: createBy,
	}
	customerDebtId := mock.Anything
	currentDebtCustomer := float32(100)
	mockErr := errors.New(mock.Anything)
	debtType := enum.Pay
	customerDebtCreate := customerdebtmodel.CustomerDebtCreate{
		Id:         customerDebtId,
		CustomerId: customerId,
		Amount:     float32(-100),
		AmountLeft: float32(0),
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
			name: "Pay customer failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(false).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because data is not validate with amount equals 0",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data: &customermodel.CustomerUpdateDebt{
					Amount: &amount0,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because data is not validate with amount > 0",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data: &customermodel.CustomerUpdateDebt{
					Amount: &positiveAmount,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because can not get debt customer",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(nil, mockErr).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because can not get debt customer",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(nil, mockErr).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(&currentDebtCustomer, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return("", mockErr).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because can not update debt of customer",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(&currentDebtCustomer, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(customerDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtCustomer",
						context.Background(),
						customerId,
						&customerUpdateDebt).
					Return(mockErr).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer failed because can not create CustomerDebt",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(&currentDebtCustomer, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(customerDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtCustomer",
						context.Background(),
						customerId,
						&customerUpdateDebt).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateCustomerDebt",
						context.Background(),
						&customerDebtCreate).
					Return(mockErr).
					Once()
			},
			want:    &customerDebtId,
			wantErr: true,
		},
		{
			name: "Pay customer successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				data:       &customerUpdateDebt,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerPayFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"GetDebtCustomer",
						context.Background(),
						customerId).
					Return(&currentDebtCustomer, nil).
					Once()

				mockGenerator.
					On(
						"GenerateId").
					Return(customerDebtId, nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtCustomer",
						context.Background(),
						customerId,
						&customerUpdateDebt).
					Return(nil).
					Once()

				mockRepo.
					On(
						"CreateCustomerDebt",
						context.Background(),
						&customerDebtCreate).
					Return(nil).
					Once()
			},
			want:    &customerDebtId,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &payCustomerBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}
			tt.mock()

			got, err := biz.PayCustomer(tt.args.ctx, tt.args.customerId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"PayCustomer() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"PayCustomer() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"PayCustomer() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
