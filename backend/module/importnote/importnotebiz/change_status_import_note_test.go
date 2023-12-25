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
	"time"
)

type mockChangeStatusImportNote struct {
	mock.Mock
}

func (m *mockChangeStatusImportNote) UpdateDebtSupplier(
	ctx context.Context,
	importNote *importnotemodel.ImportNoteUpdate) error {
	args := m.Called(ctx, importNote)
	return args.Error(0)
}
func (m *mockChangeStatusImportNote) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ImportNoteUpdate) error {
	args := m.Called(ctx, supplierDebtId, importNote)
	return args.Error(0)
}
func (m *mockChangeStatusImportNote) FindImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	args := m.Called(ctx, importNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*importnotemodel.ImportNote), args.Error(1)
}
func (m *mockChangeStatusImportNote) UpdateImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	args := m.Called(ctx, importNoteId, data)
	return args.Error(0)
}
func (m *mockChangeStatusImportNote) FindListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	args := m.Called(ctx, importNoteId)
	return args.Get(0).([]importnotedetailmodel.ImportNoteDetail),
		args.Error(1)
}
func (m *mockChangeStatusImportNote) HandleIngredientDetails(
	ctx context.Context,
	importNoteDetails []importnotedetailmodel.ImportNoteDetail) error {
	args := m.Called(ctx, importNoteDetails)
	return args.Error(0)
}
func (m *mockChangeStatusImportNote) HandleIngredient(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]int) error {
	args := m.Called(ctx, ingredientTotalAmountNeedUpdate)
	return args.Error(0)
}
func TestNewChangeStatusImportNoteBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      ChangeStatusImportNoteRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockChangeStatusImportNote)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *changeStatusImportNoteRepo
	}{
		{
			name: "Create object has type ChangeStatusImportNoteRepo",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &changeStatusImportNoteRepo{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusImportNoteBiz(
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewChangeStatusImportNoteBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusImportNoteRepo_ChangeStatusImportNote(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      ChangeStatusImportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		importNoteId string
		data         *importnotemodel.ImportNoteUpdate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockChangeStatusImportNote)
	mockRequest := new(mockRequester)

	validId := "012345678901"
	doneStatus := importnotemodel.Done
	inProgressStatus := importnotemodel.InProgress
	invalidImportNoteUpdateStatus := importnotemodel.ImportNoteUpdate{
		Status: &inProgressStatus,
	}
	importNoteUpdateStatus := importnotemodel.ImportNoteUpdate{
		Status: &doneStatus,
	}
	now := time.Now()
	closedImportNote := importnotemodel.ImportNote{
		Id:         validId,
		SupplierId: validId,
		TotalPrice: 1000,
		Status:     &doneStatus,
		CreatedBy:  validId,
		CreatedAt:  &now,
	}
	inProgressImportNote := importnotemodel.ImportNote{
		Id:         validId,
		SupplierId: validId,
		TotalPrice: 1000,
		Status:     &inProgressStatus,
		CreatedBy:  validId,
		CreatedAt:  &now,
	}
	emptyImportNoteDetail := []importnotedetailmodel.ImportNoteDetail{}
	importNoteDetails := []importnotedetailmodel.ImportNoteDetail{
		{
			ImportNoteId: validId,
			IngredientId: validId,
			Price:        10000,
			AmountImport: 100,
		},
	}
	mapIngredientDetail := map[string]int{"012345678901": 100}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Change status import note failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because data is not valid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidImportNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not get current import note",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because import note already has been closed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&closedImportNote, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not " +
				"create supplier debt when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not " +
				"update debt of supplier when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						&importNoteUpdateStatus).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not " +
				"get list of ingredient when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						validId).
					Return(emptyImportNoteDetail, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not " +
				"handle total amount of ingredient when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						validId).
					Return(importNoteDetails, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredient",
						context.Background(),
						mapIngredientDetail).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note failed because can not " +
				"update of import note when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						validId).
					Return(importNoteDetails, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredient",
						context.Background(),
						mapIngredientDetail).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateImportNote",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status import note successfully when status need to change is Done",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: validId,
				data:         &importNoteUpdateStatus,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteChangeStatusFeatureCode).
					Return(true).
					Once()

				mockRequest.
					On("GetUserId").
					Return(mock.Anything).
					Once()

				mockRepo.
					On(
						"FindImportNote",
						context.Background(),
						validId).
					Return(&inProgressImportNote, nil).
					Once()

				mockRepo.
					On(
						"CreateSupplierDebt",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateDebtSupplier",
						context.Background(),
						&importNoteUpdateStatus).
					Return(nil).
					Once()

				mockRepo.
					On(
						"FindListImportNoteDetail",
						context.Background(),
						validId).
					Return(importNoteDetails, nil).
					Once()

				mockRepo.
					On(
						"HandleIngredient",
						context.Background(),
						mapIngredientDetail).
					Return(nil).
					Once()

				mockRepo.
					On(
						"UpdateImportNote",
						context.Background(),
						validId,
						&importNoteUpdateStatus).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &changeStatusImportNoteRepo{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.ChangeStatusImportNote(tt.args.ctx, tt.args.importNoteId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ChangeStatusImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ChangeStatusImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			}
		})
	}
}
