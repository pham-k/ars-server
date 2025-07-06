package server

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/authn"
	tokenPkg "server/internal/token"
	"strings"
)

func (s *Server) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Note: This is split across multiple lines for readability. You don't // need to do this in your own code.
		w.Header().Set(
			"Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set(
			"Referrer-Policy",
			"origin-when-cross-origin")
		w.Header().Set(
			"X-Content-Type-Options",
			"nosniff")
		w.Header().Set(
			"X-Frame-Options",
			"deny")
		w.Header().Set(
			"X-XSS-Protection",
			"0")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			s.Log.Error("Invalid authorization header")
			s.setContextUser(r, authn.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
			s.Log.Error("Invalid authorization header")
			next.ServeHTTP(w, r)
			return
		}

		user, err := s.getUserFromToken(headerParts[1])
		if err != nil {
			s.Log.Error("Get user from token", "error", err)
		}

		// Set user to context
		r = s.setContextUser(r, user)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) getUserFromToken(token string) (*authn.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	value, err := s.TokenService.GetToken(ctx, tokenPkg.Authentication, token)
	if err != nil {
		return nil, err
	}

	// Parse token value
	type tokenData struct {
		ID     int64  `json:"id"`
		PID    string `json:"pid"`
		Object string `json:"object"`
	}
	data := &tokenData{}
	err = json.Unmarshal([]byte(value), data)
	if err != nil {
		return nil, err
	}

	user := &authn.User{
		ID:     data.ID,
		PID:    data.PID,
		Object: data.Object,
	}

	return user, nil
}

// ReportPanic is middleware for catching panics and reporting them.
func ReportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// TODO pham-k: report panic
			}
		}()

		next.ServeHTTP(w, r)
	})
}
