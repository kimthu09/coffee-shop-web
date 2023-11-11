package suppliermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSupplierCreate_TableName(t *testing.T) {
	type fields struct {
		Id    *string
		Name  string
		Email string
		Phone string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SupplierCreate successfully",
			fields: fields{
				Id:    nil,
				Name:  "",
				Email: "",
				Phone: "",
			},
			want: common.TableSupplier,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			supplier := &SupplierCreate{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
				Phone: tt.fields.Phone,
			}
			got := supplier.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestSupplierCreate_Validate(t *testing.T) {
	type fields struct {
		Id    *string
		Name  string
		Email string
		Phone string
	}

	inValidId := "012345678901234567890"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "SupplierCreate is invalid with invalid id",
			fields: fields{
				Id: &inValidId,
			},
			wantErr: true,
		},
		{
			name: "SupplierCreate is invalid with empty name",
			fields: fields{
				Id:   nil,
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "SupplierCreate is invalid with invalid email",
			fields: fields{
				Id:    nil,
				Name:  mock.Anything,
				Email: "aaaaaaaa",
			},
			wantErr: true,
		},
		{
			name: "SupplierCreate is invalid with invalid phone",
			fields: fields{
				Id:    nil,
				Name:  mock.Anything,
				Email: "a@gmail.com",
				Phone: "01234567890123",
			},
			wantErr: true,
		},
		{
			name: "SupplierCreate is valid",
			fields: fields{
				Id:    nil,
				Name:  mock.Anything,
				Email: "a@gmail.com",
				Phone: "0123456789",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &SupplierCreate{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
				Phone: tt.fields.Phone,
			}

			err := data.Validate()

			if tt.wantErr {
				assert.NotNil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
