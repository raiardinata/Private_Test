package models

import (
	"errors"
	"net/http"

	"private_test/db"
	"private_test/modelstruct"
)

func GetOrders(customer_id string) (Response, error) {
	var res Response
	var co []modelstruct.CoustomerOrder
	var customerValid []modelstruct.CoustomerValid

	con := db.PgNewSession()
	tx, err := con.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return res, err
	}

	if customer_id == "" {
		return res, errors.New("customer_id is empty")
	}

	tx.Select("name").
		From("customers").
		Where("customers.customer_id = ?", customer_id).Load(&customerValid)
	if customerValid == nil {
		res.Status = http.StatusBadRequest
		res.Messages = "Customer Not Valid!"
		return res, errors.New("customer not valid")
	}

	// list for admin
	if customerValid[0].CustomerName == "sup_admin" {
		tx.Select("customers.customer_id AS customer_id, customers.name AS customer_name, customers.email AS customer_email,	products.name AS product_name, products.price AS product_price, products.description AS product_description, orders.order_id,	orders.status AS order_status").
			From("orders").
			Join("order_products", "order_products.order_id = orders.order_id").
			Join("customers", "customers.customer_id = orders.customer_id").
			Join("products", "products.product_id = order_products.product_id").
			Load(&co)

		res.Status = http.StatusOK
		res.Messages = "Select Order for Admin Success!"
		res.Data = co

		return res, nil
	}

	// list for customer
	tx.Select("customers.customer_id AS customer_id, customers.name AS customer_name, customers.email AS customer_email,	products.name AS product_name, products.price AS product_price, products.description AS product_description, orders.order_id,	orders.status AS order_status").
		From("orders").
		Join("order_products", "order_products.order_id = orders.order_id").
		Join("customers", "customers.customer_id = orders.customer_id").
		Join("products", "products.product_id = order_products.product_id").
		Where("orders.customer_id = ?", customer_id).Load(&co)

	res.Status = http.StatusOK
	res.Messages = "Select Order Success!"
	res.Data = co

	return res, nil
}
