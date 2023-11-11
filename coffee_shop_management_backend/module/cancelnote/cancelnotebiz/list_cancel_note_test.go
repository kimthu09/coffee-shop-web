package cancelnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCancelNoteRepo struct {
	mock.Mock
}

func (m *mockListCancelNoteRepo) ListCancelNote(
	ctx context.Context,
	filter *cancelnotemodel.Filter,
	paging *common.Paging) ([]cancelnotemodel.CancelNote, error) {
	args := m.Called(ctx, filter, paging)
	return args.Get(0).([]cancelnotemodel.CancelNote), args.Error(1)
}

func TestNewListCancelNoteRepo(t *testing.T) {
	type args struct {
		repo      ListCancelNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListCancelNoteRepo)

	tests := []struct {
		name string
		args args
		want *listCancelNoteBiz
	}{
		{
			name: "Create object has type ListCancelNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listCancelNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewListCancelNoteRepo(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListCancelNoteRepo(%v, %v) = %v, want %v",
				tt.args.repo,
				tt.args.requester,
				got,
				tt.want)
		})
	}
}

func Test_listCancelNoteBiz_ListCancelNote(t *testing.T) {
	type fields struct {
		requester middleware.Requester
		repo      ListCancelNoteRepo
	}
	type args struct {
		ctx    context.Context
		filter *cancelnotemodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListCancelNoteRepo)
	mockRequest := new(mockRequester)
	mockFilter := cancelnotemodel.Filter{
		SearchKey: "",
		MinPrice:  nil,
		MaxPrice:  nil,
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockListCancelNote := make([]cancelnotemodel.CancelNote, 0)
	var mockEmptyListCancelNote []cancelnotemodel.CancelNote
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []cancelnotemodel.CancelNote
		wantErr bool
	}{
		{
			name: "List cancel note failed because user is not allowed",
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
					On("IsHasFeature", common.CancelNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    mockListCancelNote,
			wantErr: true,
		},
		{
			name: "List cancel note failed because can not get data from database",
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
					On("IsHasFeature", common.CancelNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListCancelNote",
						context.Background(),
						&mockFilter,
						&mockPaging,
					).
					Return(mockEmptyListCancelNote, mockErr).
					Once()
			},
			want:    mockListCancelNote,
			wantErr: true,
		},
		{
			name: "List cancel note successfully",
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
					On("IsHasFeature", common.CancelNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListCancelNote",
						context.Background(),
						&mockFilter,
						&mockPaging,
					).
					Return(mockListCancelNote, nil).
					Once()
			},
			want:    mockListCancelNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listCancelNoteBiz{
				requester: tt.fields.requester,
				repo:      tt.fields.repo,
			}

			tt.mock()

			got, err := biz.ListCancelNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCancelNote(%v, %v, %v) error = %v, wantErr %v",
					tt.args.ctx,
					tt.args.filter,
					tt.args.paging, err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListCancelNote(%v, %v, %v) error = %v, wantErr %v",
					tt.args.ctx,
					tt.args.filter,
					tt.args.paging, err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCancelNote(%v, %v, %v) want = %v, got %v",
					tt.args.ctx,
					tt.args.filter,
					tt.args.paging,
					tt.want,
					got)
			}
		})
	}
}
