package config

import "os"

type Config struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	Environment      string
	Port             int
	SentryDSN        string
	StatsdHost       string
	StatsdPort       int
}

const environmentENV = "ENVIRONMENT"

func New() Config {
	cfg := Config{
		Port:         3000,
		DatabasePort: 5432,
		SentryDSN:    os.Getenv("SENTRY_DSN"),
		StatsdHost:   "127.0.0.1",
		StatsdPort:   8125,
	}

	switch os.Getenv(environmentENV) {
	case "development", "":
		loadDevelopmentConfig(&cfg)
	case "test":
		loadTestConfig(&cfg)
	}

	return cfg
}
