package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"server/internal/helper"
	"time"
)

func (s *Server) NewRouter() chi.Router {
	r := chi.NewRouter()

	// Report panics to external service.
	r.Use(ReportPanic)

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(s.SecureHeaders)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/v1/health-check", s.HandleHealthCheck())

	r.Route("/v1/authn/", func(router chi.Router) {
		router.Post("/register-with-email", s.HandleRegisterWithEmail())
		router.Get("/validate-email/{token}", s.HandleValidateEmail())
		router.Post("/log-in-with-email", s.HandleLogInWithEmail())
		router.With(s.Authenticate).Post("/log-out", s.HandleLogOut())
	})

	return r
}

func (s *Server) HandleHealthCheck() http.HandlerFunc {
	type HealthCheckResponse struct {
		Status string `json:"status"`
	}

	return func(writer http.ResponseWriter, r *http.Request) {
		helper.WriteJson(writer, 200, &HealthCheckResponse{
			Status: "pass",
		})
	}
}
