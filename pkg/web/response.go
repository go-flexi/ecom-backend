package web

import (
	"context"
	"encoding/json"
	"net/http"
)

// JSONResponse sends v as a JSON response to the client
func JSONResponse(ctx context.Context, w http.ResponseWriter, v interface{}, code int) error {
	SetStausCode(ctx, code)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
