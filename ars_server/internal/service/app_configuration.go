package service

import (
	"ars_server/internal/postgres"
	"ars_server/internal/root"
	"context"
	"database/sql"
	"log/slog"
)

type AppConfigurationService struct {
	log *slog.Logger
	db  *postgres.DB
}

func NewAppConfigurationService(log *slog.Logger, db *postgres.DB) AppConfigurationService {
	return AppConfigurationService{
		log: log,
		db:  db,
	}
}

func (s AppConfigurationService) GetAppConfigurationsByScope(ctx context.Context, scope root.ConfigScope) ([]root.Config, error) {
	tx, err := s.db.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()
	qtx := s.db.Repo.WithTx(tx)

	appConfigurations, err := qtx.ListAppConfigsByScope(
		ctx,
		sql.NullString{
			String: string(scope),
			Valid:  true,
		},
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	var output []root.Config
	for _, row := range appConfigurations {
		output = append(output, root.Config{
			Pid:       row.Pid,
			Object:    "config",
			Scope:     row.Scope.String,
			Name:      row.Name.String,
			Type:      row.Type.String,
			TextValue: row.TextValue,
		})
	}
	return output, nil
}
