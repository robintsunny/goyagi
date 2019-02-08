package main


import (
	"net/http"

	"github.com/robintsunny/goyagi/pkg/server"
	"github.com/lob/logger-go"
)

func main() {
	log := logger.New()

	srv := server.New()

	log.Info("server started")

	err := srv.ListenAndServe()

	i err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("server stopped")
	}

	log.Info("server stopped")
}