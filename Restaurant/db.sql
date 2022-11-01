drop database d;
create database d;
use d;
CREATE TABLE Menu(
item_id int primary key  ,
type varchar(100)not null,
price varchar(11) not null 
);
-- insert into Student value(1,'negm',0111);
CREATE TABLE customer(
customer_id int primary key  ,
name varchar(100)not null,
phone varchar(11) not null,
address varchar(250) 
);
CREATE TABLE Orders(
order_id int primary key  ,
address varchar(100)not null,
description varchar(500),
date datetime,
status bool not null,
customerid int,
itemid int,
FOREIGN KEY(customerid) REFERENCES customer(customer_id),
FOREIGN KEY(itemid) REFERENCES Menu(item_id)
);
-- insert into Book value(1,'database',4);

create table Ticket(
	ticket_id int primary key,
	price int,
    date datetime,
    customerrid int,
	FOREIGN KEY(customerrid) REFERENCES customer(customer_id)
 );
 
