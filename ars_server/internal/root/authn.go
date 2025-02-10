package root

import "context"

type AuthnType string

const (
	AuthnTypeEmail AuthnType = "email"
)

type Customer struct {
	Pid    string `json:"pid"`
	Object string `json:"object"`
	Email  string `json:"email"`
}

type AuthnService interface {
	SignUpWithEmail(ctx context.Context, email, password string) (*Customer, error)
}
