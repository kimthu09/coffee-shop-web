package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateExportNoteStore struct {
	mock.Mock
}

func (m *mockCreateExportNoteStore) CreateExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreatExportNoteDetailStore struct {
	mock.Mock
}

func (m *mockCreatExportNoteDetailStore) CreateListExportNoteDetail(
	ctx context.Context,
	data []exportnotedetailmodel.ExportNoteDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockUpdateIngredientStore struct {
	mock.Mock
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
func (m *mockUpdateIngredientStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewCreateExportNoteRepo(t *testing.T) {
	type args struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	tests := []struct {
		name string
		args args
		want *createExportNoteRepo
	}{
		{
			name: "Create object has type CreateExportNoteRepo",
			args: args{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			want: &createExportNoteRepo{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateExportNoteRepo(
				tt.args.exportNoteStore,
				tt.args.exportNoteDetailStore,
				tt.args.ingredientStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateExportNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createExportNoteRepo_HandleExportNote(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
	}
	type args struct {
		ctx  context.Context
		data *exportnotemodel.ExportNoteCreate
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	mockExportNoteId := "mockId"
	mockDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "mockIng1",
			AmountExport: 10,
		},
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "mockIng2",
			AmountExport: 10,
		},
	}
	reason := exportnotemodel.OutOfDate
	createdBy := "user001"
	mockExportNoteCreate := exportnotemodel.ExportNoteCreate{
		Id:                &mockExportNoteId,
		CreatedBy:         createdBy,
		Reason:            &reason,
		ExportNoteDetails: mockDetails,
	}
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle export note failed because can not save export note to database",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockExportNote.
					On(
						"CreateExportNote",
						context.Background(),
						&mockExportNoteCreate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle export note failed because can not save export note detail to database",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockExportNote.
					On(
						"CreateExportNote",
						context.Background(),
						&mockExportNoteCreate,
					).
					Return(nil).
					Once()

				mockExportNoteDetail.
					On(
						"CreateListExportNoteDetail",
						context.Background(),
						mockDetails,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle export note successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockExportNote.
					On(
						"CreateExportNote",
						context.Background(),
						&mockExportNoteCreate,
					).
					Return(nil).
					Once()

				mockExportNoteDetail.
					On(
						"CreateListExportNoteDetail",
						context.Background(),
						mockDetails,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
			}

			tt.mock()

			err := repo.HandleExportNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleExportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleExportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createExportNoteRepo_HandleIngredientTotalAmount(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
	}
	type args struct {
		ctx                             context.Context
		exportNoteId                    string
		ingredientTotalAmountNeedUpdate map[string]int
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)

	ingredientId1 := "ingredientId1"
	mockMap := map[string]int{
		ingredientId1: 12,
	}
	outOfStockIngredient1 := ingredientmodel.Ingredient{
		Id:     ingredientId1,
		Amount: 12 - 1,
	}
	ingredient1 := ingredientmodel.Ingredient{
		Id:     ingredientId1,
		Amount: 12 + 1,
	}
	ingredientUpdate1 := ingredientmodel.IngredientUpdateAmount{Amount: -12}
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
			name: "Handle ingredient total amount failed because can not get ingredient",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId1},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount failed because exist ingredient out of stock",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId1},
						moreKeys,
					).
					Return(&outOfStockIngredient1, nil).
					Once()

			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount failed because can not update ingredient",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId1},
						moreKeys,
					).
					Return(&ingredient1, nil).
					Once()

				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientId1,
						&ingredientUpdate1,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"FindIngredient",
						context.Background(),
						map[string]interface{}{"id": ingredientId1},
						moreKeys,
					).
					Return(&ingredient1, nil).
					Once()

				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientId1,
						&ingredientUpdate1,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
			}

			tt.mock()

			err := repo.HandleIngredientTotalAmount(tt.args.ctx, tt.args.exportNoteId, tt.args.ingredientTotalAmountNeedUpdate)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
