package customermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCustomerUpdateInfo_TableName(t *testing.T) {
	type fields struct {
		Name  *string
		Email *string
		Phone *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CustomerUpdateInfo successfully",
			fields: fields{
				Name:  nil,
				Email: nil,
				Phone: nil,
			},
			want: common.TableCustomer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := &CustomerUpdateInfo{
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
				Phone: tt.fields.Phone,
			}
			got := customer.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestCustomerUpdateInfo_Validate(t *testing.T) {
	type fields struct {
		Name  *string
		Email *string
		Phone *string
	}

	validName := mock.Anything
	invalidName := ""
	validEmail := "a@gmail.com"
	invalidEmail := "12312312"
	validPhone := "0123456789"
	invalidPhone := "01234567890123456789"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "CustomerUpdateInfo is invalid with empty name",
			fields: fields{
				Name: &invalidName,
			},
			wantErr: true,
		},
		{
			name: "CustomerUpdateInfo is invalid with invalid email",
			fields: fields{
				Name:  nil,
				Email: &invalidEmail,
			},
			wantErr: true,
		},
		{
			name: "CustomerUpdateInfo is invalid with invalid phone",
			fields: fields{
				Name:  nil,
				Email: nil,
				Phone: &invalidPhone,
			},
			wantErr: true,
		},
		{
			name: "CustomerUpdateInfo is valid",
			fields: fields{
				Name:  &validName,
				Email: &validEmail,
				Phone: &validPhone,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &CustomerUpdateInfo{
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
