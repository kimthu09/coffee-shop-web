package shopgeneralmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShopGeneralUpdate_TableName(t *testing.T) {
	type fields struct {
		Name                   *string
		Email                  *string
		Phone                  *string
		Address                *string
		WifiPass               *string
		AccumulatePointPercent *float32
		UsePointPercent        *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ShopGeneralUpdate successfully",
			fields: fields{
				Name:                   nil,
				Email:                  nil,
				Phone:                  nil,
				Address:                nil,
				WifiPass:               nil,
				AccumulatePointPercent: nil,
				UsePointPercent:        nil,
			},
			want: common.TableShopGeneral,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &ShopGeneralUpdate{
				Name:                   tt.fields.Name,
				Email:                  tt.fields.Email,
				Phone:                  tt.fields.Phone,
				Address:                tt.fields.Address,
				WifiPass:               tt.fields.WifiPass,
				AccumulatePointPercent: tt.fields.AccumulatePointPercent,
				UsePointPercent:        tt.fields.UsePointPercent,
			}

			got := sh.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestShopGeneralUpdate_Validate(t *testing.T) {
	type fields struct {
		Name                   *string
		Email                  *string
		Phone                  *string
		Address                *string
		WifiPass               *string
		AccumulatePointPercent *float32
		UsePointPercent        *float32
	}
	invalidEmail := "invalid-email"
	validEmail := "john@gmail.com"
	invalidPhone := "invalid-phone"
	validPhone := "0123456789"
	invalidAccumulatePointPercent := float32(-0.001)
	validAccumulatePointPercent := float32(0.001)
	invalidUsePointPercent := float32(-1)
	validUsePointPercent := float32(1)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ShopGeneralUpdate is invalid with invalid Email",
			fields: fields{
				Email: &invalidEmail,
			},
			wantErr: true,
		},
		{
			name: "ShopGeneralUpdate is invalid with invalid Phone",
			fields: fields{
				Phone: &invalidPhone,
			},
			wantErr: true,
		},
		{
			name: "ShopGeneralUpdate is invalid with negative AccumulatePointPercent",
			fields: fields{
				AccumulatePointPercent: &invalidAccumulatePointPercent,
			},
			wantErr: true,
		},
		{
			name: "ShopGeneralUpdate is invalid with negative UsePointPercent",
			fields: fields{
				UsePointPercent: &invalidUsePointPercent,
			},
			wantErr: true,
		},
		{
			name: "ShopGeneralUpdate is valid",
			fields: fields{
				Email:                  &validEmail,
				Phone:                  &validPhone,
				AccumulatePointPercent: &validAccumulatePointPercent,
				UsePointPercent:        &validUsePointPercent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ShopGeneralUpdate{
				Name:                   tt.fields.Name,
				Email:                  tt.fields.Email,
				Phone:                  tt.fields.Phone,
				Address:                tt.fields.Address,
				WifiPass:               tt.fields.WifiPass,
				AccumulatePointPercent: tt.fields.AccumulatePointPercent,
				UsePointPercent:        tt.fields.UsePointPercent,
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
