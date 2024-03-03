package filter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/pkg/validate"
)

// set of direction for ordering
const (
	ASC  = "ASC"
	DESC = "DESC"
)

// map of direction for ordering
var directions = map[string]string{
	ASC:  "ASC",
	DESC: "DESC",
}

// OrderBy represents order filter
type OrderBy struct {
	Field     string
	Direction string
}

// NewOrderBy creates a new OrderBy.
func NewOrderBy(field, direction string) OrderBy {
	dir, ok := directions[direction]
	if !ok {
		dir = ASC
	}

	return OrderBy{
		Field:     field,
		Direction: dir,
	}
}

// ParseOrderBy parses the OrderBy from the request.
func ParseOrderBy(r *http.Request, defaultOrder OrderBy) (OrderBy, error) {
	v := chi.URLParam(r, "orderBy")
	if v == "" {
		return defaultOrder, nil
	}

	orderParts := strings.Split(v, ",")

	var by OrderBy

	switch len(orderParts) {
	case 1:
		by = NewOrderBy(orderParts[0], ASC)
	case 2:
		by = NewOrderBy(orderParts[0], orderParts[1])
	default:
		return OrderBy{}, validate.NewFieldErrors("order by", fmt.Errorf("invalid: %s", v))
	}

	return by, nil
}
