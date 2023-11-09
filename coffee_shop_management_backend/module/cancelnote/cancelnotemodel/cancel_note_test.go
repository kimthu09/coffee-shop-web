package cancelnotemodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCancelNote_TableName(t *testing.T) {
	type fields struct {
		Id         string
		TotalPrice float32
		CreateAt   *time.Time
		CreateBy   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CancelNote successfully",
			fields: fields{
				Id:         mock.Anything,
				TotalPrice: 0,
				CreateAt:   nil,
				CreateBy:   mock.Anything,
			},
			want: common.TableCancelNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := &CancelNote{
				Id:         tt.fields.Id,
				TotalPrice: tt.fields.TotalPrice,
				CreateAt:   tt.fields.CreateAt,
				CreateBy:   tt.fields.CreateBy,
			}
			assert.Equalf(t, tt.want, ca.TableName(), "TableName()")
		})
	}
}
