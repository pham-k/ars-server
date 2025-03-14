package server

import (
	"context"
	"net/http"
	"server/internal/authn"
	"server/internal/helper"
)

func (s *Server) HandleRegisterWithEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		request, err := authn.NewRequestSignUpWithEmail(r)
		if err != nil {
			s.Log.Error("Fail to parse request", "error", err)
			helper.WriteJson(w, http.StatusBadRequest, "ok")
			return
		}

		result, err := authn.RegisterWithEmail(ctx, request, s.Log, s.AuthnService, s.TokenService)
		if err != nil {
			s.Log.Error("Fail to sign up with email", "error", err)
			helper.WriteJson(w, http.StatusInternalServerError, "")
			return
		}
		response := authn.NewResponseSignUpWithEmail(result)
		helper.WriteJson(w, http.StatusOK, response)
	}
}

func (s *Server) HandleLogInWithEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		request, err := authn.NewRequestLogInWithEmail(r)
		if err != nil {
			s.Log.Error("Fail to parse request", "error", err)
			helper.WriteJson(w, http.StatusBadRequest, "ok")
			return
		}

		user, token, err := authn.LogInWithEmail(ctx, request, s.Log, s.AuthnService, s.TokenService)
		if err != nil {
			s.Log.Error("Fail to log in with email", "error", err)
			helper.WriteJson(w, http.StatusInternalServerError, "")
			return
		}
		response := authn.NewResponseLogInWithEmail(user, token)
		helper.WriteJson(w, http.StatusOK, response)
	}
}

func (s *Server) ValidateEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithCancel(context.Background())
		defer cancel()

		token := r.PathValue("token")
		s.Log.Info(token)
	}
}
