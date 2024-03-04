package filter

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
