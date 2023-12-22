package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportNoteUpdate_TableName(t *testing.T) {
	type fields struct {
		ClosedBy   string
		Id         string
		SupplierId string
		TotalPrice int
		Status     *ImportNoteStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ImportNoteUpdate successfully",
			fields: fields{
				ClosedBy:   "",
				Id:         "",
				SupplierId: "",
				TotalPrice: 0,
				Status:     nil,
			},
			want: common.TableImportNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			importNote := &ImportNoteUpdate{
				ClosedBy:   tt.fields.ClosedBy,
				Id:         tt.fields.Id,
				SupplierId: tt.fields.SupplierId,
				TotalPrice: tt.fields.TotalPrice,
				Status:     tt.fields.Status,
			}
			got := importNote.TableName()

			assert.Equal(
				t, tt.want, got,
				"TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestImportNoteUpdate_Validate(t *testing.T) {
	type fields struct {
		ClosedBy   string
		Id         string
		SupplierId string
		TotalPrice int
		Status     *ImportNoteStatus
	}

	inProgress := InProgress
	done := Done
	cancel := Cancel

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ImportNoteUpdate is invalid with InProgress status",
			fields: fields{
				Status: &inProgress,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteUpdate is invalid with nil status",
			fields: fields{
				Status: nil,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteUpdate is valid with Done status",
			fields: fields{
				Status: &done,
			},
			wantErr: false,
		},
		{
			name: "ImportNoteUpdate is valid with Cancel status",
			fields: fields{
				Status: &cancel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ImportNoteUpdate{
				ClosedBy:   tt.fields.ClosedBy,
				Id:         tt.fields.Id,
				SupplierId: tt.fields.SupplierId,
				TotalPrice: tt.fields.TotalPrice,
				Status:     tt.fields.Status,
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
