package mid

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"github.com/go-flexi/ecom-backend/pkg/web/server"
	"go.uber.org/zap"
)

// Authenticate authenticates the request
func Authenticate(core *user.Core, logger *zap.Logger) server.Middleware {
	return func(next server.Handler) server.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				return web.NewTrustedError(fmt.Errorf("authorization token is required"), http.StatusUnauthorized)
			}

			token, err := core.Authenticate(ctx, authToken)
			if errors.Is(err, user.ErrNotFound) {
				return web.NewTrustedError(fmt.Errorf("authorization token is not found"), http.StatusUnauthorized)
			}

			ctx = user.ContextWithToken(ctx, token)
			return next(ctx, w, r)
		}
	}
}
