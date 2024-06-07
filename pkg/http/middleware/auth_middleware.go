package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/cossteam/cosslan/internal/app/auth"
	pkgerrors "github.com/cossteam/cosslan/pkg/http/error"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	ErrNoAuthHeader      = errors.New("Authorization header is missing")
	ErrInvalidAuthHeader = errors.New("Authorization header is malformed")
	ErrClaimsInvalid     = errors.New("Provided claims do not match expected scopes")
)

func correctError(err *echo.HTTPError) *echo.HTTPError {
	e := err.Internal.Error()
	fmt.Println("e=>", e)
	if strings.Contains(e, "security requirements failed: authorization failed") || strings.Contains(e, "security requirements failed: Authorization header is missing") {
		err.Code = http.StatusUnauthorized
		return err
	}

	//if strings.Contains(e, "value is required but missing") {
	//	message = code.InvalidParameter.Message()
	//}

	return err
}

func HandleOpenAPIError(c echo.Context, err *echo.HTTPError) error {
	err = correctError(err)
	switch err.Code {
	case 400:
		return pkgerrors.Handler(c, pkgerrors.ErrBadRequest)
	case 401:
		return pkgerrors.Handler(c, pkgerrors.ErrUnauthorized)
	case 403:
		return pkgerrors.Handler(c, pkgerrors.ErrForbidden)
	case 404:
		return pkgerrors.Handler(c, pkgerrors.ErrNotFound)
	default:
		return pkgerrors.Handler(c, pkgerrors.ErrInternalServer)
	}

}

func HandleOpenApiAuthentication(ctx context.Context, auth auth.UseCase, input *openapi3filter.AuthenticationInput) error {

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return err
	}

	if jws == "" {
		return ErrNoAuthHeader
	}

	err = auth.Validate(ctx, jws)
	if err != nil {
		return err
	}

	return nil
}

// GetJWSFromRequest extracts a JWS string from an Authorization: Bearer <jws> header
func GetJWSFromRequest(req *http.Request) (string, error) {
	queryToken := req.URL.Query().Get("token")
	if queryToken != "" {
		return queryToken, nil
	}
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}
	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}
