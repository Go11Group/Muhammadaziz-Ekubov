create table Users (
    id int not null ,
    firstname varchar(50) default 'Noname',
    Lastname varchar(50) default 'Noname',
    age int default 18
);

create table Products (
    id int not null ,
    name varchar(100) ,
    price float4 not null ,
    quantity int not null
);

INSERT INTO Users (id, firstname, Lastname, age) VALUES
    (1, 'Alice', 'Smith', 25),
    (2, 'Bob', 'Johnson', 30),
    (3, 'Charlie', 'Williams', 18), -- age default value
    (4, 'David', 'Jones', 18), -- age default value
    (5, 'Emma', 'Brown', 22),
    (6, 'Frank', 'Davis', 40),
    (7, 'Grace', 'Miller', 18), -- age default value
    (8, 'Hannah', 'Wilson', 18), -- age default value
    (9, 'Ian', 'Moore', 28),
    (10, 'Jane', 'Taylor', 35);


INSERT INTO Products (id, name, price, quantity) VALUES
    (1, 'Laptop', 999.99, 10),
    (2, 'Smartphone', 499.99, 20),
    (3, 'Tablet', 299.99, 15),
    (4, 'Smartwatch', 199.99, 25),
    (5, 'Headphones', 99.99, 50),
    (6, 'Speaker', 149.99, 30),
    (7, 'Monitor', 249.99, 8),
    (8, 'Keyboard', 49.99, 40),
    (9, 'Mouse', 29.99, 60),
    (10, 'Charger', 19.99, 100);
