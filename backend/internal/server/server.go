package server

import (
	"context"
	"employee-management-webapp/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (server *Server) Run() error {
	return server.httpServer.ListenAndServe()
}

func (server *Server) Stop(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
