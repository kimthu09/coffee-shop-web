package importnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"fmt"
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

type mockUpdateOrCreateIngredientDetailStore struct {
	mock.Mock
}

func (m *mockUpdateOrCreateIngredientDetailStore) FindIngredientDetail(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*ingredientdetailmodel.IngredientDetail, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientdetailmodel.IngredientDetail),
		args.Error(1)
}
func (m *mockUpdateOrCreateIngredientDetailStore) UpdateIngredientDetail(
	ctx context.Context,
	ingredientId string,
	expiryDate string,
	data *ingredientdetailmodel.IngredientDetailUpdate) error {
	args := m.Called(ctx, ingredientId, expiryDate, data)
	return args.Error(0)
}
func (m *mockUpdateOrCreateIngredientDetailStore) CreateIngredientDetail(
	ctx context.Context,
	data *ingredientdetailmodel.IngredientDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
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

func (m *mockUpdateDebtOfSupplierStore) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*float32, error) {
	args := m.Called(ctx, supplierId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
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
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
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
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			want: &changeStatusImportNoteRepo{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
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
				tt.args.ingredientDetailStore,
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
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx            context.Context
		supplierDebtId string
		importNote     *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	supplierId := mock.Anything
	importNoteId := mock.Anything
	status := importnotemodel.Done
	importDebt := float32(10000)
	currentDebt := float32(30000)
	debtType := enum.Debt
	importNote := importnotemodel.ImportNoteUpdate{
		CloseBy:    mock.Anything,
		Id:         importNoteId,
		SupplierId: supplierId,
		TotalPrice: importDebt,
		Status:     &status,
	}
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         mock.Anything,
		SupplierId: supplierId,
		Amount:     importDebt,
		AmountLeft: currentDebt + importDebt,
		DebtType:   &debtType,
		CreateBy:   mock.Anything,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
						"GetDebtSupplier",
						context.Background(),
						supplierId,
						mock.Anything).
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
				ingredientDetailStore: mockIngredientDetailStore,
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
						"GetDebtSupplier",
						context.Background(),
						supplierId,
						mock.Anything).
					Return(
						&currentDebt,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
						"GetDebtSupplier",
						context.Background(),
						supplierId,
						mock.Anything).
					Return(
						&currentDebt,
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.CreateSupplierDebt(tt.args.ctx, tt.args.supplierDebtId, tt.args.importNote)
			if tt.wantErr {
				assert.NotNil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_FindImportNote(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
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
				ingredientDetailStore: mockIngredientDetailStore,
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
			want:    &importNote,
			wantErr: true,
		},
		{
			name: "Get import note successfully",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
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
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	mockErr := errors.New(mock.Anything)
	importNotes := []importnotedetailmodel.ImportNoteDetail{
		{
			ImportNoteId: importNoteId,
			IngredientId: mock.Anything,
			ExpiryDate:   mock.Anything,
			Price:        0,
			AmountImport: 0,
		},
		{
			ImportNoteId: importNoteId,
			IngredientId: mock.Anything,
			ExpiryDate:   mock.Anything,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
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

func Test_changeStatusImportNoteRepo_HandleIngredientDetails(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx               context.Context
		importNoteDetails []importnotedetailmodel.ImportNoteDetail
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	importNoteDetail := importnotedetailmodel.ImportNoteDetail{
		ImportNoteId: importNoteId,
		IngredientId: mock.Anything,
		ExpiryDate:   mock.Anything,
		Price:        0,
		AmountImport: 0,
	}
	importNoteDetails := []importnotedetailmodel.ImportNoteDetail{
		importNoteDetail,
		importNoteDetail,
	}
	ingredientDetail := ingredientdetailmodel.IngredientDetail{}
	createIngredientDetail := ingredientdetailmodel.IngredientDetailCreate{
		IngredientId: importNoteDetail.IngredientId,
		ExpiryDate:   importNoteDetail.ExpiryDate,
		Amount:       importNoteDetail.AmountImport,
	}
	updateIngredientDetail := ingredientdetailmodel.IngredientDetailUpdate{
		IngredientId: importNoteDetail.IngredientId,
		ExpiryDate:   importNoteDetail.ExpiryDate,
		Amount:       importNoteDetail.AmountImport,
	}

	mockNormalErr := errors.New(mock.Anything)
	mockRecordNotFoundErr := common.ErrRecordNotFound()

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle ingredient detail failed because database has some error",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:               context.Background(),
				importNoteDetails: importNoteDetails,
			},
			mock: func() {
				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						nil,
						mockNormalErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient detail failed because can not update ingredient detail",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:               context.Background(),
				importNoteDetails: importNoteDetails,
			},
			mock: func() {
				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						nil,
						mockRecordNotFoundErr).
					Once()

				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						&ingredientDetail,
						nil).
					Once()

				mockIngredientDetailStore.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						updateIngredientDetail.IngredientId,
						updateIngredientDetail.ExpiryDate,
						&updateIngredientDetail).
					Return(mockNormalErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient detail failed because can not create ingredient detail",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:               context.Background(),
				importNoteDetails: importNoteDetails,
			},
			mock: func() {
				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						nil,
						mockRecordNotFoundErr).
					Once()

				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						&ingredientDetail,
						nil).
					Once()

				mockIngredientDetailStore.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						updateIngredientDetail.IngredientId,
						updateIngredientDetail.ExpiryDate,
						&updateIngredientDetail).
					Return(nil).
					Once()

				mockIngredientDetailStore.
					On(
						"CreateIngredientDetail",
						context.Background(),
						&createIngredientDetail).
					Return(mockNormalErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient detail successfully because can not create ingredient detail",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:               context.Background(),
				importNoteDetails: importNoteDetails,
			},
			mock: func() {
				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						nil,
						mockRecordNotFoundErr).
					Once()

				mockIngredientDetailStore.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mock.Anything,
							"expiryDate":   mock.Anything,
						},
						mock.Anything).
					Return(
						&ingredientDetail,
						nil).
					Once()

				mockIngredientDetailStore.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						updateIngredientDetail.IngredientId,
						updateIngredientDetail.ExpiryDate,
						&updateIngredientDetail).
					Return(nil).
					Once()

				mockIngredientDetailStore.
					On(
						"CreateIngredientDetail",
						context.Background(),
						&createIngredientDetail).
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.HandleIngredientDetails(tt.args.ctx, tt.args.importNoteDetails)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredientDetails() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredientDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_HandleIngredientTotalAmount(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx                             context.Context
		ingredientTotalAmountNeedUpdate map[string]float32
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	ingredientKey1 := mock.Anything
	ingredientAmount1 := float32(30)
	ingredientKey2 := mock.Anything
	ingredientAmount2 := float32(60)
	ingredientTotalAmountNeedUpdate := map[string]float32{
		ingredientKey1: ingredientAmount1,
		ingredientKey2: ingredientAmount2,
	}
	ingredientUpdate1 := ingredientmodel.IngredientUpdateAmount{
		Amount: ingredientAmount1,
	}
	ingredientUpdate2 := ingredientmodel.IngredientUpdateAmount{
		Amount: ingredientAmount2,
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
			name: "Handle ingredient total amount failed because can not update amount of ingredient",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: ingredientTotalAmountNeedUpdate,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientKey2,
						&ingredientUpdate2).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle ingredient total amount failed because can not update amount of ingredient",
			fields: fields{
				importNoteStore:       mockImportNoteStore,
				importNoteDetailStore: mockImportNoteDetail,
				ingredientStore:       mockIngredientStore,
				ingredientDetailStore: mockIngredientDetailStore,
				supplierStore:         mockSupplierStore,
				supplierDebtStore:     mockSupplierDebtStore,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: ingredientTotalAmountNeedUpdate,
			},
			mock: func() {
				mockIngredientStore.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientKey1,
						&ingredientUpdate1).
					Return(nil).
					Once()

				mockIngredientStore.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						ingredientKey2,
						&ingredientUpdate2).
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.HandleIngredientTotalAmount(tt.args.ctx, tt.args.ingredientTotalAmountNeedUpdate)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_UpdateDebtSupplier(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}

	type args struct {
		ctx        context.Context
		importNote *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNote := importnotemodel.ImportNoteUpdate{
		SupplierId: mock.Anything,
		TotalPrice: 100,
	}
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &importNote.TotalPrice,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
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
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx          context.Context
		importNoteId string
		data         *importnotemodel.ImportNoteUpdate
	}

	mockImportNoteStore := new(mockChangeStatusImportNoteStore)
	mockImportNoteDetail := new(mockGetImportNoteDetailStore)
	mockIngredientStore := new(mockUpdateAmountIngredientStore)
	mockIngredientDetailStore := new(mockUpdateOrCreateIngredientDetailStore)
	mockSupplierStore := new(mockUpdateDebtOfSupplierStore)
	mockSupplierDebtStore := new(mockCreateSupplierDebtStorage)

	importNoteId := mock.Anything
	status := importnotemodel.Done
	importNoteUpdate := importnotemodel.ImportNoteUpdate{
		CloseBy:    mock.Anything,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: mockIngredientDetailStore,
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
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.UpdateImportNote(tt.args.ctx, tt.args.importNoteId, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CheckIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_changeStatusImportNoteRepo_createIngredientDetails(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx                      context.Context
		createdIngredientDetails []ingredientdetailmodel.IngredientDetailCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}
			tt.wantErr(t, repo.createIngredientDetails(tt.args.ctx, tt.args.createdIngredientDetails), fmt.Sprintf("createIngredientDetails(%v, %v)", tt.args.ctx, tt.args.createdIngredientDetails))
		})
	}
}

func Test_changeStatusImportNoteRepo_updateIngredientDetails(t *testing.T) {
	type fields struct {
		importNoteStore       ChangeStatusImportNoteStorage
		importNoteDetailStore GetImportNoteDetailStorage
		ingredientStore       UpdateAmountIngredientStorage
		ingredientDetailStore UpdateOrCreateIngredientDetailStorage
		supplierStore         UpdateDebtOfSupplierStorage
		supplierDebtStore     CreateSupplierDebtStorage
	}
	type args struct {
		ctx                      context.Context
		updatedIngredientDetails []ingredientdetailmodel.IngredientDetailUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusImportNoteRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
				supplierStore:         tt.fields.supplierStore,
				supplierDebtStore:     tt.fields.supplierDebtStore,
			}
			tt.wantErr(t, repo.updateIngredientDetails(tt.args.ctx, tt.args.updatedIngredientDetails), fmt.Sprintf("updateIngredientDetails(%v, %v)", tt.args.ctx, tt.args.updatedIngredientDetails))
		})
	}
}
