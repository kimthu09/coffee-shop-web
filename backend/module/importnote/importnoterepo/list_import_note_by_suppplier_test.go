package importnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockImportNoteBySupplierStore struct {
	mock.Mock
}

func (m *mockImportNoteBySupplierStore) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging,
	moreKeys ...string) ([]importnotemodel.ImportNote, error) {
	args := m.Called(ctx, supplierId, filter, paging, moreKeys)
	return args.Get(0).([]importnotemodel.ImportNote),
		args.Error(1)
}

func TestNewListImportNoteBySupplierRepo(t *testing.T) {
	type args struct {
		importNoteStore ListImportNoteBySupplierStore
	}

	mockStore := new(mockImportNoteBySupplierStore)

	tests := []struct {
		name string
		args args
		want *listImportNoteBySupplierRepo
	}{
		{
			name: "Create object has type ListIngredientDetailBiz",
			args: args{
				importNoteStore: mockStore,
			},
			want: &listImportNoteBySupplierRepo{
				importNoteStore: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteBySupplierRepo(tt.args.importNoteStore)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListImportNoteBySupplierRepo() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listImportNoteBySupplierRepo_ListImportNoteBySupplier(t *testing.T) {
	type fields struct {
		importNoteStore ListImportNoteBySupplierStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
		filter     *filter.SupplierImportFilter
		paging     *common.Paging
	}

	mockStore := new(mockImportNoteBySupplierStore)

	paging := common.Paging{
		Page: 1,
	}
	filterImportNote := filter.SupplierImportFilter{
		DateFrom: nil,
		DateTo:   nil,
	}
	supplierId := mock.Anything
	listImportNotes := make([]importnotemodel.ImportNote, 0)
	var emptyListImportNotes []importnotemodel.ImportNote
	moreKeys := []string{"Supplier"}
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
			name: "List import note failed because can not get data from database",
			fields: fields{
				importNoteStore: mockStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				mockStore.
					On(
						"ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImportNote,
						&paging,
						moreKeys,
					).
					Return(emptyListImportNotes, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List supplier debt successfully",
			fields: fields{
				importNoteStore: mockStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				mockStore.
					On(
						"ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImportNote,
						&paging,
						moreKeys,
					).
					Return(listImportNotes, nil).
					Once()
			},
			want:    listImportNotes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listImportNoteBySupplierRepo{
				importNoteStore: tt.fields.importNoteStore,
			}
			tt.mock()

			got, err := biz.ListImportNoteBySupplier(
				tt.args.ctx, tt.args.supplierId, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListImportNoteBySupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListImportNoteBySupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListImportNoteBySupplier() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
