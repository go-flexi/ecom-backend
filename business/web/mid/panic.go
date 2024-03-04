package mid

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-flexi/ecom-backend/pkg/web/server"
)

// Panic recovers from panics and logs the error
func Panic() server.Middleware {
	return func(next server.Handler) server.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("panic[%v] trace[%s]", rec, string(trace))
				}
			}()

			err = next(ctx, w, r)
			return
		}
	}
}
