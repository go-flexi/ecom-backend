package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"go.uber.org/zap"
)

type handlers struct {
	core   *user.Core
	logger *zap.Logger
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
	if errors.Is(err, user.ErrUniqueEmail) {
		return web.NewTrustedError(fmt.Errorf("email already is existed"), http.StatusBadRequest)
	}

	var response User
	response.fromCoreUser(createdUser)
	return web.JSONResponse(ctx, w, createdUser, http.StatusCreated)
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
