package lexicError

import "github.com/pkg/errors"

type LexicError struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
	Trace      string `json:"trace,omitempty"`
}

func (e *LexicError) Error() string {
	return e.Err.Error()
}

func NewLXDConnectionError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "LXD connection error",    // user-level status message
		AppCode:        504,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewLXDError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "LXD error",               // user-level status message
		AppCode:        503,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewInternalServerError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "internal server error",   // user-level status message
		AppCode:        500,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewFileError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "file error",              // user-level status message
		AppCode:        506,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewDatabaseError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "database",                // user-level status message
		AppCode:        505,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewNotFoundError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     "not found",               // user-level status message
		AppCode:        404,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}

func NewWrongInputError(err error) *LexicError {
	return &LexicError{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Wrong input error",       // user-level status message
		AppCode:        400,                       // application-specific error code
		ErrorText:      errors.Cause(err).Error(), // application-level error message, for debugging
	}
}
