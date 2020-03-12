package utils

import (
	"log"
	"github.com/kelseyhightower/envconfig"
	flag "github.com/spf13/pflag"
)

const envPrefix = "REST_SERVER"

// Config app configuration
type Config struct {
	Host     string `envconfig:"HOST"`
	Storage  int    `envconfig:"STORAGE"`
	Database string `envconfig:"DB"`
	Logging  int    `envconfig:"LOGGER"`
}

// InitConfig initial config
func InitConfig() *Config {
	cfg := Config{}
	flag.StringVarP(&cfg.Host, "host", "h", "localhost:5000", "application host")
	flag.IntVarP(&cfg.Storage, "storage", "s", 0, "application storage. 0 - inmemory, 1 - redis - not implemented")
	flag.StringVarP(&cfg.Database, "database", "d",
		"postgres://guard:password@localhost/ms_guard?sslmode=disable", "postgres connection string")
	flag.IntVarP(&cfg.Logging, "logger", "l", 1, "application logger. 0 - Disable, 1 - Standart, 2 - Verbose json")
	flag.Parse()

	err := envconfig.Process(envPrefix, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
