package categorymodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryCreate_TableName(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CategoryCreate successfully",
			fields: fields{
				Name:        "",
				Description: "",
			},
			want: common.TableCategory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := &CategoryCreate{
				Id:          tt.fields.Id,
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

func TestCategoryCreate_Validate(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Description string
	}

	name := "category name"
	description := "category description"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "CategoryUpdateInfo is invalid with empty name",
			fields: fields{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "CategoryUpdateInfo is valid",
			fields: fields{
				Name:        name,
				Description: "",
			},
			wantErr: false,
		},
		{
			name: "CategoryCreate is valid with valid name and not empty description",
			fields: fields{
				Name:        name,
				Description: description,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &CategoryCreate{
				Id:          tt.fields.Id,
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
