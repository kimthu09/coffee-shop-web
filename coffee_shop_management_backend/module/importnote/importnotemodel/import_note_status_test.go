package importnotemodel

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportNoteStatus_MarshalJSON(t *testing.T) {
	mockStatus := Done
	mockJSON := []byte("\"Done\"")

	tests := []struct {
		name             string
		importNoteStatus *ImportNoteStatus
		want             []byte
	}{
		{
			name:             "Marshal nil ImportNoteStatus to JSON successfully",
			importNoteStatus: nil,
			want:             nil,
		},
		{
			name:             "Marshal ImportNoteStatus to JSON successfully",
			importNoteStatus: &mockStatus,
			want:             mockJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.importNoteStatus.MarshalJSON()

			assert.Equal(
				t,
				tt.want,
				got,
				"Value()=%v, want=%v",
				got,
				tt.want)
		})
	}
}

func TestImportNoteStatus_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}

	status := ImportNoteStatus(0)

	tests := []struct {
		name             string
		importNoteStatus ImportNoteStatus
		args             args
		wantErr          bool
	}{
		{
			name:             "Scan successfully ImportNoteStatus",
			importNoteStatus: status,
			args:             args{value: []byte("InProgress")},
			wantErr:          false,
		},
		{
			name:             "Scan failed ImportNoteStatus with reason not exist",
			importNoteStatus: status,
			args:             args{value: []byte("InvalidReason")},
			wantErr:          true,
		},
		{
			name:             "Scan failed ImportNoteStatus with invalid type args",
			importNoteStatus: status,
			args:             args{value: 13123},
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.importNoteStatus.Scan(tt.args.value)
			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"Scan() err = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"Scan() err = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestImportNoteStatus_String(t *testing.T) {
	tests := []struct {
		name             string
		importNoteStatus ImportNoteStatus
		want             string
	}{
		{
			name:             "String model ImportNoteStatus has value 0 to string InProgress ",
			importNoteStatus: InProgress,
			want:             "InProgress",
		},
		{
			name:             "String model ImportNoteStatus has value 1 to string Done ",
			importNoteStatus: Done,
			want:             "Done",
		},
		{
			name:             "String model ImportNoteStatus has value 2 to string Cancel ",
			importNoteStatus: Cancel,
			want:             "Cancel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.importNoteStatus.String()
			assert.Equal(t, tt.want, got, "String() want=?, got=?", tt.want, got)
		})
	}
}

func TestImportNoteStatus_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}

	mockStatus := InProgress

	tests := []struct {
		name             string
		importNoteStatus ImportNoteStatus
		args             args
		wantErr          bool
	}{
		{
			name:             "Unmarshal valid JSON data successfully",
			importNoteStatus: mockStatus,
			args:             args{data: []byte("\"InProgress\"")},
			wantErr:          false,
		},
		{
			name:             "Unmarshal invalid JSON data (unknown reason) with an error",
			importNoteStatus: mockStatus,
			args:             args{data: []byte("\"UnknownStatus\"")},
			wantErr:          true,
		},
		{
			name:             "Unmarshal invalid JSON data (empty string) with an error",
			importNoteStatus: mockStatus,
			args:             args{data: []byte("\"\"")},
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.importNoteStatus.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"UnmarshalJSON(%v) error = %v, wantErr %v",
					tt.args.data,
					err,
					tt.wantErr,
				)
			} else {
				assert.Nil(
					t,
					err,
					"UnmarshalJSON(%v) error = %v, wantErr %v",
					tt.args.data,
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestImportNoteStatus_Value(t *testing.T) {
	mockInProgress := InProgress
	mockDone := Done
	mockCancel := Cancel

	tests := []struct {
		name             string
		importNoteStatus *ImportNoteStatus
		want             driver.Value
		wantErr          bool
	}{
		{
			name:             "Get value in db of nil status successfully",
			importNoteStatus: nil,
			want:             nil,
		},
		{
			name:             "Get value in db of InProgress successfully",
			importNoteStatus: &mockInProgress,
			want:             "InProgress",
		},
		{
			name:             "Get value in db of Done successfully",
			importNoteStatus: &mockDone,
			want:             "Done",
		},
		{
			name:             "Get value in db of Cancel successfully",
			importNoteStatus: &mockCancel,
			want:             "Cancel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.importNoteStatus.Value()

			assert.Equal(t, tt.want, got, "Value()=%v, want=%v", got, tt.want)
		})
	}
}
