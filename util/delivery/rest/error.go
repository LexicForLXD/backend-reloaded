package rest

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/lexicforlxd/backend-reloaded/lexicError"
	"github.com/pkg/errors"
)

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func NewErrorResponse(err error) render.Renderer {

	switch errT := errors.Cause(err).(type) {

	case *lexicError.LexicError:
		return &ErrResponse{
			Err:            errT,
			HTTPStatusCode: errT.HTTPStatusCode,
			StatusText:     errT.StatusText,
			AppCode:        errT.AppCode,
			ErrorText:      errors.Cause(errT).Error(),
		}

	default:
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: 500,
			StatusText:     err.Error(),
			ErrorText:      errors.Cause(err).Error(),
		}
	}

}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
