package sendemail

import (
	"fmt"

	"private_test/db"
	"private_test/modelstruct"

	"gopkg.in/gomail.v2"
)

type CustomerID struct {
	CustomerID string `db:"customer_id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
}

const (
	gmailSMTPHost     = "smtp.gmail.com"
	gmailSMTPPort     = 587
	gmailSMTPUsername = "semtepedetesyingsemtepede@gmail.com" // change it to your gmail email
	gmailSMTPPassword = "iglkrkbvctetxtex"                    // change it to your gmail app password
)

func sendReminderEmail(name, toEmail, products, checkoutURL string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "private_test@gmail.com")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reminder: Pending Order")
	body := fmt.Sprintf("Hello %s,\n\nYou have a pending order with the following products:\n\n%s\n\nTo complete your order, please proceed to the checkout:\n%s\n\nBest regards,\nYour Online Store", name, products, checkoutURL)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(gmailSMTPHost, gmailSMTPPort, gmailSMTPUsername, gmailSMTPPassword)
	return d.DialAndSend(m)
}

func SendEmail() {
	var customerID []CustomerID
	var co []modelstruct.CoustomerOrder
	var prdName string

	con := db.PgNewSession()
	tx, err := con.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		fmt.Println("Something wrong with send email")
	}

	tx.Select("customer_id, name, email").
		From("customers").Load(&customerID)
	for _, cust := range customerID {
		// list for customer
		tx.Select("customers.customer_id AS customer_id, customers.name AS customer_name, customers.email AS customer_email,	products.name AS product_name, products.price AS product_price, products.description AS product_description, orders.order_id,	orders.status AS order_status").
			From("orders").
			Join("order_products", "order_products.order_id = orders.order_id").
			Join("customers", "customers.customer_id = orders.customer_id").
			Join("products", "products.product_id = order_products.product_id").
			Where("orders.customer_id = ?", cust.CustomerID).
			Where("orders.status = ?", "fresh_order").Load(&co)
		for _, co := range co {
			prdName += co.ProductName + ", "
		}
		if co != nil {
			sendReminderEmail(cust.Name, cust.Email, prdName, "test.url.com")
		}
	}
}
