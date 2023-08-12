package routes

import (
	"net/http"
	"sync"
	"time"

	"private_test/handlers"
	sendemail "private_test/sendEmail"

	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

// Implement an API rate limiter that limits the number of requests per minute from a single IP address to 100.
type rateLimiter struct {
	ips map[string]int
	mu  sync.Mutex
}

func (rl *rateLimiter) addIP(ip string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.ips[ip]++
}

func (rl *rateLimiter) isRateExceeded(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	count, exists := rl.ips[ip]
	if !exists || count <= 100 {
		return false
	}

	return true
}

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	rl := &rateLimiter{ips: make(map[string]int)}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			if rl.isRateExceeded(ip) {
				return c.JSON(http.StatusTooManyRequests, "Rate limit exceeded")
			}
			rl.addIP(ip)
			return next(c)
		}
	})
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Initialize Echo Framework!")
	})

	e.GET("/_health", handlers.HealthCheck)

	// Define the endpoint to place an order
	e.POST("/createOrders", handlers.CreateOrders)
	e.GET("/getOrders", handlers.GetOrders)
	e.GET("/exportOrder", handlers.ExportOrders)
	e.GET("/sendEmail", handlers.SendEmail)

	// Route to manually trigger the email sending (for testing purposes)
	e.GET("/send_reminders", handlers.SendEmail)

	go func() {
		for {
			// Replace this with your action to be performed every 24 hours
			sendemail.SendEmail()

			// Wait for 24 hours before executing the action again
			time.Sleep(24 * time.Hour)
		}
	}()

	return e
}
