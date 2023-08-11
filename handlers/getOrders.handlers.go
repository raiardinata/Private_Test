package handlers

import (
	"net/http"

	"private_test/models"

	"github.com/labstack/echo"
)

// get customers order
func GetOrders(c echo.Context) error {
	customer_id := c.FormValue("customer_id")

	result, err := models.GetOrders(customer_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
