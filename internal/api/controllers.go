package api

import (
	"net/http"

	"github.com/ansel1/merry"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"github.com/labstack/echo/v4"
)

/* ------------------------------- Customers ------------------------------- */

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

// ListActiveCustomers list customers with is_active in true
// @Summary List active customers
// @Description List customers with is_active in true
// @Tags Customers
// @Produce json
// @Success 200 {object} []models.Customer
// @Router /customers [get]
func (s *server) ListActiveCustomers(c echo.Context) error {
	customers, err := s.customers.ListActive()
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, customers)
}

/* ------------------------------ Work orders ------------------------------ */

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

// FinishWorkOrder finishes a work order
// @Summary Finish
// @Description Finishes a work order
// @Tags Work orders
// @Accept json
// @Produce json
// @Param data body models.FinishWorkOrder true "ids to find the work order to finish"
// @Success 200 {object} models.WorkOrder
// @Router /workorders/finish [post]
func (s *server) FinishWorkOrder(c echo.Context) error {
	var req models.FinishWorkOrder

	if err := c.Bind(&req); err != nil {
		return merry.Wrap(err)
	}

	order, err := s.workOrders.Finish(req)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, order)
}

// ListAllWorkOrders list all work orders with filters by status and date
// @Summary List all work orders
// @Description List all work orders with filters by status and date
// @Tags Work orders
// @Produce json
// @Param until query string false "iso date"
// @Param since query string false "iso date"
// @Param status query string false "new, cancelled, done"
// @Success 200 {object} []models.WorkOrder
// @Router /workorders [get]
func (s *server) ListAllWorkOrders(c echo.Context) error {
	req := models.ListAllWorkOrders{
		Since:  c.QueryParam("since"),
		Until:  c.QueryParam("until"),
		Status: c.QueryParam("status"),
	}

	orders, err := s.workOrders.ListAll(req)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, orders)
}

// ListWorkOrdersByCustomer list all work orders by a customer
// @Summary List work orders by customer
// @Description List all work orders by a customer
// @Tags Work orders
// @Produce json
// @Param customerId path string true "uuid"
// @Success 200 {object} []models.WorkOrder
// @Router /workorders/customer/{customerId} [get]
func (s *server) ListWorkOrdersByCustomer(c echo.Context) error {
	customerId := c.Param("customerId")

	orders, err := s.workOrders.ListByCustomer(customerId)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, orders)
}
