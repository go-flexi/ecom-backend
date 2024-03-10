package user

import (
	"github.com/go-flexi/ecom-backend/business/core/user"
	"github.com/go-flexi/ecom-backend/business/web/mid"
	"github.com/go-flexi/ecom-backend/pkg/web/server"
	"go.uber.org/zap"
)

// Route returns user API routes
func Route(
	core *user.Core,
	logger *zap.Logger,
	middlewareRegistry mid.Middlewares,
) []server.Route {
	h := newHandlers(core, logger)
	var routes []server.Route

	middlwares := []server.Middleware{middlewareRegistry.Logger, middlewareRegistry.Errors, middlewareRegistry.Panic}

	routes = append(routes, server.Route{Method: "POST", Path: "/api/v1/users", Handler: h.create, Middlewares: middlwares})
	routes = append(routes, server.Route{Method: "POST", Path: "/api/v1/users/token", Handler: h.token, Middlewares: middlwares})

	authMiddlewares := append(middlwares, middlewareRegistry.Auth)

	routes = append(routes, server.Route{Method: "GET", Path: "/api/v1/users/{id}", Handler: h.get, Middlewares: authMiddlewares})
	routes = append(routes, server.Route{Method: "PUT", Path: "/api/v1/users/{id}", Handler: h.update, Middlewares: authMiddlewares})
	routes = append(routes, server.Route{Method: "GET", Path: "/api/v1/users", Handler: h.list, Middlewares: authMiddlewares})

	return routes
}
