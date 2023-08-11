package modelstruct

import "time"

type ExportOrder struct {
	CustomerID		string `db:"customer_id"`
	CustomerName 	string `db:"customer_name"`

	TotalPrice	float64

	OrderID     string `db:"order_id"`
	OrderStatus string `db:"order_status"`
	OrderDate   time.Time `db:"order_date"`
}
