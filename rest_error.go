package restclient

import (
	"net/http"
)

var _ RestError = (*restErrorImpl)(nil)

// RestError holds the http.Response & an error
type RestError interface {
	error
	Response() *http.Response
}

// NewRestError returns a new RestError with the given http.Response & error
func NewRestError(response *http.Response, err error) RestError {
	return &restErrorImpl{
		err:      err,
		response: response,
	}
}

type restErrorImpl struct {
	err      error
	response *http.Response
}

// Error returns the specific error message
func (r *restErrorImpl) Error() string {
	return r.err.Error()
}

// Error returns the specific error message
func (r *restErrorImpl) String() string {
	return r.err.Error()
}

// Response returns the http.Response. May be null depending on what broke during the request
func (r *restErrorImpl) Response() *http.Response {
	return r.response
}
