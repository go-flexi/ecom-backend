package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-flexi/ecom-backend/pkg/web"
	"github.com/go-flexi/ecom-backend/pkg/web/server"
	"go.uber.org/zap"
)

// Logger logs the request and response
func Logger(logger *zap.Logger) server.Middleware {
	return func(next server.Handler) server.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			v := web.GetValues(ctx)

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			logger.Info(
				"request started",
				zap.String("method", r.Method),
				zap.String("path", path),
				zap.String("traceID", v.TraceID),
				zap.String("remoteAddr", r.RemoteAddr),
			)

			err = next(ctx, w, r)

			logger.Info(
				"request completed",
				zap.String("method", r.Method),
				zap.String("path", path),
				zap.String("traceID", v.TraceID),
				zap.String("remoteAddr", r.RemoteAddr),
				zap.String("duration", time.Since(v.Now).String()),
			)

			return err
		}
	}
}
