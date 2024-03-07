package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"go.uber.org/zap"
)

// handlers is a set of HTTP handlers for user API.
type handlers struct {
	core   *user.Core
	logger *zap.Logger
}

// NewHandlers returns a new handlers.
func newHandlers(core *user.Core, logger *zap.Logger) *handlers {
	return &handlers{
		core:   core,
		logger: logger,
	}
}

func (h *handlers) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var newUser NewUser
	err := web.JSONDecode(r, &newUser)
	if err != nil {
		return web.NewTrustedError(err, http.StatusBadRequest)
	}

	coreUser, err := newUser.coreNewUser()
	if err != nil {
		return web.NewTrustedError(err, http.StatusBadRequest)
	}

	createdUser, err := h.core.Create(ctx, coreUser)
	if err != nil {
		return fmt.Errorf("core.Create: %w", web.AppErrToTrustedErr(err))
	}

	var response User
	response.set(createdUser)
	return web.JSONResponse(ctx, w, response, http.StatusCreated)
}

func (h *handlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handlers) get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handlers) update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handlers) token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
