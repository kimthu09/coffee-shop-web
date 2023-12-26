package importnoterepo

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockFindImportNoteStore struct {
	mock.Mock
}

func (m *mockFindImportNoteStore) FindImportNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*importnotemodel.ImportNote, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*importnotemodel.ImportNote), args.Error(1)
}

type mockSeeDetailImportNoteStore struct {
	mock.Mock
}

func (m *mockSeeDetailImportNoteStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	args := m.Called(ctx, importNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotedetailmodel.ImportNoteDetail), args.Error(1)
}

func TestNewSeeImportNoteDetailRepo(t *testing.T) {
	type args struct {
		importNoteStore       FindImportNoteStore
		importNoteDetailStore SeeDetailImportNoteStore
	}

	importNoteStore := new(mockFindImportNoteStore)
	importNoteDetailStore := new(mockSeeDetailImportNoteStore)

	tests := []struct {
		name string
		args args
		want *seeImportNoteDetailRepo
	}{
		{
			name: "Create object has type NewSeeImportNoteDetailRepo",
			args: args{
				importNoteStore:       importNoteStore,
				importNoteDetailStore: importNoteDetailStore,
			},
			want: &seeImportNoteDetailRepo{
				importNoteStore:       importNoteStore,
				importNoteDetailStore: importNoteDetailStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeImportNoteDetailRepo(
				tt.args.importNoteStore,
				tt.args.importNoteDetailStore,
			)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeImportNoteDetailRepo() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeImportNoteDetailRepo_SeeImportNoteDetail(t *testing.T) {
	type fields struct {
		importNoteStore       FindImportNoteStore
		importNoteDetailStore SeeDetailImportNoteStore
	}
	type args struct {
		ctx          context.Context
		importNoteId string
	}

	importNoteStore := new(mockFindImportNoteStore)
	importNoteDetailStore := new(mockSeeDetailImportNoteStore)

	importNoteId := mock.Anything
	importNote := importnotemodel.ImportNote{
		Id: importNoteId,
		Supplier: suppliermodel.SimpleSupplier{
			Id:   mock.Anything,
			Name: mock.Anything,
		},
	}
	importNoteDetail := []importnotedetailmodel.ImportNoteDetail{
		{
			ImportNoteId: importNoteId,
			IngredientId: mock.Anything,
		},
	}

	mockErr := errors.New("failed to see import note detail")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "See import note detail failed because can not get import note from database",
			fields: fields{
				importNoteStore:       importNoteStore,
				importNoteDetailStore: importNoteDetailStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				importNoteStore.
					On("FindImportNote",
						context.Background(),
						map[string]interface{}{"id": importNoteId},
						[]string{"Supplier", "CreatedByUser", "ClosedByUser"}).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See import note detail failed because can not get import note details from database",
			fields: fields{
				importNoteStore:       importNoteStore,
				importNoteDetailStore: importNoteDetailStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				importNoteStore.
					On("FindImportNote",
						context.Background(),
						map[string]interface{}{"id": importNoteId},
						[]string{"Supplier", "CreatedByUser", "ClosedByUser"}).
					Return(&importNote, nil).
					Once()

				importNoteDetailStore.
					On("ListImportNoteDetail",
						context.Background(),
						importNoteId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See import note detail successfully",
			fields: fields{
				importNoteStore:       importNoteStore,
				importNoteDetailStore: importNoteDetailStore,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
			},
			mock: func() {
				importNoteStore.
					On("FindImportNote",
						context.Background(),
						map[string]interface{}{"id": importNoteId},
						[]string{"Supplier", "CreatedByUser", "ClosedByUser"}).
					Return(&importNote, nil).
					Once()

				importNoteDetailStore.
					On("ListImportNoteDetail",
						context.Background(),
						importNoteId).
					Return(importNoteDetail, nil).
					Once()
			},
			want:    &importNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeImportNoteDetailRepo{
				importNoteStore:       tt.fields.importNoteStore,
				importNoteDetailStore: tt.fields.importNoteDetailStore,
			}

			tt.mock()

			got, err := repo.SeeImportNoteDetail(
				tt.args.ctx, tt.args.importNoteId)

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
