package user

// Filter represents a filter for querying users.
type Filter struct {
	Email   *string
	Roles   []string
	Enabled *bool
}

// NewFilter creates a new filter.
func NewFilter() Filter {
	return Filter{}
}

// WithEmail adds an email filter to the filter.
func (f *Filter) WithEmail(email string) *Filter {
	f.Email = &email
	return f
}

// WithRoles adds a role filter to the filter.
func (f *Filter) WithRoles(roles ...string) *Filter {
	f.Roles = roles
	return f
}

// WithEnabled adds an enabled filter to the filter.
func (f *Filter) WithEnabled(enabled bool) *Filter {
	f.Enabled = &enabled
	return f
}
