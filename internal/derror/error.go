package derror

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
)

type serviceError struct {
	message string
	code    codes.Code
	desc    string
}

var _ error = (*serviceError)(nil)

// Service error instances
var (
	Timeout = serviceError{
		message: "time_out",
		code:    codes.ResourceExhausted,
	}
	AccessDenied = serviceError{
		message: "access_denied",
		code:    codes.Unauthenticated,
	}
	NotImplemented = serviceError{
		message: "not_implemented",
		code:    codes.Unimplemented,
	}
	TooManyRequests = serviceError{
		message: "too_many_requests",
		code:    codes.Unavailable,
	}
	BadRequest = serviceError{
		message: "bad_request",
		code:    codes.InvalidArgument,
	}
	InternalServer = serviceError{
		message: "internal_server",
		code:    codes.Internal,
	}
	Unknown = serviceError{
		message: "unknown_error",
		code:    codes.Unknown,
	}

	ProductNotFound = serviceError{
		message: "product_not_found",
		code:    codes.NotFound,
	}
	ThemeNotFound = serviceError{
		message: "theme_not_found",
		code:    codes.NotFound,
	}
	DimensionNotFound = serviceError{
		message: "dimension_not_found",
		code:    codes.NotFound,
	}

	InvalidColor = serviceError{
		message: "invalid_color",
		code:    codes.InvalidArgument,
	}
	InvalidTheme = serviceError{
		message: "invalid_theme",
		code:    codes.InvalidArgument,
	}
	InvalidDimension = serviceError{
		message: "invalid_dimension",
		code:    codes.InvalidArgument,
	}
	InvalidProduct = serviceError{
		message: "invalid_product",
		code:    codes.InvalidArgument,
	}
	InvalidCompany = serviceError{
		message: "invalid_company",
		code:    codes.InvalidArgument,
	}
)

// Create error message formats
var (
	CreateProductRepoErrorFormat = "[ERROR] failed to create product repo, error: %v\n"
	CreateStdLogErrorFormat      = "[ERROR] failed to create std log instance, error: %v\n"
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
		return int(ce.code)
	}
	return int(Unknown.code)
}

func StatusText(err error) string {
	ce := serviceError{}
	if errors.As(err, &ce) {
		return ce.message
	}
	return Unknown.message
}
