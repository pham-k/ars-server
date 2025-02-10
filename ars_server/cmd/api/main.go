package main

import (
	"ars_server/internal/config"
	"ars_server/internal/logger"
	"ars_server/internal/postgres"
	"ars_server/internal/repository"
	"ars_server/internal/server"
	"ars_server/internal/service"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	DefaultConfigPath = "."
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	log := logger.NewLogger()

	// Get config
	cfg, err := config.NewConfig(DefaultConfigPath, log)

	// Represent the application
	m := NewMain(cfg, log)
	log.Info("Created application")

	quit := make(chan os.Signal, 1)
	go func() {
		// Intercept the signals, as before.
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		log.Info("Waiting for quit signal")

		sig := <-quit
		log.Info(fmt.Sprintf("Caught signal %v", sig.String()))

		err := m.Terminate()
		if err != nil {
			cancel()
			os.Exit(1)
		}
		cancel()
		os.Exit(0)
	}()

	err = m.Run(ctx)
	if err != nil {
		log.Error("Fail to run application", err)
		stopErr := m.Terminate()
		if stopErr != nil {
			log.Error("Fail to run application", err)
			os.Exit(1)
		}
		os.Exit(1)
	}

	// Wait for CTRL-C.
	log.Info("here 2")
	<-ctx.Done()
	log.Info("here 3")

	// Clean up program.
	if stopErr := m.Terminate(); stopErr != nil {
		log.Error("Fail to stop application", "err", stopErr)
		os.Exit(1)
	}
}

// Main represents the program.
type Main struct {
	// Configuration path and parsed config api.
	Config     config.Config
	ConfigPath string

	Log *slog.Logger

	// Database used by Postgres implementations.
	DB *postgres.DB

	// HTTP server for handling HTTP communication.
	// Postgres' services are attached to it before running.
	HTTPServer *server.Server
}

// NewMain returns a new instance of Main.
func NewMain(cfg config.Config, log *slog.Logger) *Main {
	var db = postgres.NewDB(cfg, log)
	var repo = repository.New(db.DB)
	db.Repo = repo

	return &Main{
		Config:     cfg,
		ConfigPath: DefaultConfigPath,
		Log:        log,
		//
		DB:         db,
		HTTPServer: server.NewServer(cfg),
	}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) error {
	m.Log.Info("Running application")
	if err := m.DB.Open(); err != nil {
		return err
	}

	// Instantiate SQLite-backed services.
	//authService := sqlite.NewAuthService(m.DB)
	configService := service.NewAppConfigurationService(m.Log, m.DB)
	authnService := service.NewAuthnService(m.Log, m.DB)
	//dialMembershipService := sqlite.NewDialMembershipService(m.DB)
	//userService := sqlite.NewUserService(m.DB)

	// Attach user service to Main for testing.
	//m.UserService = userService

	// Set global GA settings.
	//html.MeasurementID = m.Config.GoogleAnalytics.MeasurementID

	// Copy configuration settings to the HTTP server.
	//m.HTTPServer.Addr = m.Config.HTTP.Host
	//m.HTTPServer.Domain = m.Config.HTTP.Domain
	//m.HTTPServer.HashKey = m.Config.HTTP.HashKey
	//m.HTTPServer.BlockKey = m.Config.HTTP.BlockKey
	//m.HTTPServer.GitHubClientID = m.Config.GitHub.ClientID
	//m.HTTPServer.GitHubClientSecret = m.Config.GitHub.ClientSecret

	// Attach underlying services to the HTTP server.
	m.HTTPServer.Log = m.Log
	//m.HTTPServer.AppConfigurationService = authService
	m.HTTPServer.ConfigService = configService
	m.HTTPServer.AuthnService = authnService
	//m.HTTPServer.DialMembershipService = dialMembershipService
	//m.HTTPServer.EventService = eventService
	//m.HTTPServer.UserService = userService

	// Start the HTTP server.
	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	// If TLS enabled, redirect non-TLS connections to TLS.
	//if m.HTTPServer.UseTLS() {
	//	go func() {
	//		log.Fatal(server.ListenAndServeTLSRedirect(m.Config.HTTP.Domain))
	//	}()
	//}

	// Enable internal debug endpoints.
	//go func() { server.ListenAndServeDebug() }()

	return nil
}

// Terminate gracefully stops the program.
func (m *Main) Terminate() error {
	m.Log.Info("Start termination sequence")
	if m.HTTPServer != nil {
		err := m.HTTPServer.Shutdown()
		if err != nil {
			return err
		}
		m.Log.Info("HTTP server shutdown")
	}
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
		m.Log.Info("DB closed")
	}
	m.Log.Info("Finish termination sequence")
	return nil
}

// ParseFlags parses the command line arguments & loads the config.
//
// This exists separately from the Run() function so that we can skip it
// during end-to-end tests. Those tests will configure manually and call Run().
//func (m *Main) ParseFlags(args []string) error {
//	// Our flag set is very simple. It only includes a config path.
//	fs := flag.NewFlagSet("main", flag.ContinueOnError)
//	fs.StringVar(&m.ConfigPath, "config", DefaultConfigPath, "config path")
//	if err := fs.Parse(args); err != nil {
//		return err
//	}
//
//	// The expand() function is here to automatically expand "~" to the user's
//	// home directory. This is a common task as configuration files are typing
//	// under the home directory during local development.
//	configPath, err := expand(m.ConfigPath)
//	if err != nil {
//		return err
//	}
//
//	// Read our TOML formatted configuration file.
//	cfg, err := config.NewConfig(configPath, m.Log)
//	if os.IsNotExist(err) {
//		m.Log.Error(fmt.Sprintf("Config file not found at %s", m.ConfigPath))
//		return err
//	} else if err != nil {
//		m.Log.Error("Fail to parse flags", err)
//		return err
//	}
//	m.Config = cfg
//	m.Log.Info("Parse config")
//	return nil
//}

// expand returns path using tilde expansion. This means that a file path that
// begins with the "~" will be expanded to prefix the user's home directory.
//func expand(path string) (string, error) {
//	// Ignore if path has no leading tilde.
//	if path != "~" && !strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
//		return path, nil
//	}
//
//	// Fetch the current user to determine the home path.
//	u, err := user.Current()
//	if err != nil {
//		return path, err
//	} else if u.HomeDir == "" {
//		return path, fmt.Errorf("home directory unset")
//	}
//
//	if path == "~" {
//		return u.HomeDir, nil
//	}
//	return filepath.Join(u.HomeDir, strings.TrimPrefix(path, "~"+string(os.PathSeparator))), nil
//}
