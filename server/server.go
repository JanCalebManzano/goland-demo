package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server is the struct type for the server
type Server struct {
	router *gin.Engine
	app    *http.Server
}

// NewServer is the constructor for the server
func NewServer() *Server {
	r := NewRouter()

	app := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	return &Server{
		router: r,
		app:    app,
	}
}

// ServeHTTP runs the server router in localhost:80
func (s *Server) ServeHTTP() (err error) {
	err = s.app.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Shutdown gracefully shuts down the server without interrupting any active connections
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.app.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
