package user

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/go-flexi/ecom-backend/pkg/apperrors"
	"github.com/go-flexi/ecom-backend/pkg/filter"
	"github.com/google/uuid"
)

// list of errors
var (
	ErrNotFound = errors.New("not found")
)

const tokenDuration = time.Hour * 24

// User provides functionality to store User
type Store interface {
	Create(context.Context, User) error
	Update(ctx context.Context, uu UpdateUser) error
	ByID(context.Context, string) (User, error)
	ByIDs(context.Context, []string) ([]User, error)
	ByEmailNPassword(ctx context.Context, email mail.Address, passwordHash string) (User, error)
	Query(context.Context, Filter, filter.OrderBy, filter.Page) ([]User, error)
}

// TokenStore represents a store for managing tokens
type TokenStore interface {
	Create(ctx context.Context, token Token) error
	Get(ctx context.Context, tokenID string) (Token, error)
}

// Core represents user use case
type Core struct {
	store      Store
	tokenStore TokenStore
}

// NewCore creates a new Core
func NewCore(store Store, tokenStore TokenStore) *Core {
	return &Core{
		store:      store,
		tokenStore: tokenStore,
	}
}

// Create creates a new user
func (c *Core) Create(ctx context.Context, nu NewUser) (User, error) {
	if err := checkCreatePermission(ctx, nu); err != nil {
		return User{}, fmt.Errorf("createPermissionCheck: %w", err)
	}

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
	if errors := checkGetPermission(ctx, userID); errors != nil {
		return User{}, fmt.Errorf("gerPermissionCheck: %w", errors)
	}

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
func (c *Core) Update(ctx context.Context, uu UpdateUser) (User, error) {
	if err := checkUpdatePermission(ctx, uu); err != nil {
		return User{}, fmt.Errorf("updatePermissionCheck: %w", err)
	}

	if err := c.store.Update(ctx, uu); err != nil {
		return User{}, fmt.Errorf("store.Update[%v]: %w", uu, err)
	}

	user, err := c.ByID(ctx, uu.ID.String())
	if err != nil {
		return User{}, fmt.Errorf("Core.ByID[%s]: %w", uu.ID.String(), err)
	}

	return user, nil
}

// Query returns the users based on the filter
func (c *Core) Query(ctx context.Context, filter Filter, orderBy filter.OrderBy, page filter.Page) ([]User, error) {
	if err := checkQueryPermission(ctx); err != nil {
		return nil, fmt.Errorf("checkQueryPermission: %w", err)
	}

	users, err := c.store.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("store.Query: filter[%v]: orderBy[%v]: page[%v]: %w", filter, orderBy, page, err)
	}

	return users, nil
}

// byEmailNPassword returns the user by email and password
func (c *Core) byEmailNPassword(ctx context.Context, email mail.Address, password string) (User, error) {
	user, err := c.store.ByEmailNPassword(ctx, email, password)
	if err != nil {
		return User{}, fmt.Errorf("store.ByEmailNPassword: %w", err)
	}

	return user, nil
}

// GenerateToken generates a token for the user
func (c *Core) GenerateToken(ctx context.Context, email mail.Address, password string) (Token, error) {
	user, err := c.byEmailNPassword(ctx, email, password)
	if err != nil {
		return Token{}, fmt.Errorf("ByEmailNPassword: %w", err)
	}

	token := Token{
		ID:       uuid.New(),
		UserID:   user.ID,
		Roles:    user.Roles,
		ExpireAt: time.Now().Add(tokenDuration),
	}

	if err := c.tokenStore.Create(ctx, token); err != nil {
		return Token{}, fmt.Errorf("tokenStore.Create: %w", err)
	}

	return token, nil
}

// Authenticate authenticates the token
func (c *Core) Authenticate(ctx context.Context, tokenID string) (Token, error) {
	token, err := c.tokenStore.Get(ctx, tokenID)
	if err != nil {
		return Token{}, fmt.Errorf("tokenStore.Get[%s]: %w", tokenID, err)
	}

	return token, nil
}

func checkQueryPermission(ctx context.Context) error {
	token := GetContextToken(ctx)
	if token.Roles.Contains(RoleAdmin()) || token.Roles.Contains(RoleManager()) {
		return nil
	}
	return apperrors.NewAuthorization("user does not have permission to query users")
}

// checkGetPermission checks if the user has permission to get a user
func checkGetPermission(ctx context.Context, userID string) error {
	token := GetContextToken(ctx)
	if token.Roles.Contains(RoleAdmin()) || token.Roles.Contains(RoleManager()) {
		return nil
	}

	if token.UserID.String() == userID {
		return nil
	}

	return apperrors.NewAuthorization("user does not have permission to get other user")
}

// checkCreatePermission checks if the user has permission to create a user
func checkCreatePermission(ctx context.Context, nu NewUser) error {
	token := GetContextToken(ctx)
	if nu.Roles.Contains(RoleAdmin()) || nu.Roles.Contains(RoleManager()) {
		if !token.Roles.Contains(RoleAdmin()) {
			return apperrors.NewAuthorization("admin role is required to create admin/manager user")
		}
	}
	return nil
}

// checkUpdatePermission checks if the user has permission to update a user
func checkUpdatePermission(ctx context.Context, uu UpdateUser) error {
	token := GetContextToken(ctx)
	if err := checkPasswordUpdatePermission(token, uu); err != nil {
		return err
	}
	if err := checkRoleUPdatePermission(token, uu); err != nil {
		return err
	}
	if err := checkEnableUpdatePermission(token, uu); err != nil {
		return err
	}
	return nil
}

func checkEnableUpdatePermission(token Token, uu UpdateUser) error {
	if token.Roles.Contains(RoleAdmin()) {
		return nil
	}
	return apperrors.NewAuthorization("user does not have permission to update enabled")
}

func checkRoleUPdatePermission(token Token, uu UpdateUser) error {
	if token.Roles.Contains(RoleAdmin()) {
		return nil
	}
	return apperrors.NewAuthorization("user does not have permission to update roles")
}

func checkPasswordUpdatePermission(token Token, uu UpdateUser) error {
	if uu.Password == nil {
		return nil
	}
	if token.UserID.String() == uu.ID.String() {
		return nil
	}
	if token.Roles.Contains(RoleAdmin()) {
		return nil
	}
	return apperrors.NewAuthorization("user does not have permission to update other user's password")
}
