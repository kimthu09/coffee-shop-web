package inventorychecknoterepo

import (
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateInventoryCheckNoteStore struct {
	mock.Mock
}

func (m *mockCreateInventoryCheckNoteStore) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateInventoryCheckNoteDetailStore struct {
	mock.Mock
}

func (m *mockCreateInventoryCheckNoteDetailStore) CreateListInventoryCheckNoteDetail(
	ctx context.Context,
	data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockUpdateIngredientStore struct {
	mock.Mock
}

func (m *mockUpdateIngredientStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}
func (m *mockUpdateIngredientStore) FindIngredient(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*ingredientmodel.Ingredient, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientmodel.Ingredient), args.Error(1)
}

func TestNewCreateInventoryCheckNoteRepo(t *testing.T) {
	type args struct {
		inventoryCheckNoteStore       CreateInventoryCheckNoteStore
		inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
		ingredientStore               UpdateIngredientStore
	}

	mockInventoryCheckNote := new(mockCreateInventoryCheckNoteStore)
	mockInventoryCheckNoteDetail := new(mockCreateInventoryCheckNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	tests := []struct {
		name string
		args args
		want *createInventoryCheckNoteRepo
	}{
		{
			name: "Create object has type CreateExportNoteRepo",
			args: args{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			want: &createInventoryCheckNoteRepo{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateInventoryCheckNoteRepo(
				tt.args.inventoryCheckNoteStore,
				tt.args.inventoryCheckNoteDetailStore,
				tt.args.ingredientStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateInventoryCheckNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createInventoryCheckNoteRepo_HandleIngredientAmount(t *testing.T) {
	type fields struct {
		inventoryCheckNoteStore       CreateInventoryCheckNoteStore
		inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
		ingredientStore               UpdateIngredientStore
	}
	type args struct {
		ctx  context.Context
		data *inventorychecknotemodel.InventoryCheckNoteCreate
	}

	mockInventoryCheckNote := new(mockCreateInventoryCheckNoteStore)
	mockInventoryCheckNoteDetail := new(mockCreateInventoryCheckNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	inventoryCheckNoteId := mock.Anything
	ingredientId := "ingredient1"
	inventoryCheckNoteCreate := inventorychecknotemodel.InventoryCheckNoteCreate{
		Id: &inventoryCheckNoteId,
		Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
			{
				InventoryCheckNoteId: inventoryCheckNoteId,
				IngredientId:         ingredientId,
				Difference:           -100,
			},
		},
	}
	ingredient := ingredientmodel.Ingredient{
		Id:     ingredientId,
		Amount: 120,
	}
	outOfStockIngredient := ingredientmodel.Ingredient{
		Id:     ingredientId,
		Amount: 80,
	}
	ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: -100}
	finalInventoryCheckNoteCreate := inventorychecknotemodel.InventoryCheckNoteCreate{
		Id:                &inventoryCheckNoteId,
		AmountDifferent:   -100,
		AmountAfterAdjust: 20,
		Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
			{
				InventoryCheckNoteId: inventoryCheckNoteId,
				IngredientId:         ingredientId,
				Initial:              120,
				Difference:           -100,
				Final:                20,
			},
		},
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle ingredient amount failed because can not get ingredient",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNoteCreate,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient amount failed because ingredient out of stock",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNoteCreate,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						moreKeys,
					).
					Return(&outOfStockIngredient, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient amount failed because can not update amount of ingredient",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNoteCreate,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						moreKeys,
					).
					Return(&ingredient, nil).
					Once()

				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientId,
						&ingredientUpdate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient amount successfully",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNoteCreate,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						moreKeys,
					).
					Return(&ingredient, nil).
					Once()

				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientId,
						&ingredientUpdate,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInventoryCheckNoteRepo{
				inventoryCheckNoteStore:       tt.fields.inventoryCheckNoteStore,
				inventoryCheckNoteDetailStore: tt.fields.inventoryCheckNoteDetailStore,
				ingredientStore:               tt.fields.ingredientStore,
			}

			tt.mock()

			err := repo.HandleIngredientAmount(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredientAmount() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredientAmount() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, *tt.args.data, finalInventoryCheckNoteCreate,
					"HandleIngredientAmount() param = %v, want %v",
					*tt.args.data, finalInventoryCheckNoteCreate)
			}
		})
	}
}

func Test_createInventoryCheckNoteRepo_HandleInventoryCheckNote(t *testing.T) {
	type fields struct {
		inventoryCheckNoteStore       CreateInventoryCheckNoteStore
		inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
		ingredientStore               UpdateIngredientStore
	}
	type args struct {
		ctx  context.Context
		data *inventorychecknotemodel.InventoryCheckNoteCreate
	}

	mockInventoryCheckNote := new(mockCreateInventoryCheckNoteStore)
	mockInventoryCheckNoteDetail := new(mockCreateInventoryCheckNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	inventoryNoteId := "note001"
	createdBy := "user001"
	ingredientId := "ingredient001"
	inventoryCheckNote := inventorychecknotemodel.InventoryCheckNoteCreate{
		Id:                &inventoryNoteId,
		AmountDifferent:   -100,
		AmountAfterAdjust: 220,
		CreatedBy:         createdBy,
		Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
			{
				InventoryCheckNoteId: inventoryNoteId,
				IngredientId:         ingredientId,
				Initial:              120,
				Difference:           -100,
				Final:                20,
			},
		},
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
			name: "Handle inventory check note failed because can not save inventory check note to database",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNote,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"CreateInventoryCheckNote",
						context.Background(),
						&inventoryCheckNote,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle inventory check note failed because can not save details to database",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNote,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"CreateInventoryCheckNote",
						context.Background(),
						&inventoryCheckNote,
					).
					Return(nil).
					Once()

				mockInventoryCheckNoteDetail.
					On(
						"CreateListInventoryCheckNoteDetail",
						context.Background(),
						inventoryCheckNote.Details,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle inventory check note successfully",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
				ingredientStore:               mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &inventoryCheckNote,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"CreateInventoryCheckNote",
						context.Background(),
						&inventoryCheckNote,
					).
					Return(nil).
					Once()

				mockInventoryCheckNoteDetail.
					On(
						"CreateListInventoryCheckNoteDetail",
						context.Background(),
						inventoryCheckNote.Details,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createInventoryCheckNoteRepo{
				inventoryCheckNoteStore:       tt.fields.inventoryCheckNoteStore,
				inventoryCheckNoteDetailStore: tt.fields.inventoryCheckNoteDetailStore,
				ingredientStore:               tt.fields.ingredientStore,
			}

			tt.mock()

			err := repo.HandleInventoryCheckNote(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
