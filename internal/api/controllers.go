package api

import (
	"net/http"

	"github.com/ansel1/merry"
	"github.com/labstack/echo/v4"
)

// List list customers
// @Summary List customers
// @Description List customers
// @Tags Customers
// @Produce json
// @Param title query string false "search by title"
// @Success 200 {object} []models.Customer
// @Router /customers [get]
func (s *server) ListCustomers(c echo.Context) error {
	customers, err := s.customers.List()
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, customers)
}
