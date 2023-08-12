# Private_Test
Task :
1. 
-- Database: cognotiv_test

-- DROP DATABASE IF EXISTS cognotiv_test;

CREATE DATABASE cognotiv_test
    WITH
    OWNER = rai
    ENCODING = 'UTF8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE Customers (
    customer_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE Products (
    product_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT,
    image_url VARCHAR(255)
);

CREATE TABLE Orders (
    order_id VARCHAR(255) PRIMARY KEY,
    customer_id uuid REFERENCES Customers(customer_id) ON DELETE CASCADE,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) NOT NULL
);

CREATE TABLE Order_Products (
    order_id VARCHAR(255) REFERENCES Orders(order_id) ON DELETE CASCADE,
    product_id uuid REFERENCES Products(product_id) ON DELETE CASCADE,
    PRIMARY KEY (order_id, product_id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO customers(name, email, password) VALUES('sup_admin', 'admin@gmail.com', 'sup_admin');

INSERT INTO products(name, price, description, image_url) VALUES('chairWhite', '1', 'normal chair, white color.', 'test.com/normalChairWhite.jpeg');

INSERT INTO products(name, price, description, image_url) VALUES('chairBlack', '1', 'normal chair, black color.', 'test.com/normalChairBlack.jpeg');

INSERT INTO products(name, price, description, image_url) VALUES('chairBlue', '1', 'normal chair, blue color.', 'test.com/normalChairBlue.jpeg');

INSERT INTO Orders(customer_id, status) VALUES('ab2ea48e-180b-44e5-8909-6784e9a4a27b', 'fresh_order');

INSERT INTO Order_Products(order_id, product_id) VALUES('b061d55a-559b-43a7-a5d0-87ca23fa181e', '152943cc-8c39-4bc1-b881-c6948f0391ca');

2.
Use postman, use post method with this url http://localhost:8080/createOrders. Add the param customer_id string from table customers.customer_id and product_id from table products.product_id

3. 
Use postman, use get method with this url http://localhost:8080/getOrders. Add the param customer_id string from table customers.customer_id

4. 
Use postman, use get method with this url http://localhost:8080/getOrders. Add the param customer_id string of **admin customer_id** from table customers.customer_id

5.
Check on folder routes/routes.go line 74-79

6.
Use postman, use get method with this url http://localhost:8080/exportOrder without param.

7. 
Check on folder routes/routes.go line 15-54
