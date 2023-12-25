package customermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCustomerCreate_TableName(t *testing.T) {
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
			name: "Get TableName of CustomerCreate successfully",
			fields: fields{
				Id:    nil,
				Name:  "",
				Email: "",
				Phone: "",
			},
			want: common.TableCustomer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := &CustomerCreate{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
				Phone: tt.fields.Phone,
			}

			got := customer.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestCustomerCreate_Validate(t *testing.T) {
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
			name: "CustomerCreate is invalid with invalid id",
			fields: fields{
				Id: &inValidId,
			},
			wantErr: true,
		},
		{
			name: "CustomerCreate is invalid with empty name",
			fields: fields{
				Id:   nil,
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "CustomerCreate is invalid with invalid email",
			fields: fields{
				Id:    nil,
				Name:  mock.Anything,
				Email: "aaaaaaaa",
			},
			wantErr: true,
		},
		{
			name: "CustomerCreate is invalid with invalid phone",
			fields: fields{
				Id:    nil,
				Name:  mock.Anything,
				Email: "a@gmail.com",
				Phone: "01234567890123",
			},
			wantErr: true,
		},
		{
			name: "CustomerCreate is valid",
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
			data := &CustomerCreate{
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
