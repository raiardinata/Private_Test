package handlers

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"private_test/models"

	"github.com/labstack/echo"
)

// Insert Into orders and order_products
func CreateOrders(c echo.Context) error {
	customer_id := c.FormValue("customer_id")
	product_id := c.FormValue("product_id")
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	order_id := strings.Replace(string(newUUID[:]), "\n", "", -1)

	result, err := models.CreateOrders(customer_id, order_id, product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
