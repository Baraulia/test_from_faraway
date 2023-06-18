package client_server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(host, port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              host + ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, // 1 Mb
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return fmt.Errorf("server run: %w", s.httpServer.ListenAndServe())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return fmt.Errorf("server shutdown: %w", s.httpServer.Shutdown(ctx))
}
