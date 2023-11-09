package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
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

func (m *mockUpdateIngredientStore) GetPriceIngredient(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*float32, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}
func (m *mockUpdateIngredientStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockUpdateIngredientDetailStore struct {
	mock.Mock
}

func (m *mockUpdateIngredientDetailStore) FindIngredientDetail(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*ingredientdetailmodel.IngredientDetail, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientdetailmodel.IngredientDetail),
		args.Error(1)
}
func (m *mockUpdateIngredientDetailStore) UpdateIngredientDetail(
	ctx context.Context,
	ingredientId string,
	expiryDate string,
	data *ingredientdetailmodel.IngredientDetailUpdate,
) error {
	args := m.Called(ctx, ingredientId, expiryDate, data)
	return args.Error(0)
}

func TestNewCreateExportNoteRepo(t *testing.T) {
	type args struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

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
				ingredientDetailStore: mockIngredientDetail,
			},
			want: &createExportNoteRepo{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateExportNoteRepo(
				tt.args.exportNoteStore,
				tt.args.exportNoteDetailStore,
				tt.args.ingredientStore,
				tt.args.ingredientDetailStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateExportNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createExportNoteRepo_GetPriceIngredient(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx          context.Context
		ingredientId string
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockIngredientId := "mockId"
	mockPrice := float32(100)
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *float32
		wantErr bool
	}{
		{
			name: "Get price of ingredient successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: mockIngredientId,
			},
			mock: func() {
				mockIngredient.
					On(
						"GetPriceIngredient",
						context.Background(),
						map[string]interface{}{"id": mockIngredientId},
						mock.Anything).
					Return(
						&mockPrice,
						nil).
					Once()
			},
			want:    &mockPrice,
			wantErr: false,
		},
		{
			name: "Get price of ingredient failed with error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: mockIngredientId,
			},
			mock: func() {
				mockIngredient.
					On(
						"GetPriceIngredient",
						context.Background(),
						map[string]interface{}{"id": mockIngredientId},
						mock.Anything).
					Return(
						nil,
						mockErr,
					).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			got, err := repo.GetPriceIngredient(tt.args.ctx, tt.args.ingredientId)

			if tt.wantErr {
				assert.NotNil(t, err, "GetPriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetPriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetPriceIngredient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createExportNoteRepo_HandleExportNote(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *exportnotemodel.ExportNoteCreate
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockExportNoteId := "mockId"
	mockDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "mockIng1",
			ExpiryDate:   "mockDate",
			AmountExport: 10,
		},
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "mockIng2",
			ExpiryDate:   "mockDate",
			AmountExport: 10,
		},
	}
	mockExportNoteCreate := exportnotemodel.ExportNoteCreate{
		Id:                &mockExportNoteId,
		TotalPrice:        float32(1000),
		CreateBy:          "mockCreateBy",
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
			name: "Create export note and export note detail successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
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
		{
			name: "Create export note failed and return error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
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
			name: "Create export note successfully but " +
				"create export note detail failed then return error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
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
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx                             context.Context
		ingredientTotalAmountNeedUpdate map[string]float32
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockMap := map[string]float32{
		"ingredient1": 12,
		"ingredient2": 30,
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
			name: "Update total amount of ingredient successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						mock.Anything,
						mock.Anything,
					).
					Return(nil).
					Times(len(mockMap))
			},
			wantErr: false,
		},
		{
			name: "Update total amount of ingredient failed and return error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						mock.Anything,
						mock.Anything,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
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

func Test_createExportNoteRepo_checkIngredientDetail(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *exportnotedetailmodel.ExportNoteDetailCreate
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockExportNoteId := "mockId"
	mockIngredientId := "mockIngId"
	mockExpiryData := "mockDate"
	mockDetail := exportnotedetailmodel.ExportNoteDetailCreate{
		ExportNoteId: mockExportNoteId,
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		AmountExport: 10,
	}
	mockIngredientDetailHasBigAmountStock := ingredientdetailmodel.IngredientDetail{
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		Amount:       100,
	}
	mockIngredientDetailHasSmallAmountStock := ingredientdetailmodel.IngredientDetail{
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		Amount:       5,
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
			name: "ExportNoteDetail check found no error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(&mockIngredientDetailHasBigAmountStock, nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "ExportNoteDetail check found amount export is over stock error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(&mockIngredientDetailHasSmallAmountStock, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "ExportNoteDetail check found database error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(nil, mockErr)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.checkIngredientDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "checkIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "checkIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createExportNoteRepo_updateIngredientDetail(t *testing.T) {
	type fields struct {
		exportNoteStore       CreateExportNoteStore
		exportNoteDetailStore CreateExportNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *exportnotedetailmodel.ExportNoteDetailCreate
	}

	mockExportNote := new(mockCreateExportNoteStore)
	mockExportNoteDetail := new(mockCreatExportNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockDetail := exportnotedetailmodel.ExportNoteDetailCreate{
		ExportNoteId: "mockId",
		IngredientId: "mockIngId",
		ExpiryDate:   "mockDate",
		AmountExport: 10,
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
			name: "IngredientDetail update successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						mockDetail.IngredientId,
						mockDetail.ExpiryDate,
						mock.Anything,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "IngredientDetail failed and return error",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						mockDetail.IngredientId,
						mockDetail.ExpiryDate,
						mock.Anything,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createExportNoteRepo{
				exportNoteStore:       tt.fields.exportNoteStore,
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.updateIngredientDetail(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "updateIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "updateIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
