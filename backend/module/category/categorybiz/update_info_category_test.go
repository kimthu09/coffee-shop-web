package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoCategory struct {
	mock.Mock
}

func (m *mockUpdateInfoCategory) FindCategory(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*categorymodel.Category, error) {
	args := m.Called(ctx, conditions, moreKeys)
	return args.Get(0).(*categorymodel.Category), args.Error(1)
}
func (m *mockUpdateInfoCategory) UpdateInfoCategory(
	ctx context.Context,
	id string,
	data *categorymodel.CategoryUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateInfoCategoryBiz(t *testing.T) {
	type args struct {
		store     UpdateInfoCategoryStore
		requester middleware.Requester
	}

	mockStore := new(mockUpdateInfoCategory)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateInfoCategoryBiz
	}{
		{
			name: "Create object has type UpdateInfoCategoryBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &updateInfoCategoryBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		got := NewUpdateInfoCategoryBiz(
			tt.args.store,
			tt.args.requester,
		)

		assert.Equal(t, tt.want, got, "NewUpdateInfoCategoryBiz() = %v, want %v", got, tt.want)
	}
}

func Test_updateInfoCategoryBiz_UpdateInfoCategory(t *testing.T) {
	type fields struct {
		store     UpdateInfoCategoryStore
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *categorymodel.CategoryUpdateInfo
	}

	mockStore := new(mockUpdateInfoCategory)
	mockRequest := new(mockRequester)

	id := mock.Anything
	name := "valid name"
	emptyName := ""
	description := mock.Anything
	categoryUpdate := categorymodel.CategoryUpdateInfo{
		Name:        &name,
		Description: &description,
	}
	invalidCategoryUpdate := categorymodel.CategoryUpdateInfo{
		Name: &emptyName,
	}
	categoryFound := categorymodel.Category{
		Id:            id,
		Name:          mock.Anything,
		Description:   mock.Anything,
		AmountProduct: 0,
	}
	mockErr := errors.New(mock.Anything)
	var moreKeys []string

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update info category failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update info category failed because data is invalid",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &invalidCategoryUpdate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update info category failed because category is not exist",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"FindCategory",
						context.Background(),
						map[string]interface{}{"id": id},
						moreKeys).
					Return(&categoryFound, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update info category failed because can not save to database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"FindCategory",
						context.Background(),
						map[string]interface{}{"id": id},
						moreKeys).
					Return(&categoryFound, nil).
					Once()

				mockStore.
					On(
						"UpdateInfoCategory",
						context.Background(),
						id,
						&categoryUpdate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update info category successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdate,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryUpdateInfoFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"FindCategory",
						context.Background(),
						map[string]interface{}{"id": id},
						moreKeys).
					Return(&categoryFound, nil).
					Once()

				mockStore.
					On(
						"UpdateInfoCategory",
						context.Background(),
						id,
						&categoryUpdate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateInfoCategoryBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}
			tt.mock()

			err := biz.UpdateInfoCategory(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateInfoCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateInfoCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
