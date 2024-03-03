package user

import "errors"

// ErrInvalidRole is returned when an invalid role is used
var ErrInvalidRole = errors.New("invalid role")

// list of role
const (
	admin = "admin"
	user  = "user"
)

// set of knwon roles
var roles = map[string]Role{
	admin: RoleAdmin(),
	user:  RoleUser(),
}

// Role represents a user role
type Role struct {
	name string
}

// Name returns the role name
func (r Role) Name() string { return r.name }

func (r Role) Equal(r2 Role) bool { return r.name == r2.name }

// RoleAdmin returns the admin role
func RoleAdmin() Role { return Role{name: admin} }

// RoleUser returns the user role
func RoleUser() Role { return Role{name: user} }

// Parse returns a role from a string
func Parse(value string) (Role, error) {
	role, ok := roles[value]
	if !ok {
		return Role{}, ErrInvalidRole
	}
	return role, nil
}
