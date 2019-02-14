package config

import "os"

type Config struct {
	Environment string
	Port        int
}

const environmentENV = "ENVIRONMENT"

func New() Config {
	cfg := Config{
		Port: 3000,
	}

	switch os.Getenv(environmentENV) {
	case "development", "":
		loadDevelopmentConfig(&cfg)
	case "test":
		loadTestConfig(&cfg)
	}

	return cfg
}
