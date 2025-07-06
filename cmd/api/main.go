package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	cachePkg "server/internal/core/cache"
	configPkg "server/internal/core/config"
	"server/internal/core/database"
	ce "server/internal/core/error"
	"server/internal/core/kdb"
	loggerPkg "server/internal/core/logger"
	"server/internal/server"
	"strings"
	"syscall"
)

const (
	DefaultConfigPath = "."
	DefaultAwsRegion  = "us-west-2"
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	loggerPkg.InitializeGlobalLogger()

	// Get config
	config, err := configPkg.NewConfig(DefaultConfigPath)

	// TODO reconfigure logger
	logger := zap.L()

	app := NewApplication(config)
	logger.Info("Created application")

	quit := make(chan os.Signal, 1)
	go func() {
		// Intercept the signals, as before.
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		logger.Info("Waiting for quit signal")

		sig := <-quit
		logger.Info(fmt.Sprintf("Caught signal %v", sig.String()))

		err := app.Terminate()
		if err != nil {
			cancel()
			os.Exit(1)
		}
		cancel()
		os.Exit(0)
	}()

	err = app.Run()
	if err != nil {
		logger.Error("Run application", zap.Error(err),
			zap.String("context", err.GetOp()))
		stopErr := app.Terminate()
		if stopErr != nil {
			logger.Error("Fail to run application", zap.Error(err))
			os.Exit(1)
		}
		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Cleanup program.
	if stopErr := app.Terminate(); stopErr != nil {
		logger.Error("Fail to stop application", zap.Error(err))
		os.Exit(1)
	}
}

// Application represents the program.
type Application struct {
	ConfigPath string
	Config     configPkg.Config

	logger *zap.Logger

	Database    database.Database
	Cache       cachePkg.Cache
	KeyValStore kdb.KeyValStore

	HTTPServer *server.Server
}

func NewApplication(cfg configPkg.Config) *Application {
	return &Application{
		Config:     cfg,
		ConfigPath: DefaultConfigPath,
		logger:     zap.L(),
		HTTPServer: server.NewServer(cfg),
	}
}

func (app *Application) Run() ce.CoreError {
	app.logger.Info("Running application")

	app.Database = database.New(app.Config)
	if err := app.Database.Open(); err != nil {
		return err
	}

	app.Cache = cachePkg.New(app.Config)
	if err := app.Cache.Open(); err != nil {
		return err
	}

	app.KeyValStore = kdb.New(app.Config)
	if err := app.KeyValStore.Open(); err != nil {
		return err
	}

	app.HTTPServer.RDB = app.Database

	// Start the HTTP server.
	if err := app.HTTPServer.Open(); err != nil {
		return err
	}

	return nil
}

// Terminate gracefully stops the program.
func (app *Application) Terminate() error {
	app.logger.Info("Start termination sequence")
	if app.HTTPServer != nil {
		err := app.HTTPServer.Shutdown()
		if err != nil {
			return err
		}
	}

	if app.Database != nil {
		if err := app.Database.Close(); err != nil {
			return err
		}
	}

	if app.KeyValStore != nil {
		if err := app.KeyValStore.Close(); err != nil {
			return err
		}
	}

	app.logger.Info("Finish termination sequence")
	return nil
}

// ParseFlags parses the command line arguments and loads the config.
//
// This exists separately from the Run() function so that we can skip it
// during end-to-end tests. Those tests will configure manually and call Run().
func (app *Application) ParseFlags(args []string) error {
	// Our flag set is very simple. It only includes a config path.
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	fs.StringVar(&app.ConfigPath, "config", DefaultConfigPath, "config path")
	if err := fs.Parse(args); err != nil {
		return err
	}

	// The expand() function is here to automatically expand "~" to the user's
	// home directory. This is a common task as configuration files are typed
	// under the home directory during local development.
	configPath, err := expand(app.ConfigPath)
	if err != nil {
		return err
	}

	// Read our TOML formatted configuration file.
	cfg, err := configPkg.NewConfig(configPath)
	if os.IsNotExist(err) {
		app.logger.Error(fmt.Sprintf("config file not found at %s", app.ConfigPath))
		return err
	} else if err != nil {
		app.logger.Error("Fail to parse flags", zap.Error(err))
		return err
	}
	app.Config = cfg
	app.logger.Info("Parse config")
	return nil
}

// Expand a returned path using tilde expansion. This means that a file path that
// begins with the "~" will be expanded to prefix the user's home directory.
func expand(path string) (string, error) {
	// Ignore if a path has no leading tilde.
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
