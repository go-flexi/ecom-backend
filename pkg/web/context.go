package web

import (
	"context"
	"time"
)

type ctxKey string

const key ctxKey = "web_context"

// Values represents the values that are stored in the context
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// NewContext returns a new Context that carries values
func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return &Values{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}
	return v
}

// setStatusCode sets the status code in the context
func setStausCode(ctx context.Context, code int) {
	v := GetValues(ctx)
	v.StatusCode = code
}

// setValues sets the values in the context
func setValues(ctx context.Context, v *Values) context.Context {
	return context.WithValue(ctx, key, v)
}
