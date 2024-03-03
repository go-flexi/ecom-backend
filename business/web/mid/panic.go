package mid

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-flexi/ecom-backend/pkg/web"
)

// Panic recovers from panics and logs the error
func Panic() web.Middleware {
	return func(next web.Handler) web.Handler {
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
