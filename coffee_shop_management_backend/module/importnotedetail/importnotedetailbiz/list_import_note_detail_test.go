package importnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListImportNoteDetailStore struct {
	mock.Mock
}

func (m *mockListImportNoteDetailStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) ([]importnotedetailmodel.ImportNoteDetail, error) {
	args := m.Called(ctx, importNoteId, paging)
	return args.Get(0).([]importnotedetailmodel.ImportNoteDetail),
		args.Error(1)
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

func TestNewListImportNoteDetailBiz(t *testing.T) {
	type args struct {
		store     ListImportNoteDetailStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListImportNoteDetailStore)

	tests := []struct {
		name string
		args args
		want *listImportNoteDetailBiz
	}{
		{
			name: "Create object has type ListImportNoteDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listImportNoteDetailBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteDetailBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListExportNoteDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listImportNoteDetailBiz_ListImportNoteDetail(t *testing.T) {
	type fields struct {
		store     ListImportNoteDetailStore
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		importNoteId string
		paging       *common.Paging
	}

	mockStore := new(mockListImportNoteDetailStore)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	importNoteId := mock.Anything
	listImportNoteDetail := make([]importnotedetailmodel.ImportNoteDetail, 0)
	var emptyListImportNoteDetail []importnotedetailmodel.ImportNoteDetail
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotedetailmodel.ImportNoteDetail
		wantErr bool
	}{
		{
			name: "List import note detail failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listImportNoteDetail,
			wantErr: true,
		},
		{
			name: "List import note detail failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListImportNoteDetail",
						context.Background(),
						importNoteId,
						&paging,
					).
					Return(emptyListImportNoteDetail, mockErr).
					Once()
			},
			want:    listImportNoteDetail,
			wantErr: true,
		},
		{
			name: "List import note detail successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ImportNoteViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListImportNoteDetail",
						context.Background(),
						importNoteId,
						&paging,
					).
					Return(listImportNoteDetail, nil).
					Once()
			},
			want:    listImportNoteDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listImportNoteDetailBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListImportNoteDetail(tt.args.ctx, tt.args.importNoteId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListImportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListImportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListImportNoteDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
