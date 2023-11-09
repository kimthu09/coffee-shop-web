package cancelnoterepo

import (
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateCancelNoteStore struct {
	mock.Mock
}

func (m *mockCreateCancelNoteStore) CreateCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreatCancelNoteDetailStore struct {
	mock.Mock
}

func (m *mockCreatCancelNoteDetailStore) CreateListCancelNoteDetail(
	ctx context.Context,
	data []cancelnotedetailmodel.CancelNoteDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockUpdateIngredientStore struct {
	mock.Mock
}

func (m *mockUpdateIngredientStore) GetPriceIngredient(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*float32, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}
func (m *mockUpdateIngredientStore) UpdateAmountIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdateAmount) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockUpdateIngredientDetailStore struct {
	mock.Mock
}

func (m *mockUpdateIngredientDetailStore) FindIngredientDetail(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*ingredientdetailmodel.IngredientDetail, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ingredientdetailmodel.IngredientDetail),
		args.Error(1)
}
func (m *mockUpdateIngredientDetailStore) UpdateIngredientDetail(
	ctx context.Context,
	ingredientId string,
	expiryDate string,
	data *ingredientdetailmodel.IngredientDetailUpdate,
) error {
	args := m.Called(ctx, ingredientId, expiryDate, data)
	return args.Error(0)
}

