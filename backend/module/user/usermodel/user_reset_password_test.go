package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserResetPassword_TableName(t *testing.T) {
	type fields struct {
		UserSenderPass string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserResetPassword successfully",
			fields: fields{
				UserSenderPass: "",
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserResetPassword{
				UserSenderPass: tt.fields.UserSenderPass,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserResetPassword_Validate(t *testing.T) {
	type fields struct {
		UserSenderPass string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserResetPassword is invalid with invalid password",
			fields: fields{
				UserSenderPass: "12345",
			},
			wantErr: true,
		},
		{
			name: "UserResetPassword is invalid with valid password",
			fields: fields{
				UserSenderPass: "123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserResetPassword{
				UserSenderPass: tt.fields.UserSenderPass,
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
