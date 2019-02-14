package application

import (
	"github.com/robintsunny/goyagi/pkg/config"
)

type App struct {
	Config config.Config
}

func New() App {
	cfg := config.New()

	return App{cfg}
}
