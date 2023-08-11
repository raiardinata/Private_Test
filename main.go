package main

import (
	"time"

	"private_test/db"
	"private_test/routes"

	_ "github.com/lib/pq"
)

// Order represents the structure of an order
type Order struct {
	OrderID    int       `json:"order_id"`
	CustomerID int       `json:"customer_id"`
	OrderDate  time.Time `json:"order_date"`
	Status     string    `json:"status"`
	ProductIDs []int     `json:"product_ids"`
}

func main() {
	// Create a database connection
	db.Init()

	// Initialize Echo
	e := routes.Init()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
