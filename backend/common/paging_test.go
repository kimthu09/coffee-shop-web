package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaging_Fulfill(t *testing.T) {
	type fields struct {
		Page  int64
		Limit int64
		Total int64
	}
	tests := []struct {
		name      string
		fields    fields
		pageWant  int64
		limitWant int64
	}{
		{
			name: "Full fill page successfully with page = 0 and limit = 0",
			fields: fields{
				Page:  0,
				Limit: 0,
			},
			pageWant:  1,
			limitWant: 10,
		},
		{
			name: "Full fill page successfully with page > 0 and limit > 0",
			fields: fields{
				Page:  1,
				Limit: 1,
			},
			pageWant:  1,
			limitWant: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paging{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
				Total: tt.fields.Total,
			}

			p.Fulfill()

			assert.Equal(
				t,
				tt.pageWant,
				p.Page,
				"Fulfill() page = %v, want %v", p.Page, tt.pageWant)

			assert.Equal(
				t,
				tt.limitWant,
				p.Limit,
				"Fulfill() limit = %v, want %v", p.Limit, tt.limitWant)
		})
	}
}
