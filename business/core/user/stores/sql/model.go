package sql

import (
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/google/uuid"
)

// dbUser represents the user for the sql store
type dbUser struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Roles        string    `db:"roles"`
	Enabled      bool      `db:"enabled"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// dbUpdateUser represents the fields that can be updated
type dbUpdateUser struct {
	ID        uuid.UUID
	Name      *string   `db:"name"`
	Roles     string    `db:"roles"`
	Password  *string   `db:"password"`
	Enabled   *bool     `db:"enabled"`
	UpdatedAt time.Time `db:"updated_at"`
}

// toDBUpdateUser converts a user to a dbUser
func toDBUpdateUser(uu user.UpdateUser) (dbUpdateUser, error) {
	dbUser := dbUpdateUser{}

	var roles []string
	for _, role := range uu.Roles {
		roles = append(roles, role.Name())
	}
	dbUser.Roles = strings.Join(roles, ",")

	if uu.Password != nil {
		p := string(*uu.Password)
		dbUser.Password = &p
	}

	return dbUpdateUser{
		ID:        uu.ID,
		Name:      uu.Name,
		Enabled:   uu.Enabled,
		UpdatedAt: time.Now(),
	}, nil

}

// toDBUser converts a user to a dbUser
func toDBUser(user user.User) dbUser {
	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name())
	}

	return dbUser{
		ID:           user.ID.String(),
		Name:         user.Name,
		Email:        user.Email.String(),
		PasswordHash: string(user.PasswordHash),
		Roles:        strings.Join(roles, ","),
		Enabled:      user.Enabled,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func toCoreUser(dbUser dbUser) (user.User, error) {
	var roles []user.Role
	for _, role := range strings.Split(dbUser.Roles, ",") {
		r, err := user.ParseRole(role)
		if err != nil {
			return user.User{}, err
		}
		roles = append(roles, r)
	}

	id, err := uuid.Parse(dbUser.ID)
	if err != nil {
		return user.User{}, fmt.Errorf("parse user id: %w", err)
	}

	email, err := mail.ParseAddress(dbUser.Email)
	if err != nil {
		return user.User{}, fmt.Errorf("parse user email: %w", err)
	}

	return user.User{
		ID:           id,
		Name:         dbUser.Name,
		Email:        *email,
		PasswordHash: []byte(dbUser.PasswordHash),
		Roles:        roles,
		Enabled:      dbUser.Enabled,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
	}, nil
}
