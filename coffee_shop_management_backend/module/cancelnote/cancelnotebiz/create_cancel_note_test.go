package cancelnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) GenerateId() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockIdGenerator) IdProcess(id *string) (*string, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

type mockCreateCancelNoteRepo struct {
	mock.Mock
}

func (m *mockCreateCancelNoteRepo) GetPriceIngredient(
	ctx context.Context,
	ingredientId string) (*float32, error) {
	args := m.Called(ctx, ingredientId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float32), args.Error(1)
}

func (m *mockCreateCancelNoteRepo) HandleCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateCancelNoteRepo) HandleIngredientDetail(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateCancelNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]float32,
) error {
	args := m.Called(ctx, ingredientTotalAmountNeedUpdate)
	return args.Error(0)
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

func TestNewCreateCancelNoteBiz(t *testing.T) {
	type args struct {
		generator generator.IdGenerator
		repo      CreateCancelNoteRepo
		requester middleware.Requester
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateCancelNoteRepo)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createCancelNoteBiz
	}{
		{
			name: "Create object has type CreateCancelNoteBiz",
			args: args{
				generator: mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createCancelNoteBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateCancelNoteBiz(
				tt.args.generator,
				tt.args.repo,
				tt.args.requester,
			)

			assert.Equal(t, tt.want, got, "NewCreateCancelNoteBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createCancelNoteBiz_CreateCancelNote(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateCancelNoteRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *cancelnotemodel.CancelNoteCreate
	}

	mockGenerator := new(mockIdGenerator)
	mockRepo := new(mockCreateCancelNoteRepo)
	mockRequest := new(mockRequester)

	mockReason := cancelnotedetailmodel.Damaged
	invalidCancelNoteCreate := cancelnotemodel.CancelNoteCreate{
		Id:                      nil,
		CancelNoteCreateDetails: nil,
	}
	mockCancelNoteCreate := cancelnotemodel.CancelNoteCreate{
		Id: nil,
		CancelNoteCreateDetails: []cancelnotedetailmodel.CancelNoteDetailCreate{
			{
				IngredientId: "Ing001",
				ExpiryDate:   "08/11/2023",
				Reason:       &mockReason,
				AmountCancel: 20,
			},
			{
				IngredientId: "Ing001",
				ExpiryDate:   "09/11/2023",
				Reason:       &mockReason,
				AmountCancel: 10,
			},
		},
	}
	mockId := "123456789"
	mockErr := errors.New(mock.Anything)
	mockPrice := float32(10000)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create cancel note failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &invalidCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because can not generate Id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because can not get price ingredient",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because can not handle cancel note successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockCancelNoteCreate.CancelNoteCreateDetails))

				mockRepo.
					On("HandleCancelNote", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because can not handle ingredient detail successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockCancelNoteCreate.CancelNoteCreateDetails))

				mockRepo.
					On("HandleCancelNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note failed because can not update ingredient total amount successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockCancelNoteCreate.CancelNoteCreateDetails))

				mockRepo.
					On("HandleCancelNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientTotalAmount", context.Background(), mock.Anything).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create cancel note successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &mockCancelNoteCreate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CancelNoteCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", mock.Anything).
					Return(&mockId, nil).
					Once()

				mockRepo.
					On("GetPriceIngredient", context.Background(), mock.Anything).
					Return(&mockPrice, nil).
					Times(len(mockCancelNoteCreate.CancelNoteCreateDetails))

				mockRepo.
					On("HandleCancelNote", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientDetail", context.Background(), mock.Anything).
					Return(nil).
					Once()

				mockRepo.
					On("HandleIngredientTotalAmount", context.Background(), mock.Anything).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createCancelNoteBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateCancelNote(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "CreateCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
