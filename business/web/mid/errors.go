package mid

import (
	"context"
	"net/http"

	"github.com/go-flexi/ecom-backend/pkg/apperrors"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"github.com/go-flexi/ecom-backend/pkg/web/server"
	"go.uber.org/zap"
)

func Errors(logger *zap.Logger) server.Middleware {
	return func(next server.Handler) server.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := next(ctx, w, r)
			if err == nil {
				return nil
			}

			msg, status := parseErr(err)
			web.JSONResponse(ctx, w, map[string]string{
				"error": msg,
			}, status)

			return nil
		}
	}
}

func parseErr(err error) (errMsg string, status int) {
	trustedErr, ok := web.ToTrustedError(err)
	if ok {
		return trustedErr.Error(), trustedErr.Status
	}

	appErr, ok := apperrors.ToAuthorization(err)
	if ok {
		return appErr.Error(), http.StatusForbidden
	}

	appErr, ok = apperrors.ToFieldErrors(err)
	if ok {
		return appErr.Error(), http.StatusBadRequest
	}

	return "internal server error", http.StatusInternalServerError
}
