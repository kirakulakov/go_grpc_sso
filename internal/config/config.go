package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-required:"true"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

// MustLoad loads configuration from given path.
//
// If path is empty, it panics.
// If configuration file not found, it panics.
// If configuration failed to read, it panics.
func MustLoad() *Config {
	path := fetchConfigPath()

	log.Printf("loading config from %s", path)

	if path == "" {
		log.Panic("config path is empty")
	}

	if _, err := os.Stat(path); err != nil {
		log.Panicf("config file not found: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Panicf("failed to read config: %s", err.Error())
	}

	log.Printf("config loaded: %+v", cfg)

	return &cfg

}

// fetchConfigPath fetches config path from command line or environment variable
// Priority: flag (command line) > environment variable
// Default: empty string
func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res

}