func TestNewCreateCancelNoteRepo(t *testing.T) {
	type args struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	tests := []struct {
		name string
		args args
		want *createCancelNoteRepo
	}{
		{
			name: "Create object has type CreateCancelNoteRepo",
			args: args{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			want: &createCancelNoteRepo{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateCancelNoteRepo(
				tt.args.cancelNoteStore,
				tt.args.cancelNoteDetailStore,
				tt.args.ingredientStore,
				tt.args.ingredientDetailStore)

			assert.Equal(t,
				tt.want,
				got,
				"NewCreateCancelNoteRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_createCancelNoteRepo_GetPriceIngredient(t *testing.T) {
	type fields struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx          context.Context
		ingredientId string
	}

	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockIngredientId := "mockId"
	mockPrice := float32(100)
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *float32
		wantErr bool
	}{
		{
			name: "Get price of ingredient successfully",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: mockIngredientId,
			},
			mock: func() {
				mockIngredient.
					On(
						"GetPriceIngredient",
						context.Background(),
						map[string]interface{}{"id": mockIngredientId},
						mock.Anything).
					Return(
						&mockPrice,
						nil).
					Once()
			},
			want:    &mockPrice,
			wantErr: false,
		},
		{
			name: "Get price of ingredient failed with error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: mockIngredientId,
			},
			mock: func() {
				mockIngredient.
					On(
						"GetPriceIngredient",
						context.Background(),
						map[string]interface{}{"id": mockIngredientId},
						mock.Anything).
					Return(
						nil,
						mockErr,
					).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCancelNoteRepo{
				cancelNoteStore:       tt.fields.cancelNoteStore,
				cancelNoteDetailStore: tt.fields.cancelNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			got, err := repo.GetPriceIngredient(tt.args.ctx, tt.args.ingredientId)

			if tt.wantErr {
				assert.NotNil(t, err, "GetPriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetPriceIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetPriceIngredient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createCancelNoteRepo_HandleCancelNote(t *testing.T) {
	type fields struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *cancelnotemodel.CancelNoteCreate
	}

	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockCancelNoteId := "mockId"
	mockReason := cancelnotedetailmodel.Damaged
	mockDetails := []cancelnotedetailmodel.CancelNoteDetailCreate{
		{
			CancelNoteId: mockCancelNoteId,
			IngredientId: "mockIng1",
			ExpiryDate:   "mockDate",
			Reason:       &mockReason,
			AmountCancel: 10,
		},
		{
			CancelNoteId: mockCancelNoteId,
			IngredientId: "mockIng2",
			ExpiryDate:   "mockDate",
			Reason:       &mockReason,
			AmountCancel: 10,
		},
	}
	mockCancelNoteCreate := cancelnotemodel.CancelNoteCreate{
		Id:                      &mockCancelNoteId,
		TotalPrice:              float32(1000),
		CreateBy:                "mockCreateBy",
		CancelNoteCreateDetails: mockDetails,
	}
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create cancel note and cancel note detail successfully",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockCancelNote.
					On(
						"CreateCancelNote",
						context.Background(),
						&mockCancelNoteCreate,
					).
					Return(nil).
					Once()

				mockCancelNoteDetail.
					On(
						"CreateListCancelNoteDetail",
						context.Background(),
						mockDetails,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "Create cancel note failed and return error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockCancelNote.
					On(
						"CreateCancelNote",
						context.Background(),
						&mockCancelNoteCreate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note successfully but " +
				"create cancel note detail failed then return error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockCancelNote.
					On(
						"CreateCancelNote",
						context.Background(),
						&mockCancelNoteCreate,
					).
					Return(nil).
					Once()

				mockCancelNoteDetail.
					On(
						"CreateListCancelNoteDetail",
						context.Background(),
						mockDetails,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCancelNoteRepo{
				cancelNoteStore:       tt.fields.cancelNoteStore,
				cancelNoteDetailStore: tt.fields.cancelNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.HandleCancelNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createCancelNoteRepo_HandleIngredientTotalAmount(t *testing.T) {
	type fields struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx                             context.Context
		ingredientTotalAmountNeedUpdate map[string]float32
	}

	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockMap := map[string]float32{
		"ingredient1": 12,
		"ingredient2": 30,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update total amount of ingredient successfully",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						mock.Anything,
						mock.Anything,
					).
					Return(nil).
					Times(len(mockMap))
			},
			wantErr: false,
		},
		{
			name: "Update total amount of ingredient failed and return error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:                             context.Background(),
				ingredientTotalAmountNeedUpdate: mockMap,
			},
			mock: func() {
				mockIngredient.
					On(
						"UpdateAmountIngredient",
						context.Background(),
						mock.Anything,
						mock.Anything,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCancelNoteRepo{
				cancelNoteStore:       tt.fields.cancelNoteStore,
				cancelNoteDetailStore: tt.fields.cancelNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.HandleIngredientTotalAmount(tt.args.ctx, tt.args.ingredientTotalAmountNeedUpdate)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleIngredientTotalAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createCancelNoteRepo_checkIngredientDetail(t *testing.T) {
	type fields struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *cancelnotedetailmodel.CancelNoteDetailCreate
	}

	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockCancelNoteId := "mockId"
	mockIngredientId := "mockIngId"
	mockExpiryData := "mockDate"
	mockReason := cancelnotedetailmodel.Damaged
	mockDetail := cancelnotedetailmodel.CancelNoteDetailCreate{
		CancelNoteId: mockCancelNoteId,
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		Reason:       &mockReason,
		AmountCancel: 10,
	}
	mockIngredientDetailHasBigAmountStock := ingredientdetailmodel.IngredientDetail{
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		Amount:       100,
	}
	mockIngredientDetailHasSmallAmountStock := ingredientdetailmodel.IngredientDetail{
		IngredientId: mockIngredientId,
		ExpiryDate:   mockExpiryData,
		Amount:       5,
	}
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "CancelNoteDetail check found no error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(&mockIngredientDetailHasBigAmountStock, nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "CancelNoteDetail check found amount cancel is over stock error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(&mockIngredientDetailHasSmallAmountStock, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "CancelNoteDetail check found database error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"FindIngredientDetail",
						context.Background(),
						map[string]interface{}{
							"ingredientId": mockDetail.IngredientId,
							"expiryDate":   mockDetail.ExpiryDate,
						},
						mock.Anything,
					).
					Return(nil, mockErr)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCancelNoteRepo{
				cancelNoteStore:       tt.fields.cancelNoteStore,
				cancelNoteDetailStore: tt.fields.cancelNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.checkIngredientDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "checkIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "checkIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createCancelNoteRepo_updateIngredientDetail(t *testing.T) {
	type fields struct {
		cancelNoteStore       CreateCancelNoteStore
		cancelNoteDetailStore CreateCancelNoteDetailStore
		ingredientStore       UpdateIngredientStore
		ingredientDetailStore UpdateIngredientDetailStore
	}
	type args struct {
		ctx  context.Context
		data *cancelnotedetailmodel.CancelNoteDetailCreate
	}

	mockCancelNote := new(mockCreateCancelNoteStore)
	mockCancelNoteDetail := new(mockCreatCancelNoteDetailStore)
	mockIngredient := new(mockUpdateIngredientStore)
	mockIngredientDetail := new(mockUpdateIngredientDetailStore)

	mockReason := cancelnotedetailmodel.Damaged
	mockDetail := cancelnotedetailmodel.CancelNoteDetailCreate{
		CancelNoteId: "mockId",
		IngredientId: "mockIngId",
		ExpiryDate:   "mockDate",
		Reason:       &mockReason,
		AmountCancel: 10,
	}
	mockErr := errors.New("mockErr")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "IngredientDetail update successfully",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						mockDetail.IngredientId,
						mockDetail.ExpiryDate,
						mock.Anything,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "IngredientDetail failed and return error",
			fields: fields{
				cancelNoteStore:       mockCancelNote,
				cancelNoteDetailStore: mockCancelNoteDetail,
				ingredientStore:       mockIngredient,
				ingredientDetailStore: mockIngredientDetail,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockDetail,
			},
			mock: func() {
				mockIngredientDetail.
					On(
						"UpdateIngredientDetail",
						context.Background(),
						mockDetail.IngredientId,
						mockDetail.ExpiryDate,
						mock.Anything,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createCancelNoteRepo{
				cancelNoteStore:       tt.fields.cancelNoteStore,
				cancelNoteDetailStore: tt.fields.cancelNoteDetailStore,
				ingredientStore:       tt.fields.ingredientStore,
				ingredientDetailStore: tt.fields.ingredientDetailStore,
			}

			tt.mock()

			err := repo.updateIngredientDetail(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "updateIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "updateIngredientDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
