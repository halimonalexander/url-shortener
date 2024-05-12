package main

import (
	"golang.org/x/exp/slog"
	"link_shortener/internal/config"
	"os"
)

func main() {
	cfg := initConfig()
	log := initLogger(cfg)

	initStorage()
	initRouter()

	// TODO run server
	log.Info("Starting...")
	log.Debug("==---> debug messages are enabled")
}

func initConfig() *config.Config {
	return config.MustLoad()
}

func initLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch true {
	case cfg.IsProd():
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: cfg.ParseLogLevel()},
			),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: cfg.ParseLogLevel()},
			),
		)
	}

	return log
}

func initStorage() {

}

func initRouter() {

}
