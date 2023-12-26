package importnoterepo

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeStatusImportNoteStore struct {
	mock.Mock
}

func (m *mockChangeStatusImportNoteStore) FindImportNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*importnotemodel.ImportNote, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*importnotemodel.ImportNote), args.Error(1)
}
func (m *mockChangeStatusImportNoteStore) UpdateImportNote(
	ctx context.Context,
	id string,
	data *importnotemodel.ImportNoteUpdate) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockGetImportNoteDetailStore struct {
	mock.Mock
}

func (m *mockGetImportNoteDetailStore) FindListImportNoteDetail(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotedetailmodel.ImportNoteDetail),
		args.Error(1)
}

type mockUpdateAmountIngredientStore struct {
	mock.Mock
}

func (m *mockUpdateAmountIngredientStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockUpdateDebtOfSupplierStore struct {
	mock.Mock
}

func (m *mockUpdateDebtOfSupplierStore) FindSupplier(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*suppliermodel.Supplier, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*suppliermodel.Supplier), args.Error(1)
}
func (m *mockUpdateDebtOfSupplierStore) UpdateSupplierDebt(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateDebt) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockCreateSupplierDebtStorage struct {
	mock.Mock
}

func (m *mockCreateSupplierDebtStorage) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewChangeStatusImportNoteRepo(t *testing.T) {
	type args struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	tests := []struct {
		name string
		args args
		want *changeStatusImportNoteRepo
	}{
		{
			name: "Create object has type ChangeStatusImportNoteRepo",
			args: args{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			want: &changeStatusImportNoteRepo{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusImportNoteRepo(
				tt.args.importNoteStore,
				tt.args.importNoteDetailStore,
				tt.args.ingredientStore,
				tt.args.supplierStore,
				tt.args.supplierDebtStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewChangeStatusImportNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_changeStatusImportNoteRepo_CreateSupplierDebt(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}
	type args struct {
		ctx            context.Context
		supplierDebtId string
		importNote     *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	supplierId := mock.Anything

	conditionFindSupplier := map[string]interface{}{"id": supplierId}
	var moreKeys []string
	supplier := suppliermodel.Supplier{
		Debt: -30000,
	}

	importNoteId := mock.Anything
	status := importnotemodel.Done
	debtType := enum.Debt
	importNote := importnotemodel.ImportNoteUpdate{
		ClosedBy:   mock.Anything,
		Id:         importNoteId,
		SupplierId: supplierId,
		TotalPrice: 10000,
		Status:     &status,
	}
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         mock.Anything,
		SupplierId: supplierId,
		Amount:     -10000,
		AmountLeft: -40000,
		DebtType:   &debtType,
		CreatedBy:  mock.Anything,
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
			name: "Create supplier debt failed because can not get debt of supplier",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:            context.Background(),
				supplierDebtId: supplierId,
				importNote:     &importNote,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"FindSupplier",
						context.Background(),
						conditionFindSupplier,
						moreKeys).
					Return(
						nil,
						mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier debt failed because can not store to database",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:            context.Background(),
				supplierDebtId: supplierId,
				importNote:     &importNote,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"FindSupplier",
						context.Background(),
						conditionFindSupplier,
						moreKeys).
					Return(
						&supplier,
						nil).
					Once()

				mockSupplierDebtStore.
					On(
						"CreateSupplierDebt",
						context.Background(),
						&supplierDebtCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier debt successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:            context.Background(),
				supplierDebtId: supplierId,
				importNote:     &importNote,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"FindSupplier",
						context.Background(),
						conditionFindSupplier,
						moreKeys).
					Return(
						&supplier,
						nil).
					Once()

				mockSupplierDebtStore.
					On(
						"CreateSupplierDebt",
						context.Background(),
						&supplierDebtCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.CreateSupplierDebt(tt.args.ctx, tt.args.supplierDebtId, tt.args.importNote)
			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_FindImportNote(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	mockErr := errors.New(mock.Anything)

	importNote := importnotemodel.ImportNote{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "Get import note failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"FindImportNote",
						context.Background(),
						map[string]interface{}{"id": importNoteId},
						mock.Anything).
					Return(
						nil,
						mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get import note successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"FindImportNote",
						context.Background(),
						map[string]interface{}{"id": importNoteId},
						mock.Anything).
					Return(
						&importNote,
						nil).
					Once()
			},
			want:    &importNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			got, err := repo.FindImportNote(tt.args.ctx, tt.args.importNoteId)
			if tt.wantErr {
				assert.NotNil(t, err, "FindImportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindImportNote() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindImportNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_FindListImportNoteDetail(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	mockErr := errors.New(mock.Anything)
	importNotes := []importnotedetailmodel.ImportNoteDetail{
		{
			ImportNoteId: importNoteId,
			IngredientId: mock.Anything,
			Price:        0,
			AmountImport: 0,
		},
		{
			ImportNoteId: importNoteId,
			IngredientId: mock.Anything,
			Price:        0,
			AmountImport: 0,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotedetailmodel.ImportNoteDetail
		wantErr bool
	}{
		{
			name: "Get import note detail failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockImportNoteDetail.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						map[string]interface{}{"importNoteId": importNoteId},
						mock.Anything).
					Return(
						nil,
						mockErr).
					Once()
			},
			want:    importNotes,
			wantErr: true,
		},
		{
			name: "Get import note detail successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockImportNoteDetail.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						map[string]interface{}{"importNoteId": importNoteId},
						mock.Anything).
					Return(
						importNotes,
						nil).
					Once()
			},
			want:    importNotes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			got, err := repo.FindListImportNoteDetail(tt.args.ctx, tt.args.importNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListImportNoteDetail() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_UpdateDebtSupplier(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}

	type args struct {
		ctx        context.Context
		importNote *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNote := importnotemodel.ImportNoteUpdate{
		SupplierId: mock.Anything,
		TotalPrice: 100,
	}

	amount := -importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &amount,
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
			name: "Update debt supplier failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:        context.Background(),
				importNote: &importNote,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"UpdateSupplierDebt",
						context.Background(),
						importNote.SupplierId,
						&supplierUpdateDebt).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update debt supplier successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:        context.Background(),
				importNote: &importNote,
			},
			mock: func() {
				mockSupplierStore.
					On(
						"UpdateSupplierDebt",
						context.Background(),
						importNote.SupplierId,
						&supplierUpdateDebt).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.UpdateDebtSupplier(tt.args.ctx, tt.args.importNote)
			if tt.wantErr {
				assert.NotNil(t, err, "UpdateDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_UpdateImportNote(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}
	type args struct {
		ctx          context.Context
		importNoteId string
		data         *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	status := importnotemodel.Done
	importNoteUpdate := importnotemodel.ImportNoteUpdate{
		ClosedBy:   mock.Anything,
		Id:         importNoteId,
		SupplierId: mock.Anything,
		TotalPrice: 10000,
		Status:     &status,
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
			name: "Update import note failed",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
				data:         &importNoteUpdate,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"UpdateImportNote",
						context.Background(),
						importNoteId,
						&importNoteUpdate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update import note successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
				data:         &importNoteUpdate,
			},
			mock: func() {
				mockImportNoteStore.
					On(
						"UpdateImportNote",
						context.Background(),
						importNoteId,
						&importNoteUpdate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.UpdateImportNote(tt.args.ctx, tt.args.importNoteId, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "UpdateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_HandleIngredient(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStore
		importNoteDetailStore GetImportNoteDetailStore
		ingredientStore       UpdateAmountIngredientStore
		supplierStore         UpdateDebtOfSupplierStore
		supplierDebtStore     CreateSupplierDebtStore
	}
	type args struct {
		ctx                             context.Context
		importNoteId                    string
		ingredientTotalAmountNeedUpdate map[string]int
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := "importNote123"

	ingredientTotalAmountNeedUpdate := map[string]int{
		"ingredient1": 10,
	}
	ingredientUpdate1 := ingredientmodel.IngredientUpdateAmount{Amount: 10}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle ingredient failed because can not update amount of ingredient 1",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:                             context.Background(),
				importNoteId:                    importNoteId,
				ingredientTotalAmountNeedUpdate: ingredientTotalAmountNeedUpdate,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						"ingredient1",
						&ingredientUpdate1).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient changes successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:                             context.Background(),
				importNoteId:                    importNoteId,
				ingredientTotalAmountNeedUpdate: ingredientTotalAmountNeedUpdate,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						"ingredient1",
						&ingredientUpdate1).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.HandleIngredient(tt.args.ctx, tt.args.ingredientTotalAmountNeedUpdate)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
