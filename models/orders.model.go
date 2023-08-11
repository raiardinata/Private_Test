package models

import (
	"errors"
	"net/http"
	"strconv"

	"private_test/db"
)

func CreateOrders(customer_id, order_id, product_id string) (Response, error) {
	var obj DBCheck
	var arrobj []DBCheck
	var res Response
	var rowAffected int64

	con := db.PgNewSession()
	tx, err := con.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return res, err
	}

	if customer_id == "" {
		return res, errors.New("customer_id is empty, error creating order")
	}
	if order_id == "" {
		return res, errors.New("order_id is empty, error creating order")
	}
	if product_id == "" {
		return res, errors.New("product_id is empty, error creating order")
	}

	// insert to orders table
	sqlRes, err := tx.InsertInto("orders").Columns("order_id", "customer_id", "status").Values(order_id, customer_id, "fresh_order").Exec()
	rowAffected, _ = sqlRes.RowsAffected()
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Messages = "rowAffected when Create order : " + strconv.FormatInt(rowAffected, 10)
		return res, err
	}
	obj.Id = "customer_id : " + customer_id + "; "
	obj.Id += "status : fresh_order; "
	arrobj = append(arrobj, obj)

	// insert to Order_Products table
	sqlRes2, err2 := tx.InsertInto("order_products").Columns("order_id", "product_id").Values(order_id, product_id).Exec()
	rowAffected, _ = sqlRes2.RowsAffected()
	if err2 != nil {
		res.Status = http.StatusBadRequest
		res.Messages = "rowAffected when Create order_products : " + strconv.FormatInt(rowAffected, 10)
		return res, err
	}
	tx.Commit()
	obj.Id = "order_id : " + order_id + "; "
	obj.Id += "product_id : " + product_id + ";"
	arrobj = append(arrobj, obj)

	res.Status = http.StatusOK
	res.Messages = "Create Order Success!"
	res.Data = arrobj

	return res, nil
}
