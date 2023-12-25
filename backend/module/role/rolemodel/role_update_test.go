package rolemodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRoleUpdate_TableName(t *testing.T) {
	type fields struct {
		Name     *string
		Features *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RoleUpdate successfully",
			fields: fields{
				Name:     nil,
				Features: nil,
			},
			want: common.TableRole,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			role := &RoleUpdate{
				Name:     tt.fields.Name,
				Features: tt.fields.Features,
			}
			got := role.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestRoleUpdate_Validate(t *testing.T) {
	type fields struct {
		Name     *string
		Features *[]string
	}

	validName := mock.Anything
	invalidName := ""
	validFeatures := []string{
		"012345678901",
	}

	var emptyFeatures []string
	invalidFeatures := []string{
		"01234567890123456789",
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RoleUpdate is invalid with invalid name",
			fields: fields{
				Name: &invalidName,
			},
			wantErr: true,
		},
		{
			name: "RoleUpdate is invalid with empty features",
			fields: fields{
				Name:     nil,
				Features: &emptyFeatures,
			},
			wantErr: true,
		},
		{
			name: "RoleUpdate is invalid with invalid features",
			fields: fields{
				Name:     nil,
				Features: &invalidFeatures,
			},
			wantErr: true,
		},
		{
			name: "RoleUpdate is valid",
			fields: fields{
				Name:     &validName,
				Features: &validFeatures,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &RoleUpdate{
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
