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

type mockSeeDetailInventoryCheckNoteRepo struct {
	mock.Mock
}

func (m *mockSeeDetailInventoryCheckNoteRepo) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string) (*inventorychecknotemodel.InventoryCheckNote, error) {
	args := m.Called(ctx, inventoryCheckNoteId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*inventorychecknotemodel.InventoryCheckNote), args.Error(1)
}

func TestNewSeeDetailImportNoteBiz(t *testing.T) {
	type args struct {
		repo      SeeDetailInventoryCheckNoteRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeDetailInventoryCheckNoteRepo)

	tests := []struct {
		name string
		args args
		want *seeDetailInventoryCheckNoteBiz
	}{
		{
			name: "Create object has type SeeDetailInventoryCheckNoteBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeDetailInventoryCheckNoteBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailImportNoteBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeDetailImportNoteBiz(%v, %v) = %v, want %v",
				tt.args.repo,
				tt.args.requester,
				got,
				tt.want)
		})
	}
}

func Test_seeDetailInventoryCheckNoteBiz_SeeDetailInventoryCheckNote(t *testing.T) {
	type fields struct {
		repo      SeeDetailInventoryCheckNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx                  context.Context
		inventoryCheckNoteId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeDetailInventoryCheckNoteRepo)

	inventoryCheckNoteId := mock.Anything
	inventoryCheckNote := inventorychecknotemodel.InventoryCheckNote{
		Id: inventoryCheckNoteId,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *inventorychecknotemodel.InventoryCheckNote
		wantErr bool
	}{
		{
			name: "See inventory check note failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryCheckNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &inventoryCheckNote,
			wantErr: true,
		},
		{
			name: "See inventory check note failed because can't get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryCheckNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeDetailInventoryCheckNote",
						context.Background(),
						inventoryCheckNoteId).
					Return(nil, mockErr).
					Once()
			},
			want:    &inventoryCheckNote,
			wantErr: true,
		},
		{
			name: "See inventory check note successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: inventoryCheckNoteId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InventoryCheckNoteViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeDetailInventoryCheckNote",
						context.Background(),
						inventoryCheckNoteId).
					Return(&inventoryCheckNote, nil).
					Once()
			},
			want:    &inventoryCheckNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeDetailInventoryCheckNoteBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeDetailInventoryCheckNote(tt.args.ctx, tt.args.inventoryCheckNoteId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeDetailInventoryCheckNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeDetailInventoryCheckNote() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeDetailInventoryCheckNote() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
