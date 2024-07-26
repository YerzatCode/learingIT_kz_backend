package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:":8080"`
	Database string `yaml:"database" env-required:"true"`
}

func InitConfig() *Config {
	config := os.Getenv("CONFIG_PATH")
	cfg := Config{}

	if config == "" {
		log.Fatal("config path not found")
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("config file does not exist:%s", config)
	}

	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		log.Fatal("failed to read config: %w", err)
	}

	return &cfg
}
