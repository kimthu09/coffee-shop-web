package inventorychecknoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListInventoryCheckNoteStore struct {
	mock.Mock
}

func (m *mockListInventoryCheckNoteStore) ListInventoryCheckNote(
	ctx context.Context,
	filter *inventorychecknotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]inventorychecknotemodel.InventoryCheckNote, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]inventorychecknotemodel.InventoryCheckNote), args.Error(1)
}

func TestNewListInventoryCheckNoteRepo(t *testing.T) {
	type args struct {
		store ListInventoryCheckNoteStore
	}

	mockInventoryCheckNote := new(mockListInventoryCheckNoteStore)

	tests := []struct {
		name string
		args args
		want *listInventoryCheckNoteRepo
	}{
		{
			name: "Create object has type ListInventoryCheckNoteRepo",
			args: args{
				store: mockInventoryCheckNote,
			},
			want: &listInventoryCheckNoteRepo{
				store: mockInventoryCheckNote,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListInventoryCheckNoteRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListInventoryCheckNoteRepo(%v)", tt.args.store)
		})
	}
}

func Test_listInventoryCheckNoteRepo_ListInventoryCheckNote(t *testing.T) {
	type fields struct {
		store ListInventoryCheckNoteStore
	}
	type args struct {
		ctx    context.Context
		filter *inventorychecknotemodel.Filter
		paging *common.Paging
	}

	mockInventoryCheckNote := new(mockListInventoryCheckNoteStore)
	mockFilter := inventorychecknotemodel.Filter{
		SearchKey: "",
	}
	mockPaging := common.Paging{
		Page: 1,
	}
	mockExportNotes := []inventorychecknotemodel.InventoryCheckNote{
		{
			Id:                mock.Anything,
			AmountDifferent:   0,
			AmountAfterAdjust: 0,
			CreatedBy:         mock.Anything,
		},
		{
			Id:                mock.Anything,
			AmountDifferent:   0,
			AmountAfterAdjust: 0,
			CreatedBy:         mock.Anything,
		},
	}
	moreKeys := []string{"CreatedByUser"}
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
			name: "List inventory check note failed",
			fields: fields{
				store: mockInventoryCheckNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"ListInventoryCheckNote",
						context.Background(),
						&mockFilter,
						[]string{"InventoryCheckNote.id"},
						&mockPaging,
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List inventory check note successfully",
			fields: fields{
				store: mockInventoryCheckNote,
			},
			args: args{
				ctx:    context.Background(),
				filter: &mockFilter,
				paging: &mockPaging,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"ListInventoryCheckNote",
						context.Background(),
						&mockFilter,
						[]string{"InventoryCheckNote.id"},
						&mockPaging,
						moreKeys).
					Return(mockExportNotes, nil).
					Once()
			},
			want:    mockExportNotes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listInventoryCheckNoteRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListInventoryCheckNote(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListInventoryCheckNote() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"ListInventoryCheckNote() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListInventoryCheckNote() got = %v, want %v",
					got,
					tt.want,
				)
			}

		})
	}
}
