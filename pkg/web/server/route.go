package server

// Route is a struct that represents a route
type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Middlewares []Middleware
}

// NewRoute returns a new Route
func NewRoute(path, method string, handler Handler, middlewares ...Middleware) Route {
	return Route{
		Path:        path,
		Method:      method,
		Handler:     handler,
		Middlewares: middlewares,
	}
}
