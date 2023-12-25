package ingredientbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
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

type mockCreateIngredientStore struct {
	mock.Mock
}

func (m *mockCreateIngredientStore) CreateIngredient(
	ctx context.Context,
	data *ingredientmodel.IngredientCreate) error {
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

func TestNewCreateIngredientBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		store     CreateIngredientStore
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockStore := new(mockCreateIngredientStore)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createIngredientBiz
	}{
		{
			name: "Create object has type CreateExportNoteBiz",
			args: args{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			want: &createIngredientBiz{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateIngredientBiz(
				tt.args.gen,
				tt.args.store,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateIngredientBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createIngredientBiz_CreateIngredient(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		store     CreateIngredientStore
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *ingredientmodel.IngredientCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockStore := new(mockCreateIngredientStore)
	mockRequest := new(mockRequester)

	validId := "012345678901"
	measureType := enum.Weight
	roundedPrice := float32(0.001)
	ingredient := ingredientmodel.IngredientCreate{
		Id:          &validId,
		Name:        mock.Anything,
		MeasureType: &measureType,
		Price:       0.0006,
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
			name: "Create ingredient failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredient,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create ingredient failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &ingredientmodel.IngredientCreate{
					Name: "",
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create ingredient failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredient,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", ingredient.Id).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create ingredient failed because can not save data to database",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredient,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", ingredient.Id).
					Return(&validId, nil).
					Once()

				mockStore.
					On(
						"CreateIngredient",
						context.Background(),
						&ingredient).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		}, {
			name: "Create ingredient successfully",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredient,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", ingredient.Id).
					Return(&validId, nil).
					Once()

				mockStore.
					On(
						"CreateIngredient",
						context.Background(),
						&ingredient).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createIngredientBiz{
				gen:       tt.fields.gen,
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateIngredient(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateIngredient() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateIngredient() = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.args.data.Price, roundedPrice, "Param.Price = %v, want = %v", tt.args.data.Price, roundedPrice)
			}
		})
	}
}
