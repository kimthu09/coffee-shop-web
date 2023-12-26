package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListImportNotBySupplierRepo struct {
	mock.Mock
}

func (m *mockListImportNotBySupplierRepo) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	args := m.Called(ctx, supplierId, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotemodel.ImportNote), args.Error(1)
}

func TestNewListImportNoteBySupplierBiz(t *testing.T) {
	type args struct {
		repo      ListImportNoteBySupplierRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListImportNotBySupplierRepo)

	tests := []struct {
		name string
		args args
		want *listImportNoteBySupplierBiz
	}{
		{
			name: "Create object has type ListImportNoteBySupplierBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listImportNoteBySupplierBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteBySupplierBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListImportNoteBySupplierBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listImportNoteBySupplierBiz_ListImportNoteBySupplier(t *testing.T) {
	type fields struct {
		repo      ListImportNoteBySupplierRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		supplierId string
		filter     *filter.SupplierImportFilter
		paging     *common.Paging
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListImportNotBySupplierRepo)

	supplierId := mock.Anything
	filterImport := filter.SupplierImportFilter{
		DateFrom: nil,
		DateTo:   nil,
	}
	paging := common.Paging{
		Page:  2,
		Limit: 10,
	}
	importNotes := make([]importnotemodel.ImportNote, 0)
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
			name: "List import note by supplier failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImport,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List import note by supplier failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImport,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImport,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List import note successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImport,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListImportNoteBySupplier",
						context.Background(),
						supplierId,
						&filterImport,
						&paging,
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
			biz := &listImportNoteBySupplierBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListImportNoteBySupplier(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filter,
				tt.args.paging)

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
					"ListImportNoteBySupplier() want = %v, got %v",
					tt.want,
					got)
			}

		})
	}
}
