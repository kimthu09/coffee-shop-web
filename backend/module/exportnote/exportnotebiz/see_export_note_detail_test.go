package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeExportNoteDetailRepo struct {
	mock.Mock
}

func (m *mockSeeExportNoteDetailRepo) SeeExportNoteDetail(
	ctx context.Context,
	exportNoteId string) (*exportnotemodel.ExportNote, error) {
	args := m.Called(ctx, exportNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*exportnotemodel.ExportNote), args.Error(1)
}

func TestNewSeeExportNoteDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeExportNoteDetailRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeExportNoteDetailRepo)

	tests := []struct {
		name string
		args args
		want *seeExportNoteDetailBiz
	}{
		{
			name: "Create object has type SeeExportNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeExportNoteDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeExportNoteDetailBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeExportNoteDetailBiz(%v, %v) = %v, want %v",
				tt.args.repo,
				tt.args.requester,
				got,
				tt.want)
		})
	}
}

func Test_seeExportNoteDetailBiz_SeeExportNoteDetail(t *testing.T) {
	type fields struct {
		repo      SeeExportNoteDetailRepo
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		exportNoteId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeExportNoteDetailRepo)

	exportNoteId := "exportNote1"
	exportNote := exportnotemodel.ExportNote{
		Id:        exportNoteId,
		Reason:    nil,
		CreatedAt: nil,
		CreatedBy: "",
		Details:   nil,
	}
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
			name: "See export note failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &exportNote,
			wantErr: true,
		},
		{
			name: "See export note failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeExportNoteDetail",
						context.Background(),
						exportNoteId,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    &exportNote,
			wantErr: true,
		},
		{
			name: "See export note successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ExportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeExportNoteDetail",
						context.Background(),
						exportNoteId,
					).
					Return(&exportNote, nil).
					Once()
			},
			want:    &exportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeExportNoteDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeExportNoteDetail(tt.args.ctx, tt.args.exportNoteId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeExportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeExportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeExportNoteDetail() want = %v, got %v",
					tt.want,
					got)
			}

		})
	}
}
