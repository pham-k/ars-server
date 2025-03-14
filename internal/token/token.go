package token

import (
	"context"
	"time"
)

type Scope string

const (
	Authentication  Scope = "authentication"
	EmailValidation Scope = "email_validation"
)

type Token struct {
	Value     string
	Hashed    string
	Encrypted string
	Scope     Scope
	TTL       time.Duration
	Data      string
}

type Service interface {
	GenerateToken(ctx context.Context, scope Scope, ttl time.Duration) (*Token, error)
	StoreToken(ctx context.Context, token *Token) error
	GetToken(ctx context.Context, scope Scope, tokenValue string) (string, error)
}
