package main

import (
	"fmt"
	"log/slog"

	"github.com/stevenstr/yaml_config_loader/internal/config"
)

func main() {
	// TODO: init config: simple
	cfg, err := config.LoadConfig("./config/local.yaml")
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(cfg.Env)
	fmt.Println(cfg.Storage)
	fmt.Println(cfg.HTTPServer)
	fmt.Println(cfg.Address)
	fmt.Println()

	// TODO: init config: cleanenv
	cfgnew := config.MustLoad()
	fmt.Println(cfgnew)

}
