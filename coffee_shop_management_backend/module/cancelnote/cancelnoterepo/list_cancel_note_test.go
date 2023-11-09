package cancelnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCancelNoteStore struct {
	mock.Mock
}

func (m *mockListCancelNoteStore) ListCancelNote(
	ctx context.Context,
	filter *cancelnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
) ([]cancelnotemodel.CancelNote, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]cancelnotemodel.CancelNote), args.Error(1)
}

func TestNewListCancelNoteRepo(t *testing.T) {
	type args struct {
		store ListCancelNoteStore
	}

	mockCancelNote := new(mockListCancelNoteStore)

	tests := []struct {
		name string
		args args
		want *listCancelNoteRepo
	}{
		{
			name: "Create object has type ListCancelNoteRepo",
			args: args{
				store: mockCancelNote,
			},
			want: &listCancelNoteRepo{
				store: mockCancelNote,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCancelNoteRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListCancelNoteRepo(%v)", tt.args.store)
		})
	}
}

func Test_listCancelNoteRepo_ListCancelNote(t *testing.T) {
	type fields struct {
		store ListCancelNoteStore
	}

	type args struct {
		ctx    context.Context
		filter *cancelnotemodel.Filter
		paging *common.Paging
	}

	mockCancelNote := new(mockListCancelNoteStore)
	mockFilter := cancelnotemodel.Filter{
		SearchKey: "",
		MinPrice:  nil,
		MaxPrice:  nil,
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockCancelNotes := []cancelnotemodel.CancelNote{
		{
			Id:         mock.Anything,
			TotalPrice: 0,
			CreateAt:   nil,
			CreateBy:   mock.Anything,
		},
		{
			Id:         mock.Anything,
			TotalPrice: 0,
			CreateAt:   nil,
			CreateBy:   mock.Anything,
		},
	}
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
			name: "List cancel note successfully",
			fields: fields{
				store: mockCancelNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockCancelNote.
					On(
						"ListCancelNote",
						context.Background(),
						&mockFilter,
						mock.Anything,
						&mockPaging).
					Return(mockCancelNotes, nil).
					Once()
			},
			want:    mockCancelNotes,
			wantErr: false,
		},
		{
			name: "List cancel note failed",
			fields: fields{
				store: mockCancelNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockCancelNote.
					On(
						"ListCancelNote",
						context.Background(),
						&mockFilter,
						mock.Anything,
						&mockPaging).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listCancelNoteRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListCancelNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCancelNote() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"ListCancelNote() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCancelNote() got = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
