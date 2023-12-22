package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockSeeExportNoteDetailStore struct {
	mock.Mock
}

func (m *mockSeeExportNoteDetailStore) ListExportNoteDetail(
	ctx context.Context,
	exportNoteId string) ([]exportnotedetailmodel.ExportNoteDetail, error) {
	args := m.Called(ctx, exportNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]exportnotedetailmodel.ExportNoteDetail), args.Error(1)
}

type mockFindExportNoteStore struct {
	mock.Mock
}

func (m *mockFindExportNoteStore) FindExportNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*exportnotemodel.ExportNote, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*exportnotemodel.ExportNote), args.Error(1)
}

func TestNewSeeExportNoteDetailRepo(t *testing.T) {
	type args struct {
		exportNoteDetailStore SeeExportNoteDetailStore
		exportNoteStore       FindExportNoteStore
	}

	mockExportNoteDetail := new(mockSeeExportNoteDetailStore)
	mockExportNote := new(mockFindExportNoteStore)

	tests := []struct {
		name string
		args args
		want *seeExportNoteDetailRepo
	}{
		{
			name: "Create object has type ListExportNoteRepo",
			args: args{
				exportNoteDetailStore: mockExportNoteDetail,
				exportNoteStore:       mockExportNote,
			},
			want: &seeExportNoteDetailRepo{
				exportNoteDetailStore: mockExportNoteDetail,
				exportNoteStore:       mockExportNote,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeExportNoteDetailRepo(
				tt.args.exportNoteDetailStore,
				tt.args.exportNoteStore)
			assert.Equal(
				t, tt.want, got,
				"NewListExportNoteRepo(%v, %v)",
				tt.args.exportNoteDetailStore, tt.args.exportNoteStore)
		})
	}
}

func Test_seeExportNoteDetailRepo_SeeExportNoteDetail(t *testing.T) {
	type fields struct {
		exportNoteDetailStore SeeExportNoteDetailStore
		exportNoteStore       FindExportNoteStore
	}
	type args struct {
		ctx          context.Context
		exportNoteId string
	}

	mockExportNoteDetail := new(mockSeeExportNoteDetailStore)
	mockExportNote := new(mockFindExportNoteStore)

	exportNoteId := "exportNote1"
	reason := exportnotemodel.Damaged
	createdBy := "user001"
	createdAt := time.Date(
		2023, 12, 20,
		0, 0, 0, 0, time.UTC)
	exportNote := exportnotemodel.ExportNote{
		Id:        exportNoteId,
		Reason:    &reason,
		CreatedAt: &createdAt,
		CreatedBy: createdBy,
		Details:   nil,
	}
	ingredient1 := "Ingredient1"
	details := []exportnotedetailmodel.ExportNoteDetail{
		{
			ExportNoteId: exportNoteId,
			IngredientId: ingredient1,
			AmountExport: 100,
		},
	}
	finalExportNote := exportnotemodel.ExportNote{
		Id:        exportNoteId,
		Reason:    &reason,
		CreatedAt: &createdAt,
		CreatedBy: createdBy,
		Details:   details,
	}
	moreKeys := []string{"CreatedByUser"}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *exportnotemodel.ExportNote
		wantErr bool
	}{
		{
			name: "see export note detail failed because can not find export note",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockExportNote.
					On(
						"FindExportNote",
						context.Background(),
						map[string]interface{}{"id": exportNoteId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    &finalExportNote,
			wantErr: true,
		},
		{
			name: "see export note detail failed because can not get export note detail",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockExportNote.
					On(
						"FindExportNote",
						context.Background(),
						map[string]interface{}{"id": exportNoteId},
						moreKeys,
					).
					Return(&exportNote, nil).
					Once()

				mockExportNoteDetail.
					On(
						"ListExportNoteDetail",
						context.Background(),
						exportNoteId,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    &finalExportNote,
			wantErr: true,
		},
		{
			name: "see export note detail successfully",
			fields: fields{
				exportNoteStore:       mockExportNote,
				exportNoteDetailStore: mockExportNoteDetail,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockExportNote.
					On(
						"FindExportNote",
						context.Background(),
						map[string]interface{}{"id": exportNoteId},
						moreKeys,
					).
					Return(&exportNote, nil).
					Once()

				mockExportNoteDetail.
					On(
						"ListExportNoteDetail",
						context.Background(),
						exportNoteId,
					).
					Return(details, nil).
					Once()
			},
			want:    &finalExportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeExportNoteDetailRepo{
				exportNoteDetailStore: tt.fields.exportNoteDetailStore,
				exportNoteStore:       tt.fields.exportNoteStore,
			}

			tt.mock()

			got, err := biz.SeeExportNoteDetail(tt.args.ctx, tt.args.exportNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeExportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeExportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "SeeExportNoteDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
