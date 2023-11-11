package importnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListImportNoteStore struct {
	mock.Mock
}

func (m *mockListImportNoteStore) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotemodel.ImportNote), args.Error(1)
}

func TestNewListImportNoteRepo(t *testing.T) {
	type args struct {
		store ListImportNoteStore
	}

	mockStore := new(mockListImportNoteStore)

	tests := []struct {
		name string
		args args
		want *listImportNoteRepo
	}{
		{
			name: "Create object has type ListImportNoteRepo",
			args: args{
				store: mockStore,
			},
			want: &listImportNoteRepo{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListImportNoteRepo(%v)", tt.args.store)
		})
	}
}

func Test_listImportNoteRepo_ListImportNote(t *testing.T) {
	type fields struct {
		store ListImportNoteStore
	}
	type args struct {
		ctx    context.Context
		filter *importnotemodel.Filter
		paging *common.Paging
	}

	mockStore := new(mockListImportNoteStore)
	filter := importnotemodel.Filter{
		SearchKey: "",
		MinPrice:  nil,
		MaxPrice:  nil,
		Status:    "",
	}
	paging := common.Paging{
		Page: 1,
	}
	importNotes := []importnotemodel.ImportNote{
		{
			Id:         mock.Anything,
			TotalPrice: 0,
			CreateAt:   nil,
			CreateBy:   mock.Anything,
		},
		{
			Id:         mock.Anything,
			TotalPrice: 0,
			CreateAt:   nil,
			CreateBy:   mock.Anything,
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "List import note successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockStore.
					On(
						"ListImportNote",
						context.Background(),
						&filter,
						mock.Anything,
						&paging).
					Return(importNotes, nil).
					Once()
			},
			want:    importNotes,
			wantErr: false,
		},
		{
			name: "List import note successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockStore.
					On(
						"ListImportNote",
						context.Background(),
						&filter,
						mock.Anything,
						&paging).
					Return(nil, mockErr).
					Once()
			},
			want:    importNotes,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listImportNoteRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListImportNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"ListImportNote() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListIngredient() got = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
