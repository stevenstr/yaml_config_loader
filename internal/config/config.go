package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

// simple config loader without safety.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return &cfg, nil
}

func MustLoad() *Config {
	// грузим имя файла из env
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// проверяем есть ли файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exists: %w", configPath)
	}

	// читаем файл
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("can't read config file: %w", configPath)
	}

	return &cfg
}
