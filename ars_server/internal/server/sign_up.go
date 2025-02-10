package server

import (
	"ars_server/internal/api/api_sign_up"
	"ars_server/internal/helper"
	"context"
	"net/http"
)

func (s *Server) HandleSignUpWithEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		request, err := api_sign_up.NewRequestSignUpWithEmail(r)
		if err != nil {
			s.Log.Error("Fail to parse request", "error", err)
			helper.WriteJson(w, http.StatusBadRequest, "ok")
			return
		}

		result, err := api_sign_up.SignUpWithEmail(ctx, s.AuthnService, request)
		if err != nil {
			s.Log.Error("Fail to sign up with email", "error", err)
			helper.WriteJson(w, http.StatusInternalServerError, "")
			return
		}
		response := api_sign_up.NewResponseSignUpWithEmail(result)
		helper.WriteJson(w, http.StatusOK, response)
	}
}
