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
}

const environmentENV = "ENVIRONMENT"

func New() Config {
	cfg := Config{
		Port:         3000,
		DatabasePort: 5432,
	}

	switch os.Getenv(environmentENV) {
	case "development", "":
		loadDevelopmentConfig(&cfg)
	case "test":
		loadTestConfig(&cfg)
	}

	return cfg
}
