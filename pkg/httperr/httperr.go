package httperr

import "net/http"

type HttpError interface {
	error
}

type httpErrorImpl struct {
	Err        string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"-"`
}

func (he *httpErrorImpl) Error() string {
	return he.Err
}

var (
	ErrUnauthorized = func(message string, err error) HttpError {
		e := &httpErrorImpl{
			Message:    message,
			StatusCode: http.StatusUnauthorized,
		}
		if err != nil {
			e.Err = err.Error()
		}
		return e
	}
	ErrBadRequest = func(message string, err error) HttpError {
		e := &httpErrorImpl{
			Message:    message,
			StatusCode: http.StatusBadRequest,
		}
		if err != nil {
			e.Err = err.Error()
		}
		return e
	}
	ErrNotFound = func(message string, err error) HttpError {
		e := &httpErrorImpl{
			Message:    message,
			StatusCode: http.StatusNotFound,
		}
		if err != nil {
			e.Err = err.Error()
		}
		return e
	}
	ErrInternalServerError = func(message string, err error) HttpError {
		e := &httpErrorImpl{
			Message:    message,
			StatusCode: http.StatusInternalServerError,
		}
		if err != nil {
			e.Err = err.Error()
		}
		return e
	}
)

// AsHTTP converts any error (including HttpError) into an HTTP status code
// and a JSON-serializable body.
func AsHTTP(err error) (int, any) {
	if err == nil {
		return http.StatusOK, nil
	}

	if he, ok := err.(*httpErrorImpl); ok {
		return he.StatusCode, he
	}

	if he, ok := err.(HttpError); ok {
		// All HttpError implementations in this package are httpErrorImpl,
		// but in case of custom implementations, default to 500.
		if impl, ok := he.(*httpErrorImpl); ok {
			return impl.StatusCode, impl
		}
	}

	// Fallback: wrap arbitrary error as internal server error.
	wrapped := &httpErrorImpl{
		Err:        err.Error(),
		Message:    "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
	return wrapped.StatusCode, wrapped
}
