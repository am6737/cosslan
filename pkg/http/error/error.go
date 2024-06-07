package error

import (
	pkghttp "github.com/cossteam/cosslan/pkg/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIError struct {
	Code    int
	Message string
}

var (
	ErrBadRequest     = &APIError{Code: http.StatusBadRequest, Message: "Bad request"}
	ErrUnauthorized   = &APIError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden      = &APIError{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound       = &APIError{Code: http.StatusNotFound, Message: "Not found"}
	ErrInternalServer = &APIError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	// Add more errors as needed
)

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// Handler 用于处理错误并返回统一的响应
func Handler(c echo.Context, err error) error {
	var apiErr *APIError
	if e, ok := err.(*APIError); ok {
		apiErr = e
	} else {
		apiErr = ErrInternalServer
	}

	return pkghttp.Fail(c, apiErr.Code, apiErr.Message)
}
