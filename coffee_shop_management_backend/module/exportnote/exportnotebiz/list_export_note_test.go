package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListExportNoteRepo struct {
	mock.Mock
}

func (m *mockListExportNoteRepo) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	paging *common.Paging) ([]exportnotemodel.ExportNote, error) {
	args := m.Called(ctx, filter, paging)
	return args.Get(0).([]exportnotemodel.ExportNote), args.Error(1)
}

type mockRequester struct {
	mock.Mock
}

func (m *mockRequester) GetUserId() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetEmail() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetRole() rolemodel.Role {
	args := m.Called()
	return args.Get(0).(rolemodel.Role)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called()
	return args.Bool(0)
}

func TestNewListExportNoteRepo(t *testing.T) {
	type args struct {
		repo      ListExportNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListExportNoteRepo)

	tests := []struct {
		name string
		args args
		want *listExportNoteBiz
	}{
		{
			name: "Create object has type ListExportNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listExportNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewListExportNoteRepo(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListExportNoteRepo(%v, %v) = %v, want %v",
				tt.args.repo,
				tt.args.requester,
				got,
				tt.want)
		})
	}
}

func Test_listExportNoteBiz_ListExportNote(t *testing.T) {
	type fields struct {
		repo      ListExportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *exportnotemodel.Filter
		paging *common.Paging
	}
	mockRepo := new(mockListExportNoteRepo)
	mockRequest := new(mockRequester)
	mockFilter := exportnotemodel.Filter{
		SearchKey: "",
		MinPrice:  nil,
		MaxPrice:  nil,
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockListExportNote := make([]exportnotemodel.ExportNote, 0)
	var mockEmptyListExportNote []exportnotemodel.ExportNote
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
			name: "List export note failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(false).
					Once()
			},
			want:    mockListExportNote,
			wantErr: true,
		},
		{
			name: "List export note failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListExportNote",
						context.Background(),
						&mockFilter,
						&mockPaging,
					).
					Return(mockEmptyListExportNote, mockErr).
					Once()
			},
			want:    mockListExportNote,
			wantErr: true,
		},
		{
			name: "List export note successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListExportNote",
						context.Background(),
						&mockFilter,
						&mockPaging,
					).
					Return(mockListExportNote, nil).
					Once()
			},
			want:    mockListExportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listExportNoteBiz{
				requester: tt.fields.requester,
				repo:      tt.fields.repo,
			}

			tt.mock()

			got, err := biz.ListExportNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListExportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListExportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListExportNote() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
