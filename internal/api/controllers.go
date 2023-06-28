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

	return c.JSON(http.StatusCreated, customer)
}

// CreateWorkOrder creates a new work order
// @Summary Create
// @Description Creates a new work order
// @Tags Work orders
// @Accept json
// @Produce json
// @Param data body models.CreateWorkOrder true "initial data to create a work order"
// @Success 201 {object} models.WorkOrder
// @Router /workorders [post]
func (s *server) CreateWorkOrder(c echo.Context) error {
	var req models.CreateWorkOrder

	if err := c.Bind(&req); err != nil {
		return merry.Wrap(err)
	}

	order, err := s.workOrders.Create(req)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusCreated, order)
}
