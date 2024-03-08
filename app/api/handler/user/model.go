package user

import (
	"net/mail"
	"time"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/google/uuid"
)

// NewUser represents a new user request
type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func (nu NewUser) toCoreNewUser() (user.NewUser, error) {
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

func (u *User) set(cu user.User) {
	u.ID = cu.ID.String()
	u.Name = cu.Name
	u.Email = cu.Email.String()
	u.Roles = cu.Roles
	u.Enabled = cu.Enabled
	u.CreatedAt = cu.CreatedAt
	u.UpdatedAt = cu.UpdatedAt
}

// UpdateUser represents a user update request
type UpdateUser struct {
	ID       string   `json:"-"`
	Name     *string  `json:"name,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Password *string  `json:"password,omitempty"`
	Enabled  *bool    `json:"enabled,omitempty"`
}

// coreUpdateUser converts the update user request to core update user
func (uu *UpdateUser) toCoreUpdateUser(userID uuid.UUID) (user.UpdateUser, error) {
	var roles user.Roles
	var err error

	if uu.Roles != nil {
		roles, err = user.ParseRoles(uu.Roles)
		return user.UpdateUser{}, err
	}

	var password *user.Password
	if uu.Password != nil {
		convertedPass := user.Password(*uu.Password)
		password = &convertedPass
	}

	return user.UpdateUser{
		ID:       userID,
		Name:     uu.Name,
		Roles:    roles,
		Password: password,
		Enabled:  uu.Enabled,
	}, nil
}

// Token represents a user token
type Token struct {
	Token string `json:"token"`
}

// set sets the token
func (t *Token) set(token user.Token) {
	t.Token = token.ID.String()
}
