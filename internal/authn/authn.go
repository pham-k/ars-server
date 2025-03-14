package authn

import (
	"context"
)

type Type string

const (
	AuthnTypeEmail  Type   = "email"
	AuthnTypeGoogle Type   = "google"
	ObjUser         string = "user"
	PIDPrefixUser          = "usr"
)

type User struct {
	ID        int64  `json:"id"`
	PID       string `json:"pid"`
	Object    string `json:"object"`
	AuthnType Type   `json:"authn_type"`
	Email     string `json:"email"`
}

type Service interface {
	RegisterWithEmail(ctx context.Context, email, password string) (*User, error)
	LogInWithEmail(ctx context.Context, email, password string) (*User, error)
	LogOut(ctx context.Context, customerPID string) error
}
