package user

import (
	"fmt"
	"net/http"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/apperrors"
	"github.com/go-flexi/ecom-backend/pkg/filter"
	webfilter "github.com/go-flexi/ecom-backend/pkg/web/filter"
)

var orderByFields = map[string]struct{}{
	user.OrderByID:        {},
	user.OrderByName:      {},
	user.OrderByCreatedAt: {},
	user.OrderByUpdatedAt: {},
}

func parseOrder(r *http.Request) (filter.OrderBy, error) {
	order, err := webfilter.ParseOrderBy(r, user.DefaultOrderBy)
	if err != nil {
		return filter.OrderBy{}, err
	}

	if _, ok := orderByFields[order.Field]; !ok {
		return filter.OrderBy{}, apperrors.NewFieldErrors("order by", fmt.Errorf("invalid: %s", order.Field))
	}

	return order, nil
}
