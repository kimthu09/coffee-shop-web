package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUpdatePassword_TableName(t *testing.T) {
	type fields struct {
		OldPassword string
		NewPassword string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserUpdatePassword successfully",
			fields: fields{
				OldPassword: "",
				NewPassword: "",
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserUpdatePassword{
				OldPassword: tt.fields.OldPassword,
				NewPassword: tt.fields.NewPassword,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserUpdatePassword_Validate(t *testing.T) {
	type fields struct {
		OldPassword string
		NewPassword string
	}

	invalidPass := "12345"
	validPass := "123456"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserUpdatePassword is invalid with invalid old password",
			fields: fields{
				OldPassword: invalidPass,
			},
			wantErr: true,
		},
		{
			name: "UserUpdatePassword is invalid with invalid new password",
			fields: fields{
				OldPassword: validPass,
				NewPassword: invalidPass,
			},
			wantErr: true,
		},
		{
			name: "UserUpdatePassword is valid",
			fields: fields{
				OldPassword: validPass,
				NewPassword: validPass,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserUpdatePassword{
				OldPassword: tt.fields.OldPassword,
				NewPassword: tt.fields.NewPassword,
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
