package server

import (
	"ars_server/internal/helper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
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
	r.Use(SecureHeaders)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/v1/health-check", s.HandleHealthCheck())
	r.Get("/v1/panic", s.HandlePanic())

	r.Get("/v1/app-configurations", s.HandleGetAppConfigurations())

	r.Route("/v1/authn/", func(router chi.Router) {
		router.Post("/sign-up-with-email", s.HandleSignUpWithEmail())
		//router.Post("/log-in", authn.LogIn())
		//router.Post("/log-out", LogOut(app))
		//router.Post("/activate", Activate(app))
	})

	return r
}

func (s *Server) HandleHealthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		helper.WriteJson(writer, 200, "ok")
	}
}

func (s *Server) HandlePanic() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		panic("Oh no")
	}
}
