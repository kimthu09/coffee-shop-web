package invoicebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateInvoiceRepo struct {
	mock.Mock
}

func (m *mockCreateInvoiceRepo) GetShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*shopgeneralmodel.ShopGeneral), args.Error(1)
}

func (m *mockCreateInvoiceRepo) HandleCheckPermissionStatus(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateInvoiceRepo) HandleData(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateInvoiceRepo) FindCustomer(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	args := m.Called(ctx, customerId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*customermodel.Customer), args.Error(1)
}

func (m *mockCreateInvoiceRepo) UpdateCustomerPoint(
	ctx context.Context,
	customerId string,
	data customermodel.CustomerUpdatePoint) error {
	args := m.Called(ctx, customerId, data)
	return args.Error(0)
}

func (m *mockCreateInvoiceRepo) HandleInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateInvoiceRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	invoiceId string,
	ingredientTotalAmountNeedUpdate map[string]int) error {
	args := m.Called(ctx, invoiceId, ingredientTotalAmountNeedUpdate)
	return args.Error(0)
}

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) GenerateId() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockIdGenerator) IdProcess(id *string) (*string, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

func TestNewCreateInvoiceBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateInvoiceRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockCreateInvoiceRepo)
	mockGen := new(mockIdGenerator)
	tests := []struct {
		name string
		args args
		want *createInvoiceBiz
	}{
		{
			name: "Create object has type CreateInvoiceBiz",
			args: args{
				gen:       mockGen,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createInvoiceBiz{
				gen:       mockGen,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateInvoiceBiz(tt.args.gen, tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewCreateInvoiceBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_createInvoiceBiz_CreateInvoice(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateInvoiceRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *invoicemodel.InvoiceCreate
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockCreateInvoiceRepo)
	mockGen := new(mockIdGenerator)

	customerId := "Customer001"
	invalidCustomerId := "1245678901234567890"
	invoiceId := "Invoice001"
	invoiceCreate := invoicemodel.InvoiceCreate{
		CustomerId: &customerId,
		IsUsePoint: true,
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				FoodId: "Food001",
				SizeId: "Size001",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id: "Topping001",
					},
				},
				Amount: 10,
			},
		},
	}
	invoiceCreateWithUsePointButWithoutCustomer := invoicemodel.InvoiceCreate{
		CustomerId: nil,
		IsUsePoint: true,
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				FoodId: "Food001",
				SizeId: "Size001",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id: "Topping001",
					},
				},
				Amount: 10,
			},
		},
	}
	invoiceCreateWithoutCustomer := invoicemodel.InvoiceCreate{
		CustomerId: nil,
		IsUsePoint: false,
		InvoiceDetails: []invoicedetailmodel.InvoiceDetailCreate{
			{
				FoodId: "Food001",
				SizeId: "Size001",
				Toppings: &invoicedetailmodel.InvoiceDetailToppings{
					{
						Id: "Topping001",
					},
				},
				Amount: 10,
			},
		},
	}
	invalidInvoiceCreate := invoicemodel.InvoiceCreate{
		CustomerId: &invalidCustomerId,
	}
	shopGeneral := shopgeneralmodel.ShopGeneral{
		Name:                   "",
		Email:                  "",
		Phone:                  "",
		Address:                "",
		WifiPass:               "",
		AccumulatePointPercent: 0.001,
		UsePointPercent:        1,
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
			name: "Create invoice failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because data is invalid",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidInvoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not generate id",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because exist inactive product",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not handle data",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not get shop general",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not handle ingredient total amount",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredientTotalAmount",
						context.Background(),
						invoiceId,
						mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not find customer",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredientTotalAmount",
						context.Background(),
						invoiceId,
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"FindCustomer",
						context.Background(),
						customerId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because use point without customer info",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreateWithUsePointButWithoutCustomer,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreateWithUsePointButWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreateWithUsePointButWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleIngredientTotalAmount",
						context.Background(),
						invoiceId,
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice failed because can not save data to database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreateWithoutCustomer,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredientTotalAmount",
						context.Background(),
						invoiceId,
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleInvoice",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create invoice successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
				gen:       mockGen,
			},
			args: args{
				ctx:  context.Background(),
				data: &invoiceCreateWithoutCustomer,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceCreateFeatureCode).
					Return(true).
					Once()

				mockGen.
					On("GenerateId").
					Return(invoiceId, nil).
					Once()

				mockRepo.
					On(
						"HandleCheckPermissionStatus",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleData",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(nil).
					Once()

				mockRepo.
					On(
						"GetShopGeneral",
						context.Background()).
					Return(&shopGeneral, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredientTotalAmount",
						context.Background(),
						invoiceId,
						mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleInvoice",
						context.Background(),
						&invoiceCreateWithoutCustomer).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createInvoiceBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateInvoice(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"CreateInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"CreateInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
			}
		})
	}
}
