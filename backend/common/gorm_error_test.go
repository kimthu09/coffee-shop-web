package common

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetGormErr(t *testing.T) {
	type args struct {
		appErr error
	}
	gormErr := GormErr{
		Number:  0,
		Message: "error parameter",
	}

	tests := []struct {
		name string
		args args
		want *GormErr
	}{
		{
			name: "Get GormErr from Error successfully",
			args: args{
				appErr: &gormErr,
			},
			want: &gormErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGormErr(tt.args.appErr)

			assert.Equal(
				t,
				tt.want,
				got,
				"GetGormErr() = %v, want %v", got, tt.want)
		})
	}
}

func TestGormErr_Error(t *testing.T) {
	type fields struct {
		Number  int
		Message string
	}

	errMsg := mock.Anything
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get error message successfully",
			fields: fields{
				Message: errMsg,
			},
			want: errMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gErr := &GormErr{
				Number:  tt.fields.Number,
				Message: tt.fields.Message,
			}

			got := gErr.Error()

			assert.Equal(
				t,
				tt.want,
				got,
				"Error() = %v, want %v", got, tt.want)
		})
	}
}

func TestGormErr_GetDuplicateErrorKey(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		gormEr *GormErr
		args   args
		want   string
	}{
		{
			name: "Get the key has been duplicated successfully",
			gormEr: &GormErr{
				Number:  GormDuplicateErrorNumber,
				Message: "name",
			},
			args: args{
				args: []string{"name", "PRIMARY"},
			},
			want: "name",
		},
		{
			name:   "Get the key has been duplicated failed because the fields is nil",
			gormEr: nil,
			args: args{
				args: []string{"name", "PRIMARY"},
			},
			want: "",
		},
		{
			name: "Get the key has been duplicated failed because the number is not key of GormError duplicate key",
			gormEr: &GormErr{
				Number:  0,
				Message: "name",
			},
			args: args{
				args: []string{"name", "PRIMARY"},
			},
			want: "",
		},
		{
			name: "Get the key has been duplicated failed because the key is not in the list",
			gormEr: &GormErr{
				Number:  GormDuplicateErrorNumber,
				Message: "someKeyHere",
			},
			args: args{
				args: []string{"name", "PRIMARY"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.gormEr.GetDuplicateErrorKey(tt.args.args...)

			assert.Equal(
				t,
				tt.want,
				got,
				"GetDuplicateErrorKey() = %v, want %v", got, tt.want)
		})
	}
}
