package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeImportNoteDetailRepo struct {
	mock.Mock
}

func (m *mockSeeImportNoteDetailRepo) SeeImportNoteDetail(
	ctx context.Context,
	supplierId string) (*importnotemodel.ImportNote, error) {
	args := m.Called(ctx, supplierId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*importnotemodel.ImportNote), args.Error(1)
}

func TestNewSeeImportNoteDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeImportNoteDetailRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeImportNoteDetailRepo)

	tests := []struct {
		name string
		args args
		want *seeImportNoteDetailBiz
	}{
		{
			name: "Create object has type SeeImportNoteDetailBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeImportNoteDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeImportNoteDetailBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeImportNoteDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeImportNoteDetailBiz_SeeImportNoteDetail(t *testing.T) {
	type fields struct {
		repo      SeeImportNoteDetailRepo
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeImportNoteDetailRepo)

	importNoteId := mock.Anything
	status := importnotemodel.Done
	validId := mock.Anything
	importNote := importnotemodel.ImportNote{
		Id: importNoteId,
		Supplier: suppliermodel.SimpleSupplier{
			Id:   mock.Anything,
			Name: mock.Anything,
		},
		TotalPrice: 100,
		Status:     &status,
		CreatedBy:  mock.Anything,
		ClosedBy:   &validId,
		CreatedAt:  nil,
		ClosedAt:   nil,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "See import note detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &importNote,
			wantErr: true,
		},
		{
			name: "See import note detail failed because can not get supplier from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeImportNoteDetail",
						context.Background(),
						importNoteId).
					Return(nil, mockErr).
					Once()
			},
			want:    &importNote,
			wantErr: true,
		},
		{
			name: "See import note detail successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeImportNoteDetail",
						context.Background(),
						importNoteId).
					Return(&importNote, nil).
					Once()
			},
			want:    &importNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeImportNoteDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeImportNoteDetail(tt.args.ctx, tt.args.importNoteId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeImportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeImportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeImportNoteDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
