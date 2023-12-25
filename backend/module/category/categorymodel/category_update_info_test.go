package categorymodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryUpdateInfo_TableName(t *testing.T) {
	type fields struct {
		Name        *string
		Description *string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CategoryUpdateInfo successfully",
			fields: fields{
				Name:        nil,
				Description: nil,
			},
			want: common.TableCategory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := &CategoryUpdateInfo{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
			}
			got := category.TableName()

			assert.Equal(
				t,
				tt.want,
				got,
				"TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestCategoryUpdateInfo_Validate(t *testing.T) {
	type fields struct {
		Name        *string
		Description *string
	}

	emptyName := ""
	validName := "categoryName"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "CategoryUpdateInfo is invalid with empty name",
			fields: fields{
				Name: &emptyName,
			},
			wantErr: true,
		},
		{
			name: "CategoryUpdateInfo is valid with nil name",
			fields: fields{
				Name: nil,
			},
			wantErr: false,
		},
		{
			name: "CategoryUpdateInfo is valid",
			fields: fields{
				Name: &validName,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &CategoryUpdateInfo{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
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
