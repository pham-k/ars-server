package server

import (
	"context"
	"net/http"
	"server/internal/authn"
)

type contextKey string

const userContextKey contextKey = "user"

func (s *Server) setContextUser(req *http.Request, user *authn.User) *http.Request {
	ctx := context.WithValue(req.Context(), userContextKey, user)
	return req.WithContext(ctx)
}

func (s *Server) getContextUser(req *http.Request) *authn.User {
	user, ok := req.Context().Value(userContextKey).(*authn.User)
	if !ok {
		panic("user not found in context")
	}
	return user
}
