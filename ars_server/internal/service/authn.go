package service

import (
	"ars_server/internal/postgres"
	"ars_server/internal/repository"
	"ars_server/internal/root"
	"context"
	"database/sql"
	"github.com/alexedwards/argon2id"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log/slog"
)

type AuthnService struct {
	log *slog.Logger
	db  *postgres.DB
}

func NewAuthnService(log *slog.Logger, db *postgres.DB) AuthnService {
	return AuthnService{
		log: log,
		db:  db,
	}
}

func (s AuthnService) SignUpWithEmail(ctx context.Context, email, password string) (*root.Customer, error) {
	passwordHash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}

	nanoID, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	pid := "cus_" + nanoID

	tx, err := s.db.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()
	qtx := s.db.Repo.WithTx(tx)

	customer, err := qtx.CreateCustomer(ctx, repository.CreateCustomerParams{
		Pid:   pid,
		Email: sql.NullString{String: email, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	authnEmail, err := qtx.CreateAuthnEmail(ctx, repository.CreateAuthnEmailParams{
		Email:        sql.NullString{String: email, Valid: true},
		Passwordhash: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	_, err = qtx.CreateAuthn(ctx, repository.CreateAuthnParams{
		Customerid: sql.NullInt32{Int32: customer.ID, Valid: true},
		Type:       sql.NullString{String: string(root.AuthnTypeEmail), Valid: true},
		Refid:      sql.NullInt32{Int32: authnEmail.ID, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	output := &root.Customer{
		Pid:    customer.Pid,
		Object: "customer",
		Email:  customer.Email.String,
	}
	return output, nil
}
