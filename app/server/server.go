package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/andresbott/Fe26/internal/httpjson"
	"github.com/rs/zerolog"
	"github.com/spf13/afero"

	"io"
	"net"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	logger *zerolog.Logger
}

type Cfg struct {
	Addr   string
	Root   string
	Logger *zerolog.Logger
}

// NewServer creates a new sever instance that can be started individually
func NewServer(cfg Cfg) *Server {
	if cfg.Addr == "" {
		cfg.Addr = ":8080"
	}

	if cfg.Logger == nil {
		l := zerolog.New(io.Discard)
		cfg.Logger = &l
	}
	//
	//fe26Hndl, err := fe26.New(cfg.Logger, cfg.Root)
	//_ = fe26Hndl
	//if err != nil {
	//	panic("unable to crate Server")
	//}

	var AppFs = afero.NewOsFs()
	httpFs := afero.NewHttpFs(AppFs)
	fileserver := httpjson.FileServer(httpFs.Dir("./"))

	mux := http.NewServeMux()
	mux.Handle("/api/v0/fs/", http.StripPrefix("/api/v0/fs/", fileserver))

	return &Server{
		logger: cfg.Logger,
		server: &http.Server{
			Addr:    cfg.Addr,
			Handler: mux,
		},
	}
}

// Start to listen on the configured address
func (srv *Server) Start() error {
	ln, err := net.Listen("tcp", srv.server.Addr)
	if err != nil {
		return err
	}
	srv.logger.Info().Msg(fmt.Sprintf("server started on: %s", srv.server.Addr))

	if err = srv.server.Serve(ln); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop shut down the server cleanly
func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.server.Shutdown(ctx); err != nil {
		srv.logger.Warn().Msg(fmt.Sprintf("server shutdown: %v", err))
	}
	srv.logger.Info().Msg("server stopped")

}
