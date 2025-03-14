package token

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"server/internal/db/idb"
	"server/internal/log"
	"time"
)

type service struct {
	log log.Log
	iDB idb.IDB
}

func NewService(log log.Log, iDB idb.IDB) Service {
	return &service{
		log: log,
		iDB: iDB,
	}
}

func (s *service) GenerateToken(ctx context.Context, scope Scope, ttl time.Duration) (*Token, error) {
	value := rand.Text()
	hashed := s.hash(value)

	return &Token{
		Value:     value,
		Hashed:    hashed,
		Encrypted: "",
		Scope:     scope,
		TTL:       ttl,
		Data:      "",
	}, nil
}

func (s *service) StoreToken(ctx context.Context, token *Token) error {
	repo := s.iDB.NewRepo()
	key := fmt.Sprintf("%v::%v", token.Scope, token.Value)
	err := repo.Set(ctx, key, token.Data, token.TTL).Err()
	if err != nil {
		s.log.Error("Store token", "error", err)
		return err
	}
	return nil
}

func (s *service) GetToken(ctx context.Context, scope Scope, tokenValue string) (string, error) {
	repo := s.iDB.NewRepo()
	key := fmt.Sprintf("%v::%v", scope, tokenValue)
	value, err := repo.Get(ctx, key).Result()
	if err != nil {
		s.log.Error("Get token", "error", err)
		return "", err
	}
	return value, nil
}

func (s *service) DeleteToken(ctx context.Context, token *Token) error {
	return nil
}

func (s *service) hash(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
