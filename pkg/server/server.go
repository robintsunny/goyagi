package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func New() *http.Server {
	e := echo.New()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3000),
		Handler: e,
	}

	return srv
}
