package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserUpdateInfo_TableName(t *testing.T) {
	type fields struct {
		Name    *string
		Phone   *string
		Address *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserUpdateInfo successfully",
			fields: fields{
				Name:    nil,
				Phone:   nil,
				Address: nil,
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserUpdateInfo{
				Name:    tt.fields.Name,
				Phone:   tt.fields.Phone,
				Address: tt.fields.Address,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserUpdateInfo_Validate(t *testing.T) {
	type fields struct {
		Name    *string
		Phone   *string
		Address *string
	}

	invalidName := ""
	validName := mock.Anything
	invalidPhone := "this is invalid phone"
	validPhone := "0123456789"
	emptyPhone := ""

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserUpdateInfo is invalid with empty name",
			fields: fields{
				Name: &invalidName,
			},
			wantErr: true,
		},
		{
			name: "UserUpdateInfo is invalid with invalid phone",
			fields: fields{
				Name:  nil,
				Phone: &invalidPhone,
			},
			wantErr: true,
		},
		{
			name: "UserUpdateInfo is valid with valid name and empty phone",
			fields: fields{
				Name:  &validName,
				Phone: &emptyPhone,
			},
			wantErr: false,
		},
		{
			name: "UserUpdateInfo is valid with valid name and phone",
			fields: fields{
				Name:  &validName,
				Phone: &validPhone,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserUpdateInfo{
				Name:    tt.fields.Name,
				Phone:   tt.fields.Phone,
				Address: tt.fields.Address,
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
