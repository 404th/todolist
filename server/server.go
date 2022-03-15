package server

import (
	"context"
	"net/http"

	"github.com/404th/todolist/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(cfg config.ServerConfig) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServerAddr,
		Handler:        cfg.ServerHandler,
		MaxHeaderBytes: cfg.ServerMaxHeaderBytes,
		ReadTimeout:    cfg.ServerReadingTimeout,
		WriteTimeout:   cfg.ServerWritingTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
