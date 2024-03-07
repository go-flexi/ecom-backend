package mid

import "github.com/go-flexi/ecom-backend/pkg/web/server"

// Middlewares holds all the middlewares
type Middlewares struct {
	Auth   server.Middleware
	Logger server.Middleware
	Panic  server.Middleware
}
