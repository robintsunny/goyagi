package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lob/logger-go"
	"github.com/robintsunny/goyagi/pkg/health"
	"github.com/robintsunny/goyagi/pkg/signals"
)

func New() *http.Server {
	log := logger.New()

	e := echo.New()

	health.RegisterRoutes(e)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3000),
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
