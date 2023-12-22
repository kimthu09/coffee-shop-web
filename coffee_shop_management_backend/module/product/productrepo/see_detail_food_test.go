package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeDetailFoodStore struct {
	mock.Mock
}

func (m *mockSeeDetailFoodStore) FindFood(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productmodel.Food, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Food), args.Error(1)
}

func TestNewSeeDetailFoodRepo(t *testing.T) {
	type args struct {
		store SeeDetailFoodStore
	}

	mockStore := new(mockSeeDetailFoodStore)

	tests := []struct {
		name string
		args args
		want *seeDetailFoodRepo
	}{
		{
			name: "Create object has type seeDetailFoodRepo",
			args: args{
				store: mockStore,
			},
			want: &seeDetailFoodRepo{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailFoodRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewSeeDetailFoodRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeDetailFoodRepo_SeeDetailFood(t *testing.T) {
	type fields struct {
		store SeeDetailFoodStore
	}
	type args struct {
		ctx    context.Context
		foodId string
	}

	mockStore := new(mockSeeDetailFoodStore)

	foodId := "Food001"

	mockErr := errors.New(mock.Anything)
	moreKeys := []string{"FoodSizes.Recipe.Details.Ingredient", "FoodCategories.Category"}
	mockFood := &productmodel.Food{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Food
		wantErr bool
	}{
		{
			name: "See detail food failed because can not get data from the mockStore",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				foodId: foodId,
			},
			mock: func() {
				mockStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{"id": foodId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail food successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				foodId: foodId,
			},
			mock: func() {
				mockStore.
					On("FindFood",
						context.Background(),
						map[string]interface{}{"id": foodId},
						moreKeys).
					Return(mockFood, nil).
					Once()
			},
			want:    mockFood,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeDetailFoodRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.SeeDetailFood(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeDetailFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeDetailFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeDetailFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
