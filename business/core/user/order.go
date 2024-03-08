package user

import "github.com/go-flexi/ecom-backend/pkg/filter"

// DefaultOrderBy is the default order by
var DefaultOrderBy = filter.NewOrderBy(OrderByID, filter.ASC)

// list of order by
const (
	OrderByID        = "id"
	OrderByName      = "name"
	OrderByCreatedAt = "created_at"
	OrderByUpdatedAt = "updated_at"
)
