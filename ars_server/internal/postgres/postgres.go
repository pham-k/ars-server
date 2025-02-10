package postgres

import (
	"ars_server/internal/config"
	"ars_server/internal/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log/slog"
	"time"
)

// DB represents the database connection.
type DB struct {
	DB     *sql.DB
	Repo   *repository.Queries
	ctx    context.Context // background context
	cancel func()          // cancel background context

	Log    *slog.Logger
	Config config.Config
	// Datasource name.
	DSN string

	// Returns the current time. Defaults to time.Now().
	// Can be mocked for tests.
	Now func() time.Time
}

// Tx wraps the SQL Tx object to provide a timestamp at the start of the transaction.
type Tx struct {
	*sql.Tx
	db  *DB
	now time.Time
}

// NewDB returns a new instance of DB associated with the given datasource name.
func NewDB(cfg config.Config, log *slog.Logger) *DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	db := &DB{
		DSN:    dsn,
		Now:    time.Now,
		Log:    log,
		Config: cfg,
	}
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

// Open opens the database connection.
func (db *DB) Open() (err error) {
	// Ensure a DSN is set before attempting to open the database.
	if db.DSN == "" {
		return fmt.Errorf("dsn required")
	}

	db.DB, err = sql.Open("pgx", db.DSN)

	if err != nil {
		db.Log.Error("Fail to open DB connection", "error", err)
		return err
	}

	duration, err := time.ParseDuration(db.Config.DB.MaxIdleConnectionLifetime)
	db.DB.SetMaxOpenConns(db.Config.DB.MaxOpenConnections)
	db.DB.SetMaxIdleConns(db.Config.DB.MaxIdleConnections)
	db.DB.SetConnMaxLifetime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.DB.PingContext(ctx)
	if err != nil {
		db.Log.Error("Fail to ping DB", "error", err)
		return err
	}
	db.Log.Info("Opened DB connection")

	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	// Cancel background context.
	db.cancel()

	// Close database.
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}

// BeginTx starts a transaction and returns a wrapper Tx type. This type
// provides a reference to the database and a fixed timestamp at the start of
// the transaction. The timestamp allows us to mock time during tests as well.
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Return wrapper Tx that includes the transaction start time.
	return &Tx{
		Tx:  tx,
		db:  db,
		now: db.Now().UTC().Truncate(time.Second),
	}, nil
}
