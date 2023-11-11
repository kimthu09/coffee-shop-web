package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUpdateStatus_TableName(t *testing.T) {
	type fields struct {
		IsActive *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserUpdateStatus successfully",
			fields: fields{
				IsActive: nil,
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserUpdateStatus{
				IsActive: tt.fields.IsActive,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserUpdateStatus_Validate(t *testing.T) {
	type fields struct {
		IsActive *bool
	}

	validStatus := true

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserUpdateStatus is invalid with nil active status",
			fields: fields{
				IsActive: nil,
			},
			wantErr: true,
		},
		{
			name: "UserUpdateStatus is valid",
			fields: fields{
				IsActive: &validStatus,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserUpdateStatus{
				IsActive: tt.fields.IsActive,
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
