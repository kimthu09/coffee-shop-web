package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		err,
		"invalid request",
		err.Error(),
		"ERROR_INVALID_REQUEST",
	)
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(
		err,
		"internal error",
		err.Error(),
		"ERROR_INTERNAL_ERROR",
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("you have no permission"),
		fmt.Sprintf("ERR_NO_PERMISSION"),
	)
}

func ErrIdIsTooLong() *AppError {
	return NewCustomError(
		errIdIsTooLong,
		fmt.Sprintf("maximum length of key is 9"),
		fmt.Sprintf("ERR_ID_IS_TOO_LONG"),
	)
}

func ErrDuplicateKey(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("the key has been existed"),
		fmt.Sprintf("ERR_DUPLICATE_KEY"),
	)
}

func ErrRecordNotFound() *AppError {
	return NewCustomError(
		errRecordNotFound,
		fmt.Sprintf(errRecordNotFound.Error()),
		fmt.Sprintf("ERR_RECORD_NOT_FOUND"),
	)
}

var (
	errRecordNotFound = errors.New("record not found")
	errIdIsTooLong    = errors.New("id is too long")
)
