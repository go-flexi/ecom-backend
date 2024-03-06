package user

import (
	"net/mail"
	"time"

	"github.com/go-flexi/ecom-backend/business/core/user"
)

// NewUser represents a new user request
type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func (nu NewUser) coreNewUser() (user.NewUser, error) {
	email, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return user.NewUser{}, err
	}

	roles, err := user.ParseRoles(nu.Roles)
	if err != nil {
		return user.NewUser{}, err
	}

	return user.NewUser{
		Name:     nu.Name,
		Email:    *email,
		Password: []byte(nu.Password),
		Roles:    roles,
	}, nil
}

// User represents a user
type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Roles     user.Roles `json:"roles"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (u *User) fromCoreUser(cu user.User) {
	u.ID = cu.ID.String()
	u.Name = cu.Name
	u.Email = cu.Email.String()
	u.Roles = cu.Roles
	u.Enabled = cu.Enabled
	u.CreatedAt = cu.CreatedAt
	u.UpdatedAt = cu.UpdatedAt
}
