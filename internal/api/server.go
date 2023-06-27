package api

import (
	_ "github.com/edwintrumpet/prueba-tecnica-t-evolvers/docs"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	"github.com/labstack/echo/v4"
)

type server struct {
	customers customers.Service
}

type Server interface {
	Start(conf config.Config)
}

func NewServer(customers customers.Service) Server {
	return &server{
		customers: customers,
	}
}

// Start configures the server, registers the paths and runs it.
func (s *server) Start(conf config.Config) {
	e := echo.New()

	s.routes(e)

	e.Logger.Fatal(e.Start(conf.GetPort()))
}
