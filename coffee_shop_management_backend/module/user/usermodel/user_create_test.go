package usermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserCreate_TableName(t *testing.T) {
	type fields struct {
		Id       string
		Name     string
		Email    string
		Password string
		Salt     string
		RoleId   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of UserCreate successfully",
			fields: fields{
				Id:       "",
				Name:     "",
				Email:    "",
				Password: "",
				Salt:     "",
				RoleId:   "",
			},
			want: common.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserCreate{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
				Salt:     tt.fields.Salt,
				RoleId:   tt.fields.RoleId,
			}

			got := user.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestUserCreate_Validate(t *testing.T) {
	type fields struct {
		Id       string
		Name     string
		Email    string
		Password string
		Salt     string
		RoleId   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "UserCreate is invalid with empty name",
			fields: fields{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "UserCreate is invalid with invalid email",
			fields: fields{
				Name:  mock.Anything,
				Email: "",
			},
			wantErr: true,
		},
		{
			name: "UserCreate is invalid with invalid role id",
			fields: fields{
				Name:  mock.Anything,
				Email: "a@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "UserCreate is valid",
			fields: fields{
				Name:   mock.Anything,
				Email:  "a@gmail.com",
				RoleId: "012345678901",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &UserCreate{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
				Salt:     tt.fields.Salt,
				RoleId:   tt.fields.RoleId,
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
