package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func New(handler http.Handler, addr string) *Server {
	return &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
