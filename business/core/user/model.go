package user

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Password represents a password
type Password []byte

// Hash converts the password to a hash
func (p Password) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
}

// User represents a user of the system
type User struct {
	ID           uuid.UUID
	Name         string
	Email        mail.Address
	PasswordHash []byte
	Roles        Roles
	Enabled      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UpdateUser represents the fields that can be updated
type UpdateUser struct {
	ID       uuid.UUID
	Name     *string
	Roles    []Role
	Password *Password
	Enabled  *bool
}

// NewUser is used to create a new user
type NewUser struct {
	Name     string
	Email    mail.Address
	Password Password
	Roles    Roles
}

// User converts the NewUser to a User
func (nu NewUser) User() (User, error) {
	now := time.Now()
	passwordHash, err := nu.Password.Hash()
	if err != nil {
		return User{}, fmt.Errorf("Password.Hash: %w", err)
	}

	return User{
		ID:           uuid.New(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: passwordHash,
		Roles:        nu.Roles,
		Enabled:      true,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// Token is used to authenticate a user
type Token struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Roles    Roles
	ExpireAt time.Time
}
