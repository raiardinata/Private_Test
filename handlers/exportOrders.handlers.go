package handlers

import (
	"net/http"

	"private_test/models"

	"github.com/labstack/echo"
)

// get customers order
func ExportOrders(c echo.Context) error {
	w := c.Response()
	r := c.Request()
	result, err := models.ExportOrders(w,r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
