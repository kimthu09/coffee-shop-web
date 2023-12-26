package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeSupplierImportNoteRepo struct {
	mock.Mock
}

func (m *mockSeeSupplierImportNoteRepo) SeeSupplierImportNote(
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

func TestNewSeeSupplierImportNoteBiz(t *testing.T) {
	type args struct {
		repo      SeeSupplierImportNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierImportNoteRepo)

	tests := []struct {
		name string
		args args
		want *seeSupplierImportNoteBiz
	}{
		{
			name: "Create object has type SeeSupplierImportNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeSupplierImportNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierImportNoteBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeSupplierImportNoteBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeSupplierImportNoteBiz_SeeSupplierImportNote(t *testing.T) {
	type fields struct {
		repo      SeeSupplierImportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		supplierId string
		filter     *filter.SupplierImportFilter
		paging     *common.Paging
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeSupplierImportNoteRepo)
	supplierId := mock.Anything
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

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "See supplier import note failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
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
			name: "See supplier import note failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierImportNote",
						context.Background(),
						supplierId,
						&filterImportNote,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See supplier import note successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     &filterImportNote,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeSupplierImportNote",
						context.Background(),
						supplierId,
						&filterImportNote,
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
			biz := &seeSupplierImportNoteBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeSupplierImportNote(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filter,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeSupplierImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeSupplierImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeSupplierImportNote() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
