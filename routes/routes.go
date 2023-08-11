package routes

import (
	"net/http"

	"private_test/handlers"

	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Initialize Echo Framework!")
	})

	e.GET("/_health", handlers.HealthCheck)

	// Define the endpoint to place an order
	e.POST("/createOrders", handlers.CreateOrders)
	e.GET("/getOrders", handlers.GetOrders)

	return e
}
