package user

import (
	"errors"
	"fmt"
)

// ErrInvalidRole is returned when an invalid role is used
var ErrInvalidRole = errors.New("invalid role")

// list of role
const (
	admin   = "admin"
	manager = "manager"
	user    = "user"
)

// set of knwon roles
var roles = map[string]Role{
	admin:   RoleAdmin(),
	user:    RoleUser(),
	manager: RoleManager(),
}

// Role represents a user role
type Role struct {
	name string
}

// Name returns the role name
func (r Role) Name() string { return r.name }

// Equal returns true if the roles are equal
func (r Role) Equal(r2 Role) bool { return r.name == r2.name }

// Roles represents a list of roles
type Roles []Role

// Contains returns true if the list of roles contains the role
func (r Roles) Contains(role Role) bool {
	for _, v := range r {
		if v.Equal(role) {
			return true
		}
	}
	return false
}

// RoleAdmin returns the admin role
func RoleAdmin() Role { return Role{name: admin} }

// RoleManager returns the manager role
func RoleManager() Role { return Role{name: manager} }

// RoleUser returns the user role
func RoleUser() Role { return Role{name: user} }

// ParseRole returns a role from a string
func ParseRole(value string) (Role, error) {
	role, ok := roles[value]
	if !ok {
		return Role{}, ErrInvalidRole
	}
	return role, nil
}

// ParseRoles returns a list of roles from a list of strings
func ParseRoles(values []string) ([]Role, error) {
	var roles []Role
	picked := map[string]bool{}

	for _, v := range values {
		if picked[v] {
			continue
		}

		role, err := ParseRole(v)
		if err != nil {
			return nil, fmt.Errorf("parseRole[%s], %w", v, err)
		}

		roles = append(roles, role)
		picked[v] = true
	}

	return roles, nil
}
