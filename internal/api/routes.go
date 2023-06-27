package api

import (
	"net/http"

	_ "github.com/edwintrumpet/prueba-tecnica-t-evolvers/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *server) routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/customers", s.CreateCustomers)
}

/*
TODO
crear orden de servicio
finalizar orden de servicio
listar clientes activos
listar ordenes de servicio (por fecha, cliente)
*/
