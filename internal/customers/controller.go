package customers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	services Services
}

type Controller interface {
	List(c echo.Context) error
}

func NewController(s Services) Controller {
	return &controller{
		services: s,
	}
}

// List list customers
// @Summary List customers
// @Description List customers
// @Tags Customers
// @Produce json
// @Param title query string false "search by title"
// @Success 200 {object} []customer
// @Router /customers [get]
func (ctr *controller) List(c echo.Context) error {
	customers, err := ctr.services.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, customers)
}
