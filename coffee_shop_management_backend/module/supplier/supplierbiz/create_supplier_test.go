package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
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

type mockCreateSupplierRepo struct {
	mock.Mock
}

func (m *mockCreateSupplierRepo) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.SupplierCreate) error {
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
func (m *mockRequester) GetRole() rolemodel.Role {
	args := m.Called()
	return args.Get(0).(rolemodel.Role)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewCreateSupplierBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateSupplierRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateSupplierRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createSupplierBiz
	}{
		{
			name: "Create object has type CreateSupplierBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createSupplierBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateSupplierBiz(tt.args.gen, tt.args.repo, tt.args.requester); !reflect.DeepEqual(got, tt.want) {
				got := NewCreateSupplierBiz(
					tt.args.gen,
					tt.args.repo,
					tt.args.requester,
				)

				assert.Equal(t, tt.want, got, "NewCreateSupplierBiz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createSupplierBiz_CreateSupplier(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateSupplierRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *suppliermodel.SupplierCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateSupplierRepo)
	mockRequest := new(mockRequester)

	validEmail := "a@gmail.com"
	validPhone := "0123456789"
	validId := "0123456789"
	supplierCreate := suppliermodel.SupplierCreate{
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
			name: "Create supplier failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &suppliermodel.SupplierCreate{
					Id:   nil,
					Name: "",
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", supplierCreate.Id).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier failed because can not save to database",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", supplierCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplier",
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
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", supplierCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplier",
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
			biz := &createSupplierBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateSupplier(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
