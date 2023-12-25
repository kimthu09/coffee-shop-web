package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
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

type mockCreateCategoryStore struct {
	mock.Mock
}

func (m *mockCreateCategoryStore) CreateCategory(
	ctx context.Context,
	data *categorymodel.CategoryCreate) error {
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

func TestNewCreateCategoryBiz(t *testing.T) {
	type args struct {
		generator generator.IdGenerator
		store     CreateCategoryStore
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockStore := new(mockCreateCategoryStore)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createCategoryBiz
	}{
		{
			name: "Create object has type CreateCategoryBiz",
			args: args{
				generator: mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			want: &createCategoryBiz{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateCategoryBiz(
				tt.args.generator,
				tt.args.store,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateCategoryBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createCategoryBiz_CreateCategory(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		store     CreateCategoryStore
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *categorymodel.CategoryCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockStore := new(mockCreateCategoryStore)
	mockRequest := new(mockRequester)

	id := mock.Anything
	categoryCreate := categorymodel.CategoryCreate{
		Name:        mock.Anything,
		Description: mock.Anything,
	}
	invalidCategoryCreate := categorymodel.CategoryCreate{
		Name: "",
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
			name: "Create category failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create category failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidCategoryCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create category failed because can not generate id",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create category failed because can not save data to database",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(id, nil).
					Once()

				categoryCreate.Id = id

				mockStore.
					On(
						"CreateCategory",
						context.Background(),
						&categoryCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create category successfully",
			fields: fields{
				gen:       mockGenerator,
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(id, nil).
					Once()

				categoryCreate.Id = id

				mockStore.
					On(
						"CreateCategory",
						context.Background(),
						&categoryCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createCategoryBiz{
				gen:       tt.fields.gen,
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateCategory(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
