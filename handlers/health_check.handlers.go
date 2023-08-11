package handlers

import (
	"net/http"
	"private_test/models"

	"github.com/labstack/echo"
)

// HealthCheck gets health check of the db
func HealthCheck(c echo.Context) error {
	result, err := models.HealthCheck()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
