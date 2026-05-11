package transport

import (
	"net/http"
	"time"

	"github.com/yousefggg/api-gateway/internal/middleware"
	"github.com/yousefggg/api-gateway/internal/router"
)

type Server struct {
	addr   string
	router *router.Router

	jwtSecret string
}

func NewServer(addr string, r *router.Router, jwtSecret string) *Server {
	return &Server{
		addr:      addr,
		router:    r,
		jwtSecret: jwtSecret,
	}
}

func (s *Server) Run() error {

	// 1. базовый router
	handler := s.router.Init()

	// 2. middleware chain
	jwt := middleware.NewJWTMiddleware(s.jwtSecret)

	finalHandler :=
		middleware.Recovery(
			middleware.Logging(
				jwt.Auth(handler),
			),
		)

	// 3. HTTP server
	server := &http.Server{
		Addr:         s.addr,
		Handler:      finalHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}