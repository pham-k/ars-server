package server

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"path"
	"server/internal/core/config"
	"server/internal/core/database"
	ce "server/internal/core/error"
	"server/internal/user/handler"
	"strings"
	"time"
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 5 * time.Second

// Server represents an HTTP server. It is meant to wrap all HTTP functionality
// used by the application so that dependent packages (such as cmd/wtfd) do not
// need to reference the "net/server" package at all.
type Server struct {
	server *http.Server

	// Bind address & root for the server's listener.
	// If root is specified, server is run on TLS using acme/autocert.
	Address string
	Domain  string

	// Services used by the various HTTP routes.
	logger *zap.Logger
	RDB    database.Database
	//AuthnService authn.Service
	//TokenService token.Service
}

// NewServer returns a new instance of Server.
func NewServer(cfg config.Config) *Server {
	address := fmt.Sprintf("%v:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	s := &Server{
		server: &http.Server{
			Addr:         address,
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		Address: address,
		logger:  zap.L(),
	}

	return s
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() ce.CoreError {
	router := s.NewRouter()
	userHandler := handler.NewHandler(router, s.RDB)
	userHandler.RegisterRoutes()

	s.server.Handler = router

	s.logger.Info(fmt.Sprintf("Starting HTTP server on %s", s.Address))
	onListenAndServeErrChan := make(chan error, 1)
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			onListenAndServeErrChan <- err
		}
	}()
	s.logger.Info(fmt.Sprintf("HTTP server is listening on %v", s.Address))

	listenAndServeErr := <-onListenAndServeErrChan
	if listenAndServeErr != nil {
		err := ce.New("Listen and serve", ce.WithErr(listenAndServeErr))
		err.AddOp("Run HTTP server")
		return err
	}

	return nil
}

// Perform some middleware-like tasks that cannot be performed by actual middleware.
// This includes changing route paths for JSON endpoints & overriding methods.
func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// Override method for forms passing "_method" value.
	if r.Method == http.MethodPost {
		switch v := r.PostFormValue("_method"); v {
		case http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete:
			r.Method = v
		}
	}

	// Override content-type for certain extensions.
	// This allows us to easily cURL API endpoints with a ".json" or ".csv"
	// extension instead of having to explicitly set Content-type & Accept headers.
	// The extensions are removed so they don't appear in the routes.
	switch ext := path.Ext(r.URL.Path); ext {
	case ".json":
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Content-type", "application/json")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	case ".csv":
		r.Header.Set("Accept", "text/csv")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	}

	// Delegate remaining HTTP handling to the gorilla router.
	s.server.Handler.ServeHTTP(w, r)
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown() error {
	// Create a context with a 5-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	// Call Shutdown() on our server, passing in the context we just made.
	// Shutdown() will return nil if the graceful shutdown was successful, or an
	// error (which may happen because of a problem closing the listeners, or
	// because the shutdown didn't complete before the 5-second context deadline is
	// hit). We relay this return value to the shutdownErr channel.
	err := s.server.Shutdown(ctx)
	if err != nil {
		s.logger.Error("HTTP server shutdown", zap.Error(err))
		return err
	}

	s.logger.Info("HTTP server shutdown")
	return nil
}
