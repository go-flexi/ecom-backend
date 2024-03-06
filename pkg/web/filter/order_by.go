package filter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/pkg/apperrors"
	"github.com/go-flexi/ecom-backend/pkg/filter"
)

// ParseOrderBy parses the OrderBy from the request.
func ParseOrderBy(r *http.Request, defaultOrder filter.OrderBy) (filter.OrderBy, error) {
	v := chi.URLParam(r, "orderBy")
	if v == "" {
		return defaultOrder, nil
	}

	orderParts := strings.Split(v, ",")

	var by filter.OrderBy

	switch len(orderParts) {
	case 1:
		by = filter.NewOrderBy(orderParts[0], filter.ASC)
	case 2:
		by = filter.NewOrderBy(orderParts[0], orderParts[1])
	default:
		return filter.OrderBy{}, apperrors.NewFieldErrors("order by", fmt.Errorf("invalid: %s", v))
	}

	return by, nil
}
