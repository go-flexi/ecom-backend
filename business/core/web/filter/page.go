package filter

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/business/core/filter"
	"github.com/go-flexi/ecom-backend/pkg/validate"
)

// ParsePage parses the Page from the request
func ParsePage(r *http.Request) (filter.Page, error) {
	skip := chi.URLParam(r, "skip")
	limit := chi.URLParam(r, "limit")

	var page filter.Page
	var err error

	page.Skip, err = strconv.Atoi(skip)
	if err != nil {
		return filter.Page{}, validate.NewFieldErrors("skip", err)
	}

	page.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return filter.Page{}, validate.NewFieldErrors("limit", err)
	}

	return page, nil
}
