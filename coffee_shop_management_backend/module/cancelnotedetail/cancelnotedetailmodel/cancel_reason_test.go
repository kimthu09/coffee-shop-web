package cancelnotedetailmodel

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelReason_MarshalJSON(t *testing.T) {
	mockReason := Damaged
	mockJSON := []byte("\"Damaged\"")
	tests := []struct {
		name         string
		cancelReason *CancelReason
		want         []byte
	}{
		{
			name:         "Marshal nil CancelReason to JSON successfully",
			cancelReason: nil,
			want:         nil,
		},
		{
			name:         "Marshal CancelReason to JSON successfully",
			cancelReason: &mockReason,
			want:         mockJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.cancelReason.MarshalJSON()

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

func TestCancelReason_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}

	reason := CancelReason(0)

	tests := []struct {
		name         string
		cancelReason CancelReason
		args         args
		wantErr      bool
	}{
		{
			name:         "Scan successfully cancel reason",
			cancelReason: reason,
			args:         args{value: []byte("Damaged")},
			wantErr:      false,
		},
		{
			name:         "Scan failed cancel reason with reason not exist",
			cancelReason: reason,
			args:         args{value: []byte("InvalidReason")},
			wantErr:      true,
		},
		{
			name:         "Scan failed cancel reason with invalid type args",
			cancelReason: reason,
			args:         args{value: 13123},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cancelReason.Scan(tt.args.value)
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

func TestCancelReason_String(t *testing.T) {
	tests := []struct {
		name         string
		cancelReason CancelReason
		want         string
	}{
		{
			name:         "String model CancelReason has value 0 to string Damaged ",
			cancelReason: Damaged,
			want:         "Damaged",
		},
		{
			name:         "String model CancelReason has value 1 to string OutOfDate ",
			cancelReason: OutOfDate,
			want:         "OutOfDate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cancelReason.String()
			assert.Equal(t, tt.want, got, "String() want=?, got=?", tt.want, got)
		})
	}
}

func TestCancelReason_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}

	mockCancelNote := Damaged

	tests := []struct {
		name         string
		cancelReason CancelReason
		args         args
		wantErr      bool
	}{
		{
			name:         "Unmarshal valid JSON data successfully",
			cancelReason: mockCancelNote,
			args:         args{data: []byte("\"Damaged\"")},
			wantErr:      false,
		},
		{
			name:         "Unmarshal invalid JSON data (unknown reason) with an error",
			cancelReason: mockCancelNote,
			args:         args{data: []byte("\"UnknownReason\"")},
			wantErr:      true,
		},
		{
			name:         "Unmarshal invalid JSON data (empty string) with an error",
			cancelReason: mockCancelNote,
			args:         args{data: []byte("\"\"")},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cancelReason.UnmarshalJSON(tt.args.data)
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

func TestCancelReason_Value(t *testing.T) {
	mockDamaged := Damaged
	mockOutOfDate := OutOfDate
	tests := []struct {
		name         string
		cancelReason *CancelReason
		want         driver.Value
	}{
		{
			name:         "Get value in db of nil reason successfully",
			cancelReason: nil,
			want:         nil,
		},
		{
			name:         "Get value in db of Damaged successfully",
			cancelReason: &mockDamaged,
			want:         "Damaged",
		},
		{
			name:         "Get value in db of OutOfDate successfully",
			cancelReason: &mockOutOfDate,
			want:         "OutOfDate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.cancelReason.Value()

			assert.Equal(t, tt.want, got, "Value()=%v, want=%v", got, tt.want)
		})
	}
}
