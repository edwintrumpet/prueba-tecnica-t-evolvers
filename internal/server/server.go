package server

import (
	"net/http"

	_ "github.com/edwintrumpet/prueba-tecnica-t-evolvers/docs"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Start configures the server, registers the paths and runs it.
func Start(customersCtr customers.Controller) {
	e := echo.New()

	routes(e, customersCtr)

	e.Logger.Fatal(e.Start(":3000"))
}

func routes(e *echo.Echo, customersCtr customers.Controller) {
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/customers", customersCtr.List)
}
