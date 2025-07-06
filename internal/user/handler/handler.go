package handler

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"server/internal/core/database"
	"server/internal/user/repository"
	"server/internal/user/service"
)

type Handler interface {
	RegisterRoutes()
	GetUsers(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	SignOut(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	router  chi.Router
	logger  *zap.Logger
	RDB     database.Database
	service service.Service
}

func NewHandler(router chi.Router, rDB database.Database) Handler {
	repo := repository.NewRepository(rDB)
	svc := service.NewService(repo)
	return &handler{
		router:  router,
		logger:  zap.L(),
		RDB:     rDB,
		service: svc,
	}
}

// RegisterRoutes registers all user-related routes
func (h *handler) RegisterRoutes() {
	h.router.Route("/v1/users", func(router chi.Router) {
		router.Get("/", h.GetUsers)
		router.Post("/sign-up", h.SignUp)
		router.Post("/sign-in", h.SignIn)
		router.Get("/sign-out", h.SignOut)
	})
}
