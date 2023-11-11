package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
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

type mockCreateImportNoteRepo struct {
	mock.Mock
}

func (m *mockCreateImportNoteRepo) CheckIngredient(
	ctx context.Context,
	ingredientId string) error {
	args := m.Called(ctx, ingredientId)
	return args.Error(0)
}
func (m *mockCreateImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
func (m *mockCreateImportNoteRepo) UpdatePriceIngredient(
	ctx context.Context,
	ingredientId string,
	price float32) error {
	args := m.Called(ctx, ingredientId, price)
	return args.Error(0)
}
func (m *mockCreateImportNoteRepo) CheckSupplier(
	ctx context.Context,
	supplierId string) error {
	args := m.Called(ctx, supplierId)
	return args.Error(0)
}

func TestNewCreateImportNoteBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateImportNoteRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateImportNoteRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createImportNoteBiz
	}{
		{
			name: "Create object has type CreateImportNoteBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createImportNoteBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateImportNoteBiz(
				tt.args.gen,
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateImportNoteBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createImportNoteBiz_CreateImportNote(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateImportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *importnotemodel.ImportNoteCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateImportNoteRepo)
	mockRequest := new(mockRequester)

	validId := "012345678901"
	validDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{
			IngredientId:   validId,
			ExpiryDate:     "11/11/2023",
			Price:          10000,
			IsReplacePrice: true,
			AmountImport:   100,
		},
		{
			IngredientId:   validId,
			ExpiryDate:     "12/11/2023",
			Price:          1000,
			IsReplacePrice: false,
			AmountImport:   100,
		},
	}
	invalidDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{
			IngredientId:   validId,
			ExpiryDate:     "11/11/2023",
			Price:          10000,
			IsReplacePrice: true,
			AmountImport:   100,
		},
		{
			IngredientId:   validId,
			ExpiryDate:     "12/11/2023",
			Price:          1000,
			IsReplacePrice: true,
			AmountImport:   100,
		},
	}
	importNoteCreate := importnotemodel.ImportNoteCreate{
		TotalPrice:        200000,
		SupplierId:        validId,
		CreateBy:          validId,
		ImportNoteDetails: validDetails,
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
			name: "Create import note failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &importnotemodel.ImportNoteCreate{
					Id:                &validId,
					TotalPrice:        1000,
					SupplierId:        mock.Anything,
					CreateBy:          mock.Anything,
					ImportNoteDetails: invalidDetails,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not check ingredient",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not handle id of import note",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(nil).
					Times(2)

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not check supplier",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(nil).
					Times(2)

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CheckSupplier",
						context.Background(),
						importNoteCreate.SupplierId).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not check supplier",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(nil).
					Times(2)

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CheckSupplier",
						context.Background(),
						importNoteCreate.SupplierId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleCreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note failed because can not update price of ingredient",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(nil).
					Times(2)

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CheckSupplier",
						context.Background(),
						importNoteCreate.SupplierId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleCreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdatePriceIngredient",
						context.Background(),
						importNoteCreate.ImportNoteDetails[0].IngredientId,
						importNoteCreate.ImportNoteDetails[0].Price).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create import note successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On("CheckIngredient", context.Background(), mock.Anything).
					Return(nil).
					Times(2)

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&validId, nil).
					Once()

				mockRepo.
					On(
						"CheckSupplier",
						context.Background(),
						importNoteCreate.SupplierId).
					Return(nil).
					Once()

				mockRepo.
					On(
						"HandleCreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdatePriceIngredient",
						context.Background(),
						importNoteCreate.ImportNoteDetails[0].IngredientId,
						importNoteCreate.ImportNoteDetails[0].Price).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createImportNoteBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateImportNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"CreateImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"CreateImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			}
		})
	}
}
