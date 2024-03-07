package web

import (
	"net/http"

	"github.com/go-flexi/ecom-backend/pkg/apperrors"
)

// AppErrToTrustedErr converts an app error to a trusted error
// if the error is not an app error, it is returned as is
func AppErrToTrustedErr(err error) error {
	if apperrors.IsAuthorization(err) {
		return NewTrustedError(err, http.StatusForbidden)
	}
	if apperrors.IsFieldErrors(err) {
		return NewTrustedError(err, http.StatusBadRequest)
	}
	return err
}
