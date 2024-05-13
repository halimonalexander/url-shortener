package main

import (
	"golang.org/x/exp/slog"
	"link_shortener/internal/config"
	"link_shortener/internal/storage/sqlite"
	"link_shortener/lib/sl"
	"os"
)

func main() {
	cfg := initConfig()
	log := initLogger(cfg)

	storage := initStorage(cfg, log)
	log.Info("Loading storage is complete")

	//id, err := storage.SaveUrl("test", "debug")
	//if err != nil {
	//	log.Error(err.Error())
	//	os.Exit(1)
	//}
	//log.Info("Created " + strconv.FormatInt(id, 10))
	//err = storage.RemoveUrl("debug")
	//if err != nil {
	//	log.Error(err.Error())
	//	os.Exit(1)
	//}
	_ = storage

	initRouter()

	// TODO run server
	log.Info("Starting server...")
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

func initStorage(cfg *config.Config, log *slog.Logger) *sqlite.Storage {
	var db *sqlite.Storage

	db, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage", sl.ErrorLog(err))
		os.Exit(1)
	}

	return db
}

func initRouter() {

}
