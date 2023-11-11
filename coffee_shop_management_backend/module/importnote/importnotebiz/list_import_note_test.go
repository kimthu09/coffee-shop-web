package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListImportNoteRepo struct {
	mock.Mock
}

func (m *mockListImportNoteRepo) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	args := m.Called(ctx, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]importnotemodel.ImportNote), args.Error(1)
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
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewListImportNoteBiz(t *testing.T) {
	type args struct {
		repo      ListImportNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListImportNoteRepo)

	tests := []struct {
		name string
		args args
		want *listImportNoteBiz
	}{
		{
			name: "Create object has type ListImportNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listImportNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListExportNoteBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listImportNoteBiz_ListImportNote(t *testing.T) {
	type fields struct {
		repo      ListImportNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *importnotemodel.Filter
		paging *common.Paging
	}
	mockRepo := new(mockListImportNoteRepo)
	mockRequest := new(mockRequester)
	filter := importnotemodel.Filter{
		SearchKey: "",
		MinPrice:  nil,
		MaxPrice:  nil,
		Status:    "",
	}
	paging := common.Paging{
		Page: 1,
	}
	listImportNote := make([]importnotemodel.ImportNote, 0)
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
			name: "List import note failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listImportNote,
			wantErr: true,
		},
		{
			name: "List import note failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListImportNote",
						context.Background(),
						&filter,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    listImportNote,
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
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListImportNote",
						context.Background(),
						&filter,
						&paging,
					).
					Return(listImportNote, nil).
					Once()
			},
			want:    listImportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listImportNoteBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListImportNote(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListImportNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListImportNote() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
