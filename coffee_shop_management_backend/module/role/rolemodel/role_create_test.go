package rolemodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRoleCreate_TableName(t *testing.T) {
	type fields struct {
		Id       string
		Name     string
		Features []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RoleCreate successfully",
			fields: fields{
				Id:       "",
				Name:     "",
				Features: nil,
			},
			want: common.TableRole,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			role := &RoleCreate{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Features: tt.fields.Features,
			}

			got := role.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestRoleCreate_Validate(t *testing.T) {
	type fields struct {
		Id       string
		Name     string
		Features []string
	}
	
	var emptyFeatures []string
	invalidFeatures := []string{
		"1233123123213213",
	}
	validFeatures := []string{
		"012345678901",
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RoleCreate is invalid with empty name",
			fields: fields{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "RoleCreate is invalid with nil features",
			fields: fields{
				Name:     mock.Anything,
				Features: nil,
			},
			wantErr: true,
		},
		{
			name: "RoleCreate is invalid with empty features",
			fields: fields{
				Name:     mock.Anything,
				Features: emptyFeatures,
			},
			wantErr: true,
		},
		{
			name: "RoleCreate is invalid with features has element is invalid id",
			fields: fields{
				Name:     mock.Anything,
				Features: invalidFeatures,
			},
			wantErr: true,
		},
		{
			name: "RoleCreate is valid with features has element is invalid id",
			fields: fields{
				Name:     mock.Anything,
				Features: validFeatures,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &RoleCreate{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Features: tt.fields.Features,
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
