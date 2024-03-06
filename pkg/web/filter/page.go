package filter

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/pkg/apperrors"
	"github.com/go-flexi/ecom-backend/pkg/filter"
)

// ParsePage parses the Page from the request
func ParsePage(r *http.Request) (filter.Page, error) {
	skip := chi.URLParam(r, "skip")
	limit := chi.URLParam(r, "limit")

	var page filter.Page
	var err error

	page.Skip, err = strconv.Atoi(skip)
	if err != nil {
		return filter.Page{}, apperrors.NewFieldErrors("skip", err)
	}

	page.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return filter.Page{}, apperrors.NewFieldErrors("limit", err)
	}

	return page, nil
}
