package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-flexi/ecom-backend/pkg/filter"
)

// list of errors
var (
	ErrNotFound    = errors.New("user not found")
	ErrUniqueEmail = errors.New("email already exists")
)

// User provides functionality to store User
type Store interface {
	Create(context.Context, User) error
	Update(ctx context.Context, userID string, uu UpdateUser) error
	ByID(context.Context, string) (User, error)
	ByIDs(context.Context, []string) ([]User, error)
	Query(context.Context, Filter, filter.OrderBy, filter.Page) ([]User, error)
}

// Core represents user use case
type Core struct {
	store Store
}

// NewCore creates a new Core
func NewCore(store Store) *Core {
	return &Core{store: store}
}

// Create creates a new user
func (c *Core) Create(ctx context.Context, nu NewUser) (User, error) {
	user, err := nu.User()
	if err != nil {
		return user, fmt.Errorf("NewUser.User: %w", err)
	}
	if err := c.store.Create(ctx, user); err != nil {
		return User{}, fmt.Errorf("store.Create: %w", err)
	}

	return user, nil
}

// ByID returns the User by id
func (c *Core) ByID(ctx context.Context, userID string) (User, error) {
	user, err := c.store.ByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("store.ByID[%s]: %w", userID, err)
	}

	return user, nil
}

// ByIDs returns the Users by ids
func (c *Core) ByIDs(ctx context.Context, userIDs []string) ([]User, error) {
	users, err := c.store.ByIDs(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("store.ByIDs[%v]: %w", userIDs, err)
	}

	return users, nil
}

// Update updates the user
func (c *Core) Update(ctx context.Context, userID string, uu UpdateUser) (User, error) {
	if err := c.store.Update(ctx, userID, uu); err != nil {
		return User{}, fmt.Errorf("store.Update[%s, %v]: %w", userID, uu, err)
	}

	user, err := c.ByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("Core.ByID[%s]: %w", userID, err)
	}

	return user, nil
}

// Query returns the users based on the filter
func (c *Core) Query(ctx context.Context, filter Filter, orderBy filter.OrderBy, page filter.Page) ([]User, error) {
	users, err := c.store.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("store.Query: filter[%v]: orderBy[%v]: page[%v]: %w", filter, orderBy, page, err)
	}

	return users, nil
}
