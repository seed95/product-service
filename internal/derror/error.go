package derror

import (
	"errors"
	"fmt"
	"net/http"
)

type serviceError struct {
	message string
	code    int
	desc    string
}

var _ error = (*serviceError)(nil)

//452-499, 512-599 Unassigned
const (
	StatusUnknown = 552 // Unknown error
)

// Service error instances
var (
	Timeout = serviceError{
		message: "time_out",
		code:    http.StatusRequestTimeout,
	}
	AccessDenied = serviceError{
		message: "access_denied",
		code:    http.StatusBadRequest,
	}
	NotImplemented = serviceError{
		message: "not_implemented",
		code:    http.StatusNotImplemented,
	}
	TooManyRequests = serviceError{
		message: "too_many_requests",
		code:    http.StatusTooManyRequests,
	}
	BadRequest = serviceError{
		message: "bad_request",
		code:    http.StatusBadRequest,
	}
	InternalServer = serviceError{
		message: "internal_server",
		code:    http.StatusInternalServerError,
	}
	Unknown = serviceError{
		message: "unknown_error",
		code:    StatusUnknown,
	}

	NilProduct = serviceError{
		message: "nil_product",
		code:    http.StatusBadRequest,
	}
)

// Create error message formats
var (
	CreateProductRepoErrorFormat = "[ERROR] failed to create product repo, error: %v\n"
)

func (se *serviceError) SetDesc(desc string) {
	se.desc = desc
}

func New(se serviceError, desc string) error {
	return serviceError{
		message: se.message,
		code:    se.code,
		desc:    desc,
	}
}

// Error return error message if not empty
func (se serviceError) Error() string {
	if len(se.desc) != 0 {
		return fmt.Sprintf("(StatusCode:%d) %s, desc: %s", se.code, se.message, se.desc)
	}
	return fmt.Sprintf("(StatusCode:%d) %s", se.code, se.message)
}

func StatusCode(err error) int {
	ce := serviceError{}
	if errors.As(err, &ce) {
		return ce.code
	}
	return http.StatusInternalServerError
}

func StatusText(err error) string {
	ce := serviceError{}
	if errors.As(err, &ce) {
		return ce.message
	}
	return Unknown.message
}
