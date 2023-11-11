package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
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

type mockCreateExportNoteRepo struct {
	mock.Mock
}

func (m *mockCreateExportNoteRepo) GetPriceIngredient(
	ctx context.Context,
	ingredientId string) (*float32, error) {
	args := m.Called(ctx, ingredientId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}

func (m *mockCreateExportNoteRepo) HandleExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateExportNoteRepo) HandleIngredientDetail(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateExportNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]float32,
) error {
	args := m.Called(ctx, ingredientTotalAmountNeedUpdate)
	return args.Error(0)
}

func TestNewCreateExportNoteBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateExportNoteRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateExportNoteRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createExportNoteBiz
	}{
		{
			name: "Create object has type CreateExportNoteBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createExportNoteBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateExportNoteBiz(
				tt.args.gen,
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateExportNoteBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createExportNoteBiz_CreateExportNote(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateExportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *exportnotemodel.ExportNoteCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateExportNoteRepo)
	mockRequest := new(mockRequester)

	invalidExportNoteCreate := exportnotemodel.ExportNoteCreate{
		Id: nil,
	}
	mockExportNoteCreate := exportnotemodel.ExportNoteCreate{
		Id: nil,
		ExportNoteDetails: []exportnotedetailmodel.ExportNoteDetailCreate{
			{
				IngredientId: "Ing001",
				ExpiryDate:   "08/11/2023",
				AmountExport: 20,
			},
			{
				IngredientId: "Ing001",
				ExpiryDate:   "09/11/2023",
				AmountExport: 10,
			},
		},
	}
	mockId := "123456789"
	mockErr := errors.New(mock.Anything)
	mockPrice := float32(10000)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create export note failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because can not generate Id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because can not get price ingredient",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because can not handle export note successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockExportNoteCreate.ExportNoteDetails))

				mockRepo.
					On("HandleExportNote", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because can not handle ingredient detail successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockExportNoteCreate.ExportNoteDetails))

				mockRepo.
					On("HandleExportNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note failed because can not update ingredient total amount successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockExportNoteCreate.ExportNoteDetails))

				mockRepo.
					On("HandleExportNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientTotalAmount", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create export note successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockExportNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockExportNoteCreate.ExportNoteDetails))

				mockRepo.
					On("HandleExportNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientTotalAmount", context.Background(), mock.Anything).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createExportNoteBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateExportNote(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "CreateExportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateExportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
