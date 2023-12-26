package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeDetailToppingStore struct {
	mock.Mock
}

func (m *mockSeeDetailToppingStore) FindTopping(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productmodel.Topping, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Topping), args.Error(1)
}

func TestNewSeeDetailToppingRepo(t *testing.T) {
	type args struct {
		store SeeDetailToppingStore
	}

	mockStore := new(mockSeeDetailToppingStore)

	tests := []struct {
		name string
		args args
		want *seeDetailToppingRepo
	}{
		{
			name: "Create object has type seeDetailToppingRepo",
			args: args{
				store: mockStore,
			},
			want: &seeDetailToppingRepo{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeDetailToppingRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewSeeDetailToppingRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeDetailToppingRepo_SeeDetailTopping(t *testing.T) {
	type fields struct {
		store SeeDetailToppingStore
	}
	type args struct {
		ctx       context.Context
		toppingId string
	}

	mockStore := new(mockSeeDetailToppingStore)

	toppingId := "Topping001"

	mockErr := errors.New(mock.Anything)

	mockTopping := &productmodel.Topping{}
	moreKeys := []string{"Recipe.Details.Ingredient"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Topping
		wantErr bool
	}{
		{
			name: "See detail topping failed because can not get data from the mockStore",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:       context.Background(),
				toppingId: toppingId,
			},
			mock: func() {
				mockStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{"id": toppingId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See detail topping successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:       context.Background(),
				toppingId: toppingId,
			},
			mock: func() {
				mockStore.
					On("FindTopping",
						context.Background(),
						map[string]interface{}{"id": toppingId},
						moreKeys).
					Return(mockTopping, nil).
					Once()
			},
			want:    mockTopping,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeDetailToppingRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.SeeDetailTopping(tt.args.ctx, tt.args.toppingId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeDetailTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeDetailTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeDetailTopping() = %v, want %v", got, tt.want)
			}
		})
	}
}
