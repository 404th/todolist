package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		Handler:        handler,
	}

	return nil
}
