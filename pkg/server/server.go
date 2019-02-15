package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lob/logger-go"
	"github.com/robintsunny/go/pkg/application"
	"github.com/robintsunny/go/pkg/movies"
	"github.com/robintsunny/goyagi/pkg/health"
	"github.com/robintsunny/goyagi/pkg/signals"
)

func New(app application.App) *http.Server {
	log := logger.New()

	e := echo.New()

	health.RegisterRoutes(e)
	movies.RegisterRoutes(e, app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: e,
	}

	graceful := signals.Setup()

	go func() {
		<-graceful
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Err(err).Error("server shutdown")
		}
	}()

	return srv
}
