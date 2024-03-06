package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-flexi/ecom-backend/pkg/web"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Handler is a function that handles an HTTP request
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// Middleware is a function that takes a handler and returns a new handler
type Middleware func(Handler) Handler

// Server is the web server
type Server struct {
	logger *zap.Logger
	port   string
	router *chi.Mux
	server *http.Server
}

// NewServer returns a new Server
func NewServer(port string, logger *zap.Logger) *Server {
	return &Server{
		router: chi.NewRouter(),
		port:   port,
		logger: logger,
	}
}

// AddRouter adds a new route to the server
func (s *Server) AddRouter(route Route) {
	handler := route.Handler
	for i := len(route.Middlewares) - 1; i >= 0; i-- {
		handler = route.Middlewares[i](handler)
	}

	s.router.Method(route.Method, route.Path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := web.Values{
			TraceID: uuid.NewString(),
			Now:     time.Now(),
		}
		ctx := web.SetValues(r.Context(), &v)

		if err := handler(ctx, w, r); err != nil {
			s.logger.Error("error handling request", zap.Error(err))
		}
	}))
}

// Start starts the server
func (s *Server) Start() error {
	server := &http.Server{
		Addr:    s.port,
		Handler: s.router,
	}
	return server.ListenAndServe()
}

// Stop stops the server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
