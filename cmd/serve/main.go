package main

import (
	"net/http"

	"github.com/lob/logger-go"
	"github.com/robintsunny/go/pkg/application"
	"github.com/robintsunny/goyagi/pkg/server"
)

func main() {
	log := logger.New()

	app := application.New()

	srv := server.New()

	log.Info("server started", logger.Data{"port": app.Config.Port})

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("server stopped")
	}

	log.Info("server stopped")
}
