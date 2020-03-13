package utils

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

const envPrefix = "REST_SERVER"

// Config app configuration
type Config struct {
	Host     string `envconfig:"HOST"`
	Database string `envconfig:"DB"`
	Logging  int    `envconfig:"LOGGER"`
	LogFile  string `envconfig:"LOGFILE"`
}

// InitConfig initial config
func InitConfig() *Config {
	cfg := Config{}

	err := envconfig.Process(envPrefix, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Logging == 1 {
		setLogFile(cfg)
	}
	return &cfg
}

func setLogFile(cfg Config) {
	file, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.Print(cfg.LogFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
}
