package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	addr := ":8080"

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
