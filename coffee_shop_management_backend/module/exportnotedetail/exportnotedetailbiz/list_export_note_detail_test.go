package exportnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListExportNoteDetailStore struct {
	mock.Mock
}

func (m *mockListExportNoteDetailStore) ListExportNoteDetail(
	ctx context.Context,
	exportNoteId string,
	paging *common.Paging) ([]exportnotedetailmodel.ExportNoteDetail, error) {
	args := m.Called(ctx, exportNoteId, paging)
	return args.Get(0).([]exportnotedetailmodel.ExportNoteDetail), args.Error(1)
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
	args := m.Called()
	return args.Bool(0)
}

func TestNewListExportNoteDetailBiz(t *testing.T) {
	type args struct {
		store     ListExportNoteDetailStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListExportNoteDetailStore)

	tests := []struct {
		name string
		args args
		want *listExportNoteDetailBiz
	}{
		{
			name: "Create object has type ListExportNoteDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listExportNoteDetailBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListExportNoteDetailBiz(tt.args.store, tt.args.requester)

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

func Test_listExportNoteDetailBiz_ListExportNoteDetail(t *testing.T) {
	type fields struct {
		store     ListExportNoteDetailStore
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		exportNoteId string
		paging       *common.Paging
	}

	mockStore := new(mockListExportNoteDetailStore)
	mockRequest := new(mockRequester)
	mockPaging := common.Paging{
		Page: 1,
	}
	mockExportNoteId := mock.Anything
	mockListExportNoteDetail := make([]exportnotedetailmodel.ExportNoteDetail, 0)
	var mockEmptyListExportNoteDetail []exportnotedetailmodel.ExportNoteDetail
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []exportnotedetailmodel.ExportNoteDetail
		wantErr bool
	}{
		{
			name: "List export note detail failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: mockExportNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(false).
					Once()
			},
			want:    mockListExportNoteDetail,
			wantErr: true,
		},
		{
			name: "List export note detail failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: mockExportNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockStore.
					On(
						"ListExportNoteDetail",
						context.Background(),
						mockExportNoteId,
						&mockPaging,
					).
					Return(mockEmptyListExportNoteDetail, mockErr).
					Once()
			},
			want:    mockListExportNoteDetail,
			wantErr: true,
		},
		{
			name: "List export note successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				exportNoteId: mockExportNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockStore.
					On(
						"ListExportNoteDetail",
						context.Background(),
						mockExportNoteId,
						&mockPaging,
					).
					Return(mockListExportNoteDetail, nil).
					Once()
			},
			want:    mockListExportNoteDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listExportNoteDetailBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListExportNoteDetail(tt.args.ctx, tt.args.exportNoteId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListExportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListExportNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListExportNoteDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
