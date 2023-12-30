package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	type fields struct {
		StatusCode int
		RootErr    error
		Message    string
		Log        string
		Key        string
	}

	errorStr := mock.Anything
	err := errors.New(errorStr)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get error message from root error successfully",
			fields: fields{
				RootErr: err,
			},
			want: errorStr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &AppError{
				StatusCode: tt.fields.StatusCode,
				RootErr:    tt.fields.RootErr,
				Message:    tt.fields.Message,
				Log:        tt.fields.Log,
				Key:        tt.fields.Key,
			}

			got := e.Error()

			assert.Equal(
				t,
				tt.want,
				got,
				"Error() = %v, want %v", got, tt.want)
		})
	}
}

func TestAppError_RootError(t *testing.T) {
	type fields struct {
		StatusCode int
		RootErr    error
		Message    string
		Log        string
		Key        string
	}

	errString := mock.Anything
	rootErr1 := errors.New(errString)
	rootErr2 := AppError{RootErr: rootErr1}
	err3 := AppError{RootErr: &rootErr2}

	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "Get root error if root error is AppError successfully",
			fields: fields{
				RootErr: &err3,
			},
			want: rootErr1,
		},
		{
			name: "Get root error if root error is Error successfully",
			fields: fields{
				RootErr: &rootErr2,
			},
			want: rootErr1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &AppError{
				StatusCode: tt.fields.StatusCode,
				RootErr:    tt.fields.RootErr,
				Message:    tt.fields.Message,
				Log:        tt.fields.Log,
				Key:        tt.fields.Key,
			}

			got := e.RootError()

			assert.Equal(
				t,
				tt.want,
				got,
				"RootError() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrDB(t *testing.T) {
	type args struct {
		err error
	}

	err := errors.New(mock.Anything)

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError for error database successfully",
			args: args{
				err: err,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    "Đã có lỗi xảy ra với database. Xin hãy thử lại sau.",
				Log:        err.Error(),
				Key:        "DB_ERROR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrDB(tt.args.err)

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrDB() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrDuplicateKey(t *testing.T) {
	type args struct {
		err error
	}

	err := errors.New(mock.Anything)

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError for error duplicate key successfully",
			args: args{
				err: err,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    err.Error(),
				Log:        err.Error(),
				Key:        "ERR_DUPLICATE_KEY",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrDuplicateKey(tt.args.err)

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrDuplicateKey() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrIdIsTooLong(t *testing.T) {
	err := errors.New("id is too long")

	tests := []struct {
		name string
		want *AppError
	}{
		{
			name: "Get AppError for error id is too long successfully",
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    err.Error(),
				Log:        err.Error(),
				Key:        "ERR_ID_IS_TOO_LONG",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrIdIsTooLong(); !reflect.DeepEqual(got, tt.want) {
				got := ErrIdIsTooLong()

				assert.Equal(
					t,
					tt.want,
					got,
					"ErrIdIsTooLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrInternal(t *testing.T) {
	type args struct {
		err error
	}

	err := errors.New(mock.Anything)

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError for error internal successfully",
			args: args{
				err: err,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    "Đã có lỗi xảy ra, xin hãy thử lại sau.",
				Log:        err.Error(),
				Key:        "ERROR_INTERNAL_ERROR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrInternal(tt.args.err)

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrInternal() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrInvalidRequest(t *testing.T) {
	type args struct {
		err error
	}

	err := errors.New(mock.Anything)

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError for error invalid request successfully",
			args: args{
				err: err,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    "Request không hợp lệ.",
				Log:        err.Error(),
				Key:        "ERROR_INVALID_REQUEST",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrInvalidRequest(tt.args.err)

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrInvalidRequest() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrNoPermission(t *testing.T) {
	type args struct {
		err error
	}

	err := errors.New(mock.Anything)

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError for error no permission successfully",
			args: args{
				err: err,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    err,
				Message:    err.Error(),
				Log:        err.Error(),
				Key:        "ERR_NO_PERMISSION",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrNoPermission(tt.args.err)

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrNoPermission() = %v, want %v", got, tt.want)
		})
	}
}

func TestErrRecordNotFound(t *testing.T) {
	errRecordNotFound = errors.New("record not found")

	tests := []struct {
		name string
		want *AppError
	}{
		{
			name: "Get AppError for error record not found successfully",
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    errRecordNotFound,
				Message:    errRecordNotFound.Error(),
				Log:        errRecordNotFound.Error(),
				Key:        "ERR_RECORD_NOT_FOUND",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrRecordNotFound()

			assert.Equal(
				t,
				tt.want,
				got,
				"ErrRecordNotFound() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewCustomError(t *testing.T) {
	type args struct {
		root error
		msg  string
		key  string
	}

	root := errors.New(mock.Anything)
	msg := "msg"
	key := "key"

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError from root, msg and key successfully",
			args: args{
				root: root,
				msg:  msg,
				key:  key,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    root,
				Message:    msg,
				Log:        root.Error(),
				Key:        key,
			},
		},
		{
			name: "Get AppError from nil root, msg and key successfully",
			args: args{
				root: nil,
				msg:  msg,
				key:  key,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    errors.New(msg),
				Message:    msg,
				Log:        msg,
				Key:        key,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCustomError(tt.args.root, tt.args.msg, tt.args.key)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewCustomError() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewErrorResponse(t *testing.T) {
	type args struct {
		root error
		msg  string
		log  string
		key  string
	}

	root := errors.New(mock.Anything)
	msg := "msg"
	key := "key"
	log := "log"

	tests := []struct {
		name string
		args args
		want *AppError
	}{
		{
			name: "Get AppError from root, msg, log and key successfully",
			args: args{
				root: root,
				msg:  msg,
				log:  log,
				key:  key,
			},
			want: &AppError{
				StatusCode: http.StatusBadRequest,
				RootErr:    root,
				Message:    msg,
				Log:        log,
				Key:        key,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewErrorResponse(tt.args.root, tt.args.msg, tt.args.log, tt.args.key)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewErrorResponse() = %v, want %v", got, tt.want)
		})
	}
}
