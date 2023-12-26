package inventorychecknotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListInventoryCheckNoteRepo struct {
	mock.Mock
}

func (m *mockListInventoryCheckNoteRepo) ListInventoryCheckNote(
	ctx context.Context,
	filter *inventorychecknotemodel.Filter,
	paging *common.Paging) ([]inventorychecknotemodel.InventoryCheckNote, error) {
	args := m.Called(ctx, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]inventorychecknotemodel.InventoryCheckNote), args.Error(1)
}
func TestNewListInventoryCheckNoteBiz(t *testing.T) {
	type args struct {
		repo      ListInventoryCheckNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListInventoryCheckNoteRepo)

	tests := []struct {
		name string
		args args
		want *listInventoryCheckNoteBiz
	}{
		{
			name: "Create object has type ListInventoryCheckNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listInventoryCheckNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListInventoryCheckNoteBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListInventoryCheckNoteBiz(%v, %v) = %v, want %v",
				tt.args.repo,
				tt.args.requester,
				got,
				tt.want)

		})
	}
}

func Test_listInventoryCheckNoteBiz_ListInventoryCheckNote(t *testing.T) {
	type fields struct {
		repo      ListInventoryCheckNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *inventorychecknotemodel.Filter
		paging *common.Paging
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListInventoryCheckNoteRepo)
	mockFilter := inventorychecknotemodel.Filter{
		SearchKey: "",
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockListInventoryCheckNote := make([]inventorychecknotemodel.InventoryCheckNote, 0)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []inventorychecknotemodel.InventoryCheckNote
		wantErr bool
	}{
		{
			name: "List inventory check note failed because user is not allowed",
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
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List inventory check note failed because can not get data from database",
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
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListInventoryCheckNote",
						context.Background(),
						&mockFilter,
						&mockPaging).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List inventory check note successfully",
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
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListInventoryCheckNote",
						context.Background(),
						&mockFilter,
						&mockPaging).
					Return(mockListInventoryCheckNote, nil).
					Once()
			},
			want:    mockListInventoryCheckNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listInventoryCheckNoteBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListInventoryCheckNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListInventoryCheckNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListInventoryCheckNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListInventoryCheckNote() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
