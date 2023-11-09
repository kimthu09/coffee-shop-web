package importnoterepo

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateImportNoteStore struct {
	mock.Mock
}

func (m *mockCreateImportNoteStore) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateImportNoteDetailStore struct {
	mock.Mock
}

func (m *mockCreateImportNoteDetailStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []importnotedetailmodel.ImportNoteDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockUpdatePriceIngredientStore struct {
	mock.Mock
}

func (m *mockUpdatePriceIngredientStore) FindIngredient(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*ingredientmodel.Ingredient, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientmodel.Ingredient), args.Error(1)
}

func (m *mockUpdatePriceIngredientStore) UpdatePriceIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdatePrice) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockCheckSupplierStore struct {
	mock.Mock
}

func (m *mockCheckSupplierStore) FindSupplier(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*suppliermodel.Supplier, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*suppliermodel.Supplier), nil
}

func TestNewCreateImportNoteRepo(t *testing.T) {
	type args struct {
		importNoteStore            CreateImportNoteStore
		importNoteDetailStore      CreateImportNoteDetailStore
		updatePriceIngredientStore UpdatePriceIngredientStore
		supplierStore              CheckSupplierStore
	}

	mockImportNoteStore := new(mockCreateImportNoteStore)
	mockImportNoteDetailStore := new(mockCreateImportNoteDetailStore)
	mockIngredientStore := new(mockUpdatePriceIngredientStore)
	mockSupplierStore := new(mockCheckSupplierStore)

	tests := []struct {
		name string
		args args
		want *createImportNoteRepo
	}{
		{
			name: "Create object has type CreateImportNoteRepo",
			args: args{
				importNoteStore:            mockImportNoteStore,
				importNoteDetailStore:      mockImportNoteDetailStore,
				updatePriceIngredientStore: mockIngredientStore,
				supplierStore:              mockSupplierStore,
			},
			want: &createImportNoteRepo{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateImportNoteRepo(
				tt.args.importNoteStore,
				tt.args.importNoteDetailStore,
				tt.args.updatePriceIngredientStore,
				tt.args.supplierStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateImportNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createImportNoteRepo_CheckIngredient(t *testing.T) {
	type fields struct {
		importNoteStore       CreateImportNoteStore
		importNoteDetailStore CreateImportNoteDetailStore
		ingredientStore       UpdatePriceIngredientStore
		supplierStore         CheckSupplierStore
	}
	type args struct {
		ctx          context.Context
		ingredientId string
	}

	mockImportNoteStore := new(mockCreateImportNoteStore)
	mockImportNoteDetailStore := new(mockCreateImportNoteDetailStore)
	mockIngredientStore := new(mockUpdatePriceIngredientStore)
	mockSupplierStore := new(mockCheckSupplierStore)

	ingredientId := mock.Anything
	mockErr := errors.New(mock.Anything)
	ingredient := ingredientmodel.Ingredient{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Check ingredient exist failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: ingredientId,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Check ingredient exist successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: ingredientId,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId},
						mock.Anything).
					Return(&ingredient, nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
			}

			tt.mock()

			err := repo.CheckIngredient(tt.args.ctx, tt.args.ingredientId)
			if tt.wantErr {
				assert.NotNil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createImportNoteRepo_CheckSupplier(t *testing.T) {
	type fields struct {
		importNoteStore       CreateImportNoteStore
		importNoteDetailStore CreateImportNoteDetailStore
		ingredientStore       UpdatePriceIngredientStore
		supplierStore         CheckSupplierStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
	}

	mockImportNoteStore := new(mockCreateImportNoteStore)
	mockImportNoteDetailStore := new(mockCreateImportNoteDetailStore)
	mockIngredientStore := new(mockUpdatePriceIngredientStore)
	mockSupplierStore := new(mockCheckSupplierStore)

	supplierId := mock.Anything
	mockErr := errors.New(mock.Anything)
	supplier := suppliermodel.Supplier{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Check supplier exist successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"FindSupplier",
						context.Background(),
						map[string]interface{}{"id": supplierId},
						mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Check supplier exist successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"FindSupplier",
						context.Background(),
						map[string]interface{}{"id": supplierId},
						mock.Anything).
					Return(&supplier, nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
			}

			tt.mock()

			err := repo.CheckSupplier(tt.args.ctx, tt.args.supplierId)

			if tt.wantErr {
				assert.NotNil(t, err, "CheckSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CheckSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createImportNoteRepo_HandleCreateImportNote(t *testing.T) {
	type fields struct {
		importNoteStore       CreateImportNoteStore
		importNoteDetailStore CreateImportNoteDetailStore
		ingredientStore       UpdatePriceIngredientStore
		supplierStore         CheckSupplierStore
	}
	type args struct {
		ctx  context.Context
		data *importnotemodel.ImportNoteCreate
	}

	mockImportNoteStore := new(mockCreateImportNoteStore)
	mockImportNoteDetailStore := new(mockCreateImportNoteDetailStore)
	mockIngredientStore := new(mockUpdatePriceIngredientStore)
	mockSupplierStore := new(mockCheckSupplierStore)

	id := mock.Anything
	importNoteCreate := importnotemodel.ImportNoteCreate{
		Id:         &id,
		TotalPrice: 0,
		SupplierId: mock.Anything,
		CreateBy:   mock.Anything,
		ImportNoteDetails: []importnotedetailmodel.ImportNoteDetailCreate{
			{
				ImportNoteId:   id,
				IngredientId:   mock.Anything,
				ExpiryDate:     mock.Anything,
				Price:          0,
				IsReplacePrice: false,
				AmountImport:   0,
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
			name: "Create import note failed because can not save import note to db",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"CreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not save import note details to db",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"CreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(nil).
					Once()

				mockImportNoteDetailStore.
					On(
						"CreateListImportNoteDetail",
						context.Background(),
						importNoteCreate.ImportNoteDetails).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"CreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(nil).
					Once()

				mockImportNoteDetailStore.
					On(
						"CreateListImportNoteDetail",
						context.Background(),
						importNoteCreate.ImportNoteDetails).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
			}

			tt.mock()

			err := repo.HandleCreateImportNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleCreateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleCreateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createImportNoteRepo_UpdatePriceIngredient(t *testing.T) {
	type fields struct {
		importNoteStore       CreateImportNoteStore
		importNoteDetailStore CreateImportNoteDetailStore
		ingredientStore       UpdatePriceIngredientStore
		supplierStore         CheckSupplierStore
	}
	type args struct {
		ctx          context.Context
		ingredientId string
		price        float32
	}

	mockImportNoteStore := new(mockCreateImportNoteStore)
	mockImportNoteDetailStore := new(mockCreateImportNoteDetailStore)
	mockIngredientStore := new(mockUpdatePriceIngredientStore)
	mockSupplierStore := new(mockCheckSupplierStore)

	id := mock.Anything
	price := float32(0)
	mockErr := errors.New(mock.Anything)
	ingredientUpdatePrice := ingredientmodel.IngredientUpdatePrice{
		Price: &price,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update price ingredient failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: id,
				price:        price,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdatePriceIngredient",
						context.Background(),
						id,
						&ingredientUpdatePrice).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update price ingredient successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetailStore,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: id,
				price:        price,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdatePriceIngredient",
						context.Background(),
						id,
						&ingredientUpdatePrice).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
			}

			tt.mock()

			err := repo.UpdatePriceIngredient(tt.args.ctx, tt.args.ingredientId, tt.args.price)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdatePriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdatePriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
