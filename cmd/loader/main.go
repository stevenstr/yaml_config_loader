package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/stevenstr/yaml_config_loader/internal/config"
	"github.com/stevenstr/yaml_config_loader/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Done: init config: simple
	// cfg, err := config.LoadConfig("./config/local.yaml")
	// if err != nil {
	// 	slog.Error(err.Error())
	// }
	// fmt.Println(cfg.Env)
	// fmt.Println(cfg.Storage)
	// fmt.Println(cfg.HTTPServer)
	// fmt.Println(cfg.Address)
	// fmt.Println()

	// Done: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// Done: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting loader service...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enable!")

	// TODO: storage: sqlite3
	storage, err := sqlite.New(cfg.Storage)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	_ = storage

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		// для прода, json для grafana kibana
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
