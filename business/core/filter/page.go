package filter

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/pkg/validate"
)

// Page represents the pagination filter
type Page struct {
	Skip  int
	Limit int
}

// NewPage creates a new Page
func NewPage(skip, limit int) Page {
	return Page{
		Skip:  skip,
		Limit: limit,
	}
}

// ParsePage parses the Page from the request
func ParsePage(r *http.Request) (Page, error) {
	skip := chi.URLParam(r, "skip")
	limit := chi.URLParam(r, "limit")

	var page Page
	var err error

	page.Skip, err = strconv.Atoi(skip)
	if err != nil {
		return Page{}, validate.NewFieldErrors("skip", err)
	}

	page.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return Page{}, validate.NewFieldErrors("limit", err)
	}

	return page, nil
}
