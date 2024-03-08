package user

// Filter represents a filter for querying users.
type Filter struct {
	Email   *string
	Name    *string
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

// WithName adds a name filter to the filter.
func (f *Filter) WithName(name string) *Filter {
	f.Name = &name
	return f
}

// WithEnabled adds an enabled filter to the filter.
func (f *Filter) WithEnabled(enabled bool) *Filter {
	f.Enabled = &enabled
	return f
}
