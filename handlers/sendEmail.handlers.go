package handlers

import (
	"net/http"

	sendemail "private_test/sendEmail"

	"github.com/labstack/echo"
)

// get customers order
func SendEmail(c echo.Context) error {
	sendemail.SendEmail()
	return c.JSON(http.StatusOK, nil)
}
