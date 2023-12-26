package ingredientbiz

import (
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAllIngredientStore struct {
	mock.Mock
}

func (m *mockGetAllIngredientStore) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.Ingredient, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]ingredientmodel.Ingredient), args.Error(1)
}

func TestNewGetAllIngredientBiz(t *testing.T) {
	type args struct {
		store GetAllIngredientStore
	}

	mockStore := new(mockGetAllIngredientStore)

	tests := []struct {
		name string
		args args
		want *getAllIngredientBiz
	}{
		{
			name: "Create object has type GetAllIngredientBiz",
			args: args{
				store: mockStore,
			},
			want: &getAllIngredientBiz{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllIngredientBiz(tt.args.store)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewGetAllIngredientBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_getAllIngredientBiz_GetAllUser(t *testing.T) {
	type fields struct {
		repo GetAllIngredientStore
	}
	type args struct {
		ctx context.Context
	}

	mockRepo := new(mockGetAllIngredientStore)
	listIngredients := []ingredientmodel.Ingredient{
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
		want    []ingredientmodel.Ingredient
		wantErr bool
	}{
		{
			name: "Get all ingredient failed because can not get data from database",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllIngredient", context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get all ingredient successfully",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRepo.
					On("GetAllIngredient", context.Background()).
					Return(listIngredients, nil).
					Once()
			},
			want:    listIngredients,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &getAllIngredientBiz{
				store: tt.fields.repo,
			}

			tt.mock()

			got, err := biz.GetAllIngredient(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"GetAllIngredient() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"GetAllIngredient() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"GetAllIngredient() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
