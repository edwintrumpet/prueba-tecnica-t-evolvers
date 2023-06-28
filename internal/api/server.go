package api

import (
	"github.com/ansel1/merry"
	_ "github.com/edwintrumpet/prueba-tecnica-t-evolvers/docs"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	workorders "github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/work_orders"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type server struct {
	customers  customers.Service
	workOrders workorders.Service
}

type Server interface {
	Start(conf config.Config)
}

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

func NewServer(
	customers customers.Service,
	workOrders workorders.Service,
) Server {
	return &server{
		customers:  customers,
		workOrders: workOrders,
	}
}

// Start configures the server, registers the paths and runs it.
func (s *server) Start(conf config.Config) {
	e := echo.New()

	if conf.IsDev() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	e.HTTPErrorHandler = errorHandler
	e.Pre(middleware.RemoveTrailingSlash())

	s.routes(e)

	e.Logger.Fatal(e.Start(conf.GetPort()))
}

func errorHandler(err error, c echo.Context) {
	userMessage := merry.UserMessage(err)
	statusCode := merry.HTTPCode(err)

	if userMessage == "" {
		userMessage = merry.Message(err)
	}

	logrus.WithFields(logrus.Fields{
		"stack":       merry.Stacktrace(err),
		"userMessage": userMessage,
		"statusCode":  statusCode,
		"error":       err.Error(),
	}).Debug("api error")

	c.JSON(statusCode, ErrorResponse{Error: userMessage})
}
