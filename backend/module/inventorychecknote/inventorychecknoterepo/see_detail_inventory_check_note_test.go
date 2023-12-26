package inventorychecknoterepo

import (
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockSeeDetailInventoryCheckNoteStore struct {
	mock.Mock
}

func (m *mockSeeDetailInventoryCheckNoteStore) ListInventoryCheckNoteDetail(
	ctx context.Context,
	inventoryCheckNoteId string) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error) {
	args := m.Called(ctx, inventoryCheckNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]inventorychecknotedetailmodel.InventoryCheckNoteDetail), args.Error(1)
}

type mockFindInventoryCheckNoteStore struct {
	mock.Mock
}

func (m *mockFindInventoryCheckNoteStore) FindInventoryCheckNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*inventorychecknotemodel.InventoryCheckNote, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*inventorychecknotemodel.InventoryCheckNote), args.Error(1)
}

func TestNewSeeDetailInventoryCheckNoteRepo(t *testing.T) {
	type args struct {
		inventoryCheckNoteStore       FindInventoryCheckNoteStore
		inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore
	}

	mockInventoryCheckNoteDetail := new(mockSeeDetailInventoryCheckNoteStore)
	mockInventoryCheckNote := new(mockFindInventoryCheckNoteStore)

	tests := []struct {
		name string
		args args
		want *seeDetailInventoryCheckNoteRepo
	}{
		{
			name: "Create object has type ListExportNoteRepo",
			args: args{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
			},
			want: &seeDetailInventoryCheckNoteRepo{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailInventoryCheckNoteRepo(
				tt.args.inventoryCheckNoteStore,
				tt.args.inventoryCheckNoteDetailStore)
			assert.Equal(
				t, tt.want, got,
				"NewSeeDetailInventoryCheckNoteRepo(%v, %v)",
				tt.args.inventoryCheckNoteStore, tt.args.inventoryCheckNoteDetailStore)
		})
	}
}

func Test_seeDetailInventoryCheckNoteRepo_SeeDetailInventoryCheckNote(t *testing.T) {
	type fields struct {
		inventoryCheckNoteStore       FindInventoryCheckNoteStore
		inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore
	}
	type args struct {
		ctx                  context.Context
		inventoryCheckNoteId string
	}

	mockInventoryCheckNoteDetail := new(mockSeeDetailInventoryCheckNoteStore)
	mockInventoryCheckNote := new(mockFindInventoryCheckNoteStore)

	inventoryNoteId := "Inventory001"
	moreKeys := []string{"CreatedByUser"}
	createdBy := "user001"
	ingredientId := "Ingredient001"
	mockErr := errors.New(mock.Anything)
	inventoryNote := inventorychecknotemodel.InventoryCheckNote{
		Id:                inventoryNoteId,
		AmountDifferent:   100,
		AmountAfterAdjust: 120,
		CreatedBy:         createdBy,
		CreatedAt:         &time.Time{},
	}
	details := []inventorychecknotedetailmodel.InventoryCheckNoteDetail{
		{
			InventoryCheckNoteId: inventoryNoteId,
			IngredientId:         ingredientId,
			Initial:              20,
			Difference:           100,
			Final:                120,
		},
	}
	finalInventoryNote := inventorychecknotemodel.InventoryCheckNote{
		Id:                inventoryNoteId,
		AmountDifferent:   100,
		AmountAfterAdjust: 120,
		CreatedBy:         createdBy,
		CreatedAt:         &time.Time{},
		Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetail{
			{
				InventoryCheckNoteId: inventoryNoteId,
				IngredientId:         ingredientId,
				Initial:              20,
				Difference:           100,
				Final:                120,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *inventorychecknotemodel.InventoryCheckNote
		wantErr bool
	}{
		{
			name: "See inventory check note detail failed because can not find export note",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryNoteId,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"FindInventoryCheckNote",
						context.Background(),
						map[string]interface{}{"id": inventoryNoteId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See inventory check note detail failed because can not find export note",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryNoteId,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"FindInventoryCheckNote",
						context.Background(),
						map[string]interface{}{"id": inventoryNoteId},
						moreKeys,
					).
					Return(&inventoryNote, nil).
					Once()

				mockInventoryCheckNoteDetail.
					On(
						"ListInventoryCheckNoteDetail",
						context.Background(),
						inventoryNoteId,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See inventory check note detail successfully",
			fields: fields{
				inventoryCheckNoteStore:       mockInventoryCheckNote,
				inventoryCheckNoteDetailStore: mockInventoryCheckNoteDetail,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryNoteId,
			},
			mock: func() {
				mockInventoryCheckNote.
					On(
						"FindInventoryCheckNote",
						context.Background(),
						map[string]interface{}{"id": inventoryNoteId},
						moreKeys,
					).
					Return(&inventoryNote, nil).
					Once()

				mockInventoryCheckNoteDetail.
					On(
						"ListInventoryCheckNoteDetail",
						context.Background(),
						inventoryNoteId,
					).
					Return(details, nil).
					Once()
			},
			want:    &finalInventoryNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeDetailInventoryCheckNoteRepo{
				inventoryCheckNoteStore:       tt.fields.inventoryCheckNoteStore,
				inventoryCheckNoteDetailStore: tt.fields.inventoryCheckNoteDetailStore,
			}

			tt.mock()

			got, err := repo.SeeDetailInventoryCheckNote(tt.args.ctx, tt.args.inventoryCheckNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeDetailInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeDetailInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "SeeDetailInventoryCheckNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
