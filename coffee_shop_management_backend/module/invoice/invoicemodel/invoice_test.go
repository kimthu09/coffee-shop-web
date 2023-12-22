package invoicemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvoice_TableName(t *testing.T) {
	type fields struct {
		Id                  string
		CustomerId          string
		Customer            SimpleCustomer
		TotalPrice          float32
		AmountReceived      int
		AmountPriceUsePoint int
		CreatedBy           string
		CreatedByUser       usermodel.SimpleUser
		CreatedAt           *time.Time
		Details             []invoicedetailmodel.InvoiceDetail
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InvoiceCreate successfully",
			fields: fields{
				Id:                  "",
				CustomerId:          "",
				Customer:            SimpleCustomer{},
				TotalPrice:          0,
				AmountReceived:      0,
				AmountPriceUsePoint: 0,
				CreatedBy:           "",
				CreatedByUser:       usermodel.SimpleUser{},
				CreatedAt:           nil,
				Details:             nil,
			},
			want: common.TableInvoice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoice := &Invoice{
				Id:                  tt.fields.Id,
				CustomerId:          tt.fields.CustomerId,
				Customer:            tt.fields.Customer,
				TotalPrice:          tt.fields.TotalPrice,
				AmountReceived:      tt.fields.AmountReceived,
				AmountPriceUsePoint: tt.fields.AmountPriceUsePoint,
				CreatedBy:           tt.fields.CreatedBy,
				CreatedByUser:       tt.fields.CreatedByUser,
				CreatedAt:           tt.fields.CreatedAt,
				Details:             tt.fields.Details,
			}
			got := invoice.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}
