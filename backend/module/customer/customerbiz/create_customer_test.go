package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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

type mockCreateCustomerRepo struct {
	mock.Mock
}

func (m *mockCreateCustomerRepo) CreateCustomer(
	ctx context.Context,
	data *customermodel.CustomerCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockRequester struct {
	mock.Mock
}

func (m *mockRequester) GetUserId() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetEmail() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetRoleId() string {
	args := m.Called()
	return args.Get(0).(string)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewCreateCustomerBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateCustomerRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateCustomerRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createCustomerBiz
	}{
		{
			name: "Create object has type CreateCustomerBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createCustomerBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateCustomerBiz(
				tt.args.gen,
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateCustomerBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createCustomerBiz_CreateCustomer(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateCustomerRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *customermodel.CustomerCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateCustomerRepo)
	mockRequest := new(mockRequester)

	validEmail := "a@gmail.com"
	validPhone := "0123456789"
	validId := "0123456789"
	customerCreate := customermodel.CustomerCreate{
		Id:    nil,
		Name:  mock.Anything,
		Email: validEmail,
		Phone: validPhone,
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
			name: "Create customer failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create customer failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &customermodel.CustomerCreate{
					Id:   nil,
					Name: "",
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create customer failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", customerCreate.Id).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create customer failed because can not save to database",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", customerCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CreateCustomer",
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
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &customerCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", customerCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CreateCustomer",
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
			biz := &createCustomerBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateCustomer(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
