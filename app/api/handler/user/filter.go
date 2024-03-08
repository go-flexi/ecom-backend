package user

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/apperrors"
)

// parseFilter parses the filter from the request
func parseFilter(r *http.Request) (user.Filter, error) {
	query := r.URL.Query()

	filter := user.Filter{}

	if email := query.Get("email"); email != "" {
		if _, err := mail.ParseAddress(email); err != nil {
			return user.Filter{}, apperrors.NewFieldErrors("email", err)
		}
		filter.WithEmail(email)
	}

	if name := query.Get("name"); name != "" {
		filter.WithName(name)
	}

	if enabled := query.Get("enabled"); enabled != "" {
		if enabled != "true" && enabled != "false" {
			return user.Filter{}, apperrors.NewFieldErrors("enabled", fmt.Errorf("invalied boolean"))
		}

		if enabled == "false" {
			filter.WithEnabled(false)
		} else {
			filter.WithEnabled(true)
		}
	}

	return user.Filter{}, nil
}
