package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUpdateRole_TableName(t *testing.T) {
	type fields struct {
		RoleId string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserUpdateRole successfully",
			fields: fields{
				RoleId: "",
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserUpdateRole{
				RoleId: tt.fields.RoleId,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserUpdateRole_Validate(t *testing.T) {
	type fields struct {
		RoleId string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserUpdateRole is invalid with invalid role id",
			fields: fields{
				RoleId: "",
			},
			wantErr: true,
		},
		{
			name: "UserUpdateRole is valid",
			fields: fields{
				RoleId: "1234567890",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserUpdateRole{
				RoleId: tt.fields.RoleId,
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
