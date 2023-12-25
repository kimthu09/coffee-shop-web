package exportnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListExportNoteStore struct {
	mock.Mock
}

func (m *mockListExportNoteStore) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string,
) ([]exportnotemodel.ExportNote, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]exportnotemodel.ExportNote), args.Error(1)
}

func TestNewListExportNoteRepo(t *testing.T) {
	type args struct {
		store ListExportNoteStore
	}

	mockExportNote := new(mockListExportNoteStore)

	tests := []struct {
		name string
		args args
		want *listExportNoteRepo
	}{
		{
			name: "Create object has type ListExportNoteRepo",
			args: args{
				store: mockExportNote,
			},
			want: &listExportNoteRepo{
				store: mockExportNote,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListExportNoteRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListExportNoteRepo(%v)", tt.args.store)
		})
	}
}

func Test_listExportNoteRepo_ListExportNote(t *testing.T) {
	type fields struct {
		store ListExportNoteStore
	}
	type args struct {
		ctx    context.Context
		filter *exportnotemodel.Filter
		paging *common.Paging
	}

	mockExportNote := new(mockListExportNoteStore)
	mockFilter := exportnotemodel.Filter{
		SearchKey: "",
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockExportNotes := []exportnotemodel.ExportNote{
		{
			Id:        mock.Anything,
			Reason:    nil,
			CreatedAt: nil,
			CreatedBy: mock.Anything,
		},
		{
			Id:        mock.Anything,
			Reason:    nil,
			CreatedAt: nil,
			CreatedBy: mock.Anything,
		},
	}
	moreKeys := []string{"CreatedByUser"}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []exportnotemodel.ExportNote
		wantErr bool
	}{
		{
			name: "List export note successfully",
			fields: fields{
				store: mockExportNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockExportNote.
					On(
						"ListExportNote",
						context.Background(),
						&mockFilter,
						mock.Anything,
						&mockPaging,
						moreKeys).
					Return(mockExportNotes, nil).
					Once()
			},
			want:    mockExportNotes,
			wantErr: false,
		},
		{
			name: "List export note successfully",
			fields: fields{
				store: mockExportNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockExportNote.
					On(
						"ListExportNote",
						context.Background(),
						&mockFilter,
						mock.Anything,
						&mockPaging,
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    mockExportNotes,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listExportNoteRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListExportNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListExportNote() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"ListExportNote() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListExportNote() got = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
