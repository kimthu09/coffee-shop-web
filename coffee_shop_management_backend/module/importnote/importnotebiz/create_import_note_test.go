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

	importNoteId := "importNote1"
	ingredient1 := "ingredient1"
	ingredient2 := "ingredient2"
	price1 := float32(0.0015)
	price2 := float32(0.2)
	roundedPrice1 := float32(0.002)
	amountImport := 10
	validDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{
			IngredientId:   ingredient1,
			Price:          price1,
			IsReplacePrice: true,
			AmountImport:   amountImport,
		},
		{
			IngredientId:   ingredient2,
			Price:          price2,
			IsReplacePrice: false,
			AmountImport:   amountImport,
		},
	}
	finalDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{
			ImportNoteId:   importNoteId,
			IngredientId:   ingredient1,
			Price:          roundedPrice1,
			IsReplacePrice: true,
			AmountImport:   amountImport,
			TotalUnit:      float32(0.02),
		},
		{
			ImportNoteId:   importNoteId,
			IngredientId:   ingredient2,
			Price:          price2,
			IsReplacePrice: false,
			AmountImport:   amountImport,
			TotalUnit:      float32(2),
		},
	}
	invalidDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{
			IngredientId:   ingredient1,
			Price:          price1,
			IsReplacePrice: true,
			AmountImport:   amountImport,
		},
		{
			IngredientId:   ingredient1,
			Price:          price1,
			IsReplacePrice: true,
			AmountImport:   amountImport,
		},
	}
	supplierId := "supplier1"
	createdBy := "user1"
	importNoteCreate := importnotemodel.ImportNoteCreate{
		SupplierId:        supplierId,
		CreatedBy:         createdBy,
		ImportNoteDetails: validDetails,
	}
	finalImportNoteCreate := importnotemodel.ImportNoteCreate{
		Id:                &importNoteId,
		TotalPrice:        2,
		SupplierId:        supplierId,
		CreatedBy:         createdBy,
		ImportNoteDetails: finalDetails,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name       string
		fields     fields
		args       args
		mock       func()
		finalParam *importnotemodel.ImportNoteCreate
		wantErr    bool
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
			finalParam: nil,
			wantErr:    true,
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
					Id:                &importNoteId,
					TotalPrice:        1000,
					SupplierId:        mock.Anything,
					CreatedBy:         mock.Anything,
					ImportNoteDetails: invalidDetails,
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteCreateFeatureCode).
					Return(true).
					Once()
			},
			finalParam: nil,
			wantErr:    true,
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

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(nil, mockErr).
					Once()
			},
			finalParam: nil,
			wantErr:    true,
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

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&importNoteId, nil).
					Once()

				mockRepo.
					On(
						"HandleCreateImportNote",
						context.Background(),
						&importNoteCreate).
					Return(mockErr).
					Once()
			},
			finalParam: nil,
			wantErr:    true,
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

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&importNoteId, nil).
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
			finalParam: nil,
			wantErr:    true,
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

				mockGenerator.
					On("IdProcess", importNoteCreate.Id).
					Return(&importNoteId, nil).
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
			finalParam: &finalImportNoteCreate,
			wantErr:    false,
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
				if tt.finalParam != nil {
					assert.Equal(
						t,
						tt.args.data,
						tt.finalParam,
						"data = %v, want %v",
						tt.args.data,
						tt.finalParam)
				}
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
