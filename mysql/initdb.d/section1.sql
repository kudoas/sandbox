DROP DATABASE IF EXISTS section1;
CREATE DATABASE section1;
USE section1;

DROP TABLE IF EXISTS order_details;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
name VARCHAR(100) NOT NULL,
address VARCHAR(100),
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(id)
);

CREATE TABLE items (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
name VARCHAR(100) NOT NULL,
price INT UNSIGNED NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(id)
);

CREATE TABLE orders (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
order_date DATE NOT NULL,
customer_id INT UNSIGNED NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(id),
FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE order_details (
order_id INT UNSIGNED NOT NULL,
item_id INT UNSIGNED NOT NULL,
item_quantity INT UNSIGNED NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(order_id,item_id),
FOREIGN KEY (order_id) REFERENCES orders(id),
FOREIGN KEY (item_id) REFERENCES items(id)
);

insert into customers(id,name,address,created_at,updated_at) values(1,'A商事','東京都',now(),now()),(2,'B商会','埼玉県',now(),now()),(3,'C商店','神奈川県',now(),now());
insert into items(id,name,price,created_at,updated_at) values(1,'シャツ',1000,now(),now()),(2,'パンツ',950,now(),now()),(3,'マフラー',1200,now(),now()),(4,'ブルゾン',1800,now(),now());
insert into orders(id,order_date,customer_id,created_at,updated_at) values(1 , '2013-10-01',1,now(),now()),(2 , '2013-10-01',2,now(),now()),(3 , '2013-10-02',2,now(),now()),(4 , '2013-10-02',3,now(),now());
insert into order_details(order_id,item_id,item_quantity,created_at,updated_at) values(1 , 1 ,3,now(),now()),(1 , 2 ,2,now(),now()),(2 , 1 ,1,now(),now()),(2 , 3 ,10,now(),now()),(2 , 4 ,5,now(),now()),(3 , 2 ,80,now(),now()),(4 , 3 ,25,now(),now());