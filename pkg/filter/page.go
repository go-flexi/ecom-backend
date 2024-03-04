package filter

// Page represents the pagination filter
type Page struct {
	Skip  int
	Limit int
}

// NewPage creates a new Page
func NewPage(skip, limit int) Page {
	return Page{
		Skip:  skip,
		Limit: limit,
	}
}
