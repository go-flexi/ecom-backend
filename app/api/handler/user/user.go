package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"github.com/go-flexi/ecom-backend/pkg/web/filter"
	"github.com/google/uuid"
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
	if err := web.JSONDecode(r, &newUser); err != nil {
		return web.NewTrustedError(err, http.StatusBadRequest)
	}

	coreUser, err := newUser.toCoreNewUser()
	if err != nil {
		return web.NewTrustedError(fmt.Errorf("toCoreNewUser: %w", err), http.StatusBadRequest)
	}

	createdUser, err := h.core.Create(ctx, coreUser)
	if err != nil {
		return fmt.Errorf("core.Create: %w", err)
	}

	var response User
	response.set(createdUser)
	return web.JSONResponse(ctx, w, response, http.StatusCreated)
}

func (h *handlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := filter.ParsePage(r)
	if err != nil {
		return fmt.Errorf("parsePage: %w", err)
	}

	order, err := parseOrder(r)
	if err != nil {
		return fmt.Errorf("parseOrder: %w", err)
	}

	filter, err := parseFilter(r)
	if err != nil {
		return fmt.Errorf("parseFilter: %w", err)
	}

	users, err := h.core.Query(ctx, filter, order, page)
	if err != nil {
		return fmt.Errorf("core.Query: %w", err)
	}

	var response []User
	for _, u := range users {
		var user User
		user.set(u)
		response = append(response, user)
	}

	return web.JSONResponse(ctx, w, response, http.StatusOK)
}

func (h *handlers) get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	coreUser, err := h.core.ByID(ctx, id)
	if errors.Is(err, user.ErrNotFound) {
		return web.NewTrustedError(fmt.Errorf("user not found by id %s", id), http.StatusNotFound)
	}
	if err != nil {
		return fmt.Errorf("core.ByID[%s]: %w", id, err)
	}

	var response User
	response.set(coreUser)
	return web.JSONResponse(ctx, w, response, http.StatusOK)
}

func (h *handlers) update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return web.NewTrustedError(fmt.Errorf("invalid user id %s", id), http.StatusBadRequest)
	}

	var updateUser UpdateUser
	if err := web.JSONDecode(r, &updateUser); err != nil {
		return web.NewTrustedError(err, http.StatusBadRequest)
	}

	coreUpdateUser, err := updateUser.toCoreUpdateUser(userID)
	if err != nil {
		return web.NewTrustedError(fmt.Errorf("toCoreUpdateUser: %w", err), http.StatusBadRequest)
	}

	updatedUser, err := h.core.Update(ctx, coreUpdateUser)
	if errors.Is(err, user.ErrNotFound) {
		return web.NewTrustedError(fmt.Errorf("user not found by id %s", id), http.StatusNotFound)
	}
	if err != nil {
		return fmt.Errorf("core.Update: %w", err)
	}

	var response User
	response.set(updatedUser)
	return web.JSONResponse(ctx, w, response, http.StatusOK)
}

func (h *handlers) token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	email, pass, ok := r.BasicAuth()
	if !ok {
		return web.NewTrustedError(errors.New("must provide email and password"), http.StatusUnauthorized)
	}

	emailAddr, err := mail.ParseAddress(email)
	if err != nil {
		return web.NewTrustedError(fmt.Errorf("invalid email %s", email), http.StatusBadRequest)
	}

	userToken, err := h.core.GenerateToken(ctx, *emailAddr, pass)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return web.NewTrustedError(fmt.Errorf("user not found by email %s and password", email), http.StatusNotFound)
		}
		return fmt.Errorf("core.GenerateToken: %w", err)
	}

	var token Token
	token.set(userToken)
	return web.JSONResponse(ctx, w, token, http.StatusOK)
}
