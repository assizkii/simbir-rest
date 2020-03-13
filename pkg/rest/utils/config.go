package utils

import (
	"github.com/kelseyhightower/envconfig"
	flag "github.com/spf13/pflag"
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
	flag.StringVarP(&cfg.Host, "host", "h", "localhost:5000", "application host")
	flag.StringVarP(&cfg.Database, "database", "d",
		"postgres://guard:password@localhost/ms_guard?sslmode=disable", "postgres connection string")
	flag.IntVarP(&cfg.Logging, "logger", "l", 1, "application logger. 0 - Disable, 1 - Standart, 2 - Verbose json")
	flag.StringVarP(&cfg.LogFile, "log", "h", "", "logfile")
	flag.Parse()

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
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
}
