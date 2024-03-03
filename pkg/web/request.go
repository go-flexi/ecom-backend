package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Param returns the URL parameter from the request r by key
func Param(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// JSONResponse sends v as a JSON response to the client
func JSONDecode(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
