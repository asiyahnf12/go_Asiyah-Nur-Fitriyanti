CREATE TABLE users (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    status smallint,
    gender char(1),
    dob DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE address (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    street varchar(255),
    city varchar(255),
    province varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE operators (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_types (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    payment_method int(11),
    status varchar(10),
    total_qty int(11),
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(payment_method) REFERENCES payment_methods(id)
);

CREATE TABLE products (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    product_type_id int(11),
    operator_id int(11),
    code varchar(50),
    name varchar(100),
    status smallint,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_type_id) REFERENCES product_types(id),
    FOREIGN KEY(operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    product_id int(11),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_id) REFERENCES products(id) 
);

CREATE TABLE transaction_details (
  id int(11) not null AUTO_INCREMENT PRIMARY KEY,
  transaction_id int(11),
  product_id int(11),
  status varchar(10),
  qty int(11),
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE payment_methods (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    status smallint,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payment_method_descriptions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    payment_method_id int(11),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE user_payment_method_details (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    payment_method int(11),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method) REFERENCES payment_methods(id)
);

CREATE TABLE shipping (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    ongkos_dasar DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO operators (name) values ('TOKO Asiyah');
INSERT INTO operators (name) values ('TOKO Purnama');
INSERT INTO operators (name) values ('TOKO Ajaib');
INSERT INTO operators (name) values ('TOKO Barokah');
INSERT INTO operators (name) values ('TOKO Sekawan');

use alta_online_shop;

INSERT INTO transactions_details (product_id, qty)
VALUES (1, 3),(2, 10),(3, 7);

INSERT INTO transactions (user_id, total_qty, total_price) VALUES 
(1, '18', 100000),
(2, '19', 250000),
(3, '20', 500000);

INSERT INTO product_types (name) values ('Elektronik');
INSERT INTO product_types (name) values ('Kerudung');
INSERT INTO product_types (name) values ('Sepatu');

INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (1,3,'PHILIP2', 'Lampu taman 5 Watt', 1),(2,1,'RABBANI3', 'Bergo', 1),(3,4,'VPN3', 'Flatshoes', 1);

INSERT INTO products_descriptions (product_id, description)
VALUES (1, "Lampu taman sangat bagus"),
(2, "Bergo sangat simple"), (3, "Flatshoes sangat nyaman dipakai");

INSERT INTO payment_methods (name, status) values ('BRI',1),('DANA',2),('Shopeepay',3);

INSERT INTO users (status, gender, dob)
VALUES ('1', 'L', '1995-01-01'),
    ('1', 'P', '1998-03-14'),
    ('1', 'L', '2000-05-20'),
    ('1', 'P', '1989-09-10'),
    ('1', 'L', '1992-11-30');

SELECT status FROM users WHERE gender='L';

SELECT * FROM products WHERE id=3;

SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND dob LIKE '%2000%';

SELECT * FROM users WHERE gender='P';

SELECT * FROM users ORDER BY status ASC;

SELECT * FROM products LIMIT 5;

UPDATE products
SET name = 'product dummy'
WHERE id = 1;

UPDATE transactions_details
SET qty = 3
WHERE product_id = 1;

DELETE FROM products
WHERE operator_id = 1;

DELETE FROM products
WHERE product_type_id = 1;

SELECT * FROM transactions
WHERE user_id = 1
UNION ALL
SELECT * FROM transactions
WHERE user_id = 2;

SELECT SUM(total_price) as total_price FROM transactions
WHERE user_id = 1
GROUP BY user_id;

SELECT COUNT(*) as total_transactions FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2
GROUP BY p.product_type_id;

SELECT products.*, product_types.name FROM products
JOIN product_types ON products.product_type_id = product_types.id;

SELECT transactions.*, products.*, users.* FROM transactions
JOIN transactions_details ON transactions.id = transactions_details.transactions_id
JOIN products ON transactions_details.product_id = products.id
JOIN users ON transactions.user_id = users.id;

CREATE TRIGGER delete_transaction_details AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details
    WHERE transaction_id = OLD.id;
END;

CREATE TRIGGER update_total_qty AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE id = OLD.transaction_id;
END;

SELECT * FROM products
WHERE id NOT IN (SELECT DISTINCT product_id FROM transactions_details);
