package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/exp/slog"
	"link_shortener/lib/e"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-default:"prod"`
	LogLevel    string     `yaml:"log_level" env-default:"warn"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server" env-required:"true"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

func (c Config) IsLocal() bool {
	return c.Env == "local"
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func (c Config) ParseLogLevel() slog.Level {
	switch c.LogLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func MustLoad() *Config {
	var configPath string
	if configPath = getConfigPath(); configPath == "" {
		log.Fatal("Config path is not set")
	}

	config, err := readConfig(configPath)
	if err != nil {
		log.Fatal("Unable to read config file: " + err.Error())
	}

	return config
}

func getConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()
	if configPath != "" {
		return configPath
	}

	return os.Getenv("CONFIG_PATH")
}

func readConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); err != nil {
		return nil, e.Wrap("Unable to read config file "+configPath, err)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, e.Wrap("Unable to parse config file:"+configPath, err)
	}

	return &config, nil
}
