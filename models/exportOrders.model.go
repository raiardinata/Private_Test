package models

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"private_test/db"
	"private_test/modelstruct"
)

func ExportOrders(w http.ResponseWriter, r *http.Request) (Response, error) {
	var res Response
	var eo []modelstruct.ExportOrder

	con := db.PgNewSession()
	tx, err := con.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Messages = "Export Order Failed. RollbackUnlessCommitted."
		res.Data = eo
		return res, err
	}

	tx.Select("orders.order_id, customers.customer_id, customers.name AS customer_name, orders.order_date, orders.status AS order_status").
		From("orders").
		Join("customers", "customers.customer_id = orders.customer_id").
		OrderBy("orders.order_id ASC").Load(&eo)


	// Calculate total price for each order
	for i := range eo {
		orderID := eo[i].OrderID
		var total float64
		_, err = tx.Select("SUM(products.price)").
			From("order_products").
			Join("products", "products.product_id = order_products.product_id").
			Where("order_products.order_id = ?", orderID).
			Load(&total)
		if err != nil {
			res.Status = http.StatusBadRequest
			res.Messages = "Export Order Failed. Fail to select total price."
			res.Data = eo
			return res, err
		}
		eo[i].TotalPrice = total
	}

	// Create a CSV file
	file, err := os.Create("./order_report.csv")
	if err != nil {
		res.Status = http.StatusBadRequest
			res.Messages = "Export Order Failed. Fail to create order report."
			res.Data = eo
			return res, err
	}
	defer file.Close()

	// Write CSV header
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Order ID", "Customer Name", "Order Date", "Total Price", "Status"})

	// Write order data to CSV
	for _, eo := range eo {
		writer.Write([]string{
			fmt.Sprintf("%v", eo.OrderID),
			eo.CustomerName,
			eo.OrderDate.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%.2f", eo.TotalPrice),
			eo.OrderStatus,
		})
	}

	writer.Flush()
	file.Close()

	// Serve the CSV file for download
	w.Header().Set("Content-Disposition", "attachment; filename=order_report.csv")
	http.ServeFile(w, r, file.Name())

	res.Status = http.StatusOK
	res.Messages = "Select Order Success!"
	res.Data = eo
	return res, nil
}
