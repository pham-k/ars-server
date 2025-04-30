package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"server/internal/authn"
	"server/internal/config"
	"server/internal/db/idb"
	"server/internal/db/kdb"
	"server/internal/db/rdb"
	logPkg "server/internal/log"
	"server/internal/server"
	"server/internal/token"
	"strings"
	"syscall"
)

const (
	DefaultConfigPath = "."
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	log := logPkg.NewLog()

	// Get config
	cfg, err := config.NewConfig(DefaultConfigPath, log)

	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion("us-west-2"),
	)
	if err != nil {
		log.Error("Load AWS config", "error", err)
	}

	// Represent the application
	m := NewApp(cfg, awsCfg, log)
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
		log.Error("Fail to run application", "error", err)
		stopErr := m.Terminate()
		if stopErr != nil {
			log.Error("Fail to run application", "error", err)
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
		log.Error("Fail to stop application", "error", stopErr)
		os.Exit(1)
	}
}

// App represents the program.
type App struct {
	// Configuration path and parsed config api.
	Config     config.Config
	ConfigPath string
	AwsConfig  aws.Config

	Log logPkg.Log

	// Relational storage
	RDB rdb.RDB
	// In-memory  storage
	IDB idb.IDB
	// Key-value storage
	KDB kdb.KDB

	// HTTP server for handling HTTP communication.
	// Postgres' services are attached to it before running.
	HTTPServer *server.Server
}

// NewApp returns a new instance of App.
func NewApp(cfg config.Config, awsCfg aws.Config, log logPkg.Log) *App {
	var rDB = rdb.NewDB(cfg, log)
	var iDB = idb.NewIDB(cfg, log)
	var kDB = kdb.NewKDB(cfg, log)

	return &App{
		Config:     cfg,
		ConfigPath: DefaultConfigPath,
		AwsConfig:  awsCfg,
		Log:        log,
		RDB:        rDB,
		IDB:        iDB,
		KDB:        kDB,
		HTTPServer: server.NewServer(cfg),
	}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (a *App) Run(ctx context.Context) error {
	a.Log.Info("Running application")
	if err := a.RDB.Open(); err != nil {
		return err
	}
	if err := a.IDB.Open(); err != nil {
		return err
	}
	if err := a.KDB.Open(); err != nil {
		return err
	}

	// Instantiate SQLite-backed services.
	authnService := authn.NewService(a.Log, a.RDB)
	tokenService := token.NewService(a.Log, a.IDB)

	// Attach underlying services to the HTTP server.
	a.HTTPServer.Log = a.Log
	a.HTTPServer.AuthnService = authnService
	a.HTTPServer.TokenService = tokenService

	// Start the HTTP server.
	if err := a.HTTPServer.Open(); err != nil {
		return err
	}

	return nil
}

// Terminate gracefully stops the program.
func (a *App) Terminate() error {
	a.Log.Info("Start termination sequence")
	if a.HTTPServer != nil {
		err := a.HTTPServer.Shutdown()
		if err != nil {
			return err
		}
	}

	if a.RDB != nil {
		if err := a.RDB.Close(); err != nil {
			return err
		}
	}

	if a.KDB != nil {
		if err := a.KDB.Close(); err != nil {
			return err
		}
	}

	a.Log.Info("Finish termination sequence")
	return nil
}

// ParseFlags parses the command line arguments & loads the config.
//
// This exists separately from the Run() function so that we can skip it
// during end-to-end tests. Those tests will configure manually and call Run().
func (a *App) ParseFlags(args []string) error {
	// Our flag set is very simple. It only includes a config path.
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	fs.StringVar(&a.ConfigPath, "config", DefaultConfigPath, "config path")
	if err := fs.Parse(args); err != nil {
		return err
	}

	// The expand() function is here to automatically expand "~" to the user's
	// home directory. This is a common task as configuration files are typing
	// under the home directory during local development.
	configPath, err := expand(a.ConfigPath)
	if err != nil {
		return err
	}

	// Read our TOML formatted configuration file.
	cfg, err := config.NewConfig(configPath, a.Log)
	if os.IsNotExist(err) {
		a.Log.Error(fmt.Sprintf("Config file not found at %s", a.ConfigPath))
		return err
	} else if err != nil {
		a.Log.Error("Fail to parse flags", err)
		return err
	}
	a.Config = cfg
	a.Log.Info("Parse config")
	return nil
}

// expand returns path using tilde expansion. This means that a file path that
// begins with the "~" will be expanded to prefix the user's home directory.
func expand(path string) (string, error) {
	// Ignore if path has no leading tilde.
	if path != "~" && !strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		return path, nil
	}

	// Fetch the current user to determine the home path.
	u, err := user.Current()
	if err != nil {
		return path, err
	} else if u.HomeDir == "" {
		return path, fmt.Errorf("home directory unset")
	}

	if path == "~" {
		return u.HomeDir, nil
	}
	return filepath.Join(u.HomeDir, strings.TrimPrefix(path, "~"+string(os.PathSeparator))), nil
}
