package api

import (
	"net/http"

	"github.com/ansel1/merry"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"github.com/labstack/echo/v4"
)

// CreateCustomers creates a new customer
// @Summary Create
// @Description Creates a new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param data body models.CreateCustomer true "initial data to create a customer"
// @Success 201 {object} models.Customer
// @Router /customers [post]
func (s *server) CreateCustomers(c echo.Context) error {
	var req models.CreateCustomer

	if err := c.Bind(&req); err != nil {
		return merry.Wrap(err)
	}

	customer, err := s.customers.Create(req)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, customer)
}
