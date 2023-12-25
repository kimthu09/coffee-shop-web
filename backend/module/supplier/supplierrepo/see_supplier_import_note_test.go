package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeSupplierImportNoteStore struct {
	mock.Mock
}

func (m *mockSeeSupplierImportNoteStore) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging,
	moreKeys ...string) ([]importnotemodel.ImportNote, error) {
	args := m.Called(ctx, supplierId, filter, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotemodel.ImportNote), args.Error(1)
}

func TestNewSeeSupplierImportNoteRepo(t *testing.T) {
	type args struct {
		importNoteStore ListSupplierImportNoteStore
	}

	mockStore := new(mockSeeSupplierImportNoteStore)

	tests := []struct {
		name string
		args args
		want *seeSupplierImportNoteRepo
	}{
		{
			name: "Create object has type SeeSupplierImportNoteRepo",
			args: args{
				importNoteStore: mockStore,
			},
			want: &seeSupplierImportNoteRepo{
				importNoteStore: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierImportNoteRepo(tt.args.importNoteStore)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeSupplierImportNoteRepo() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeSupplierImportNoteRepo_SeeSupplierImportNote(t *testing.T) {
	type fields struct {
		importNoteStore ListSupplierImportNoteStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
		filter     *filter.SupplierImportFilter
		paging     *common.Paging
	}

	store := new(mockSeeSupplierImportNoteStore)

	supplierId := "123"
	date := int64(123)
	filterImportNote := filter.SupplierImportFilter{
		DateFrom: &date,
		DateTo:   &date,
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
		Total: 12,
	}
	status := importnotemodel.Done
	closedBy := mock.Anything
	importNotes := []importnotemodel.ImportNote{
		{
			Id:         mock.Anything,
			SupplierId: supplierId,
			Supplier: suppliermodel.SimpleSupplier{
				Id:   supplierId,
				Name: mock.Anything,
			},
			TotalPrice: 0,
			Status:     &status,
			CreatedBy:  mock.Anything,
			ClosedBy:   &closedBy,
		},
	}
	mockErr := errors.New(mock.Anything)
	moreKeys := []string{"CreatedByUser", "ClosedByUser"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "See supplier import note failed because can not get data from database",
			fields: fields{
				importNoteStore: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				store.
					On("ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImportNote,
						&paging,
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    importNotes,
			wantErr: true,
		},
		{
			name: "See supplier import note successfully",
			fields: fields{
				importNoteStore: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				store.
					On("ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImportNote,
						&paging,
						moreKeys,
					).
					Return(importNotes, nil).
					Once()
			},
			want:    importNotes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeSupplierImportNoteRepo{
				importNoteStore: tt.fields.importNoteStore,
			}

			tt.mock()

			got, err := biz.SeeSupplierImportNote(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filter,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeSupplierImportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeSupplierImportNote() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeSupplierImportNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
