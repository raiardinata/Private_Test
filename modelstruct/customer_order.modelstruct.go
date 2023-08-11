package modelstruct

type CoustomerOrder struct {
	CustomerID string `db:"customer_id"`
	Name       string `db:"customer_name"`
	Email      string `db:"customer_email" csv:"customer_email"`

	ProductID    string `db:"product_id" csv:"product_id"`
	ProductName  string `db:"product_name" csv:"product_name"`
	ProductPrice string `db:"product_price" csv:"product_price"`
	ProductDesc  string `db:"product_description" csv:"product_description"`

	OrderID     string `db:"order_id" csv:"order_id"`
	OrderStatus string `db:"order_status" csv:"order_status"`
}

type CoustomerValid struct {
	CustomerName string `db:"name"`
}
