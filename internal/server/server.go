package server

import (
	"net/http"

	_ "github.com/edwintrumpet/prueba-tecnica-t-evolvers/docs"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Start configures the server, registers the paths and runs it.
func Start(conf config.Config, customersCtr customers.Controller) {
	e := echo.New()

	routes(e, customersCtr)

	e.Logger.Fatal(e.Start(conf.GetPort()))
}

func routes(e *echo.Echo, customersCtr customers.Controller) {
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/customers", customersCtr.List)
}
