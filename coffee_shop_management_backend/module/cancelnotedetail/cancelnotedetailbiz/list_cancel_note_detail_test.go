package cancelnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCancelNoteDetailStore struct {
	mock.Mock
}

func (m *mockListCancelNoteDetailStore) ListCancelNoteDetail(
	ctx context.Context,
	cancelNoteId string,
	paging *common.Paging) ([]cancelnotedetailmodel.CancelNoteDetail, error) {
	args := m.Called(ctx, cancelNoteId, paging)
	return args.Get(0).([]cancelnotedetailmodel.CancelNoteDetail), args.Error(1)
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

func TestNewListCancelNoteDetailBiz(t *testing.T) {
	type args struct {
		store     ListCancelNoteDetailStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListCancelNoteDetailStore)

	tests := []struct {
		name string
		args args
		want *listCancelNoteDetailBiz
	}{
		{
			name: "Create object has type ListCancelNoteDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listCancelNoteDetailBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCancelNoteDetailBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListCancelNoteDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listCancelNoteDetailBiz_ListCancelNoteDetail(t *testing.T) {
	type fields struct {
		store     ListCancelNoteDetailStore
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		cancelNoteId string
		paging       *common.Paging
	}

	mockStore := new(mockListCancelNoteDetailStore)
	mockRequest := new(mockRequester)
	mockPaging := common.Paging{
		Page: 1,
	}
	mockCancelNoteId := mock.Anything
	mockListCancelNoteDetail := make([]cancelnotedetailmodel.CancelNoteDetail, 0)
	var mockEmptyListCancelNoteDetail []cancelnotedetailmodel.CancelNoteDetail
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []cancelnotedetailmodel.CancelNoteDetail
		wantErr bool
	}{
		{
			name: "List cancel note detail failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: mockCancelNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(false).
					Once()
			},
			want:    mockListCancelNoteDetail,
			wantErr: true,
		},
		{
			name: "List cancel note detail failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: mockCancelNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCancelNoteDetail",
						context.Background(),
						mockCancelNoteId,
						&mockPaging,
					).
					Return(mockEmptyListCancelNoteDetail, mockErr).
					Once()
			},
			want:    mockListCancelNoteDetail,
			wantErr: true,
		},
		{
			name: "List cancel note successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: mockCancelNoteId,
				paging:       &mockPaging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", mock.Anything).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCancelNoteDetail",
						context.Background(),
						mockCancelNoteId,
						&mockPaging,
					).
					Return(mockListCancelNoteDetail, nil).
					Once()
			},
			want:    mockListCancelNoteDetail,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listCancelNoteDetailBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListCancelNoteDetail(tt.args.ctx, tt.args.cancelNoteId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCancelNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListCancelNoteDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCancelNoteDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
