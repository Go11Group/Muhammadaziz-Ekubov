create type ic_dc as enum ('credit', 'debit');

create table if not exists Users (
    id uuid primary key default gen_random_uuid(),
    name varchar(50),
    phone varchar(12),
    age int
);

create table if not exists Card(
    id uuid primary key default gen_random_uuid(),
    number varchar(16),
    user_id uuid references Users(id)
);

create table if not exists Transaction (
    id uuid primary key default gen_random_uuid(),
    card_id uuid references Card(id),
    amount int default 0,
    transaction_id uuid default null,
    transaction_type ic_dc
);

create table if not exists Station (
    id uuid primary key default gen_random_uuid(),
    name varchar
);

create table if not exists Terminal (
    id uuid primary key default gen_random_uuid(),
    station_id uuid references  Station(id)
);

--Insert

INSERT INTO Users (name, phone, age) VALUES
    ('Alice Johnson', '5551234567', 30),
    ('Bob Smith', '5552345678', 45),
    ('Charlie Brown', '5553456789', 25),
    ('Diana Prince', '5554567890', 35),
    ('Ethan Hunt', '5555678901', 40),
    ('Fiona Gallagher', '5556789012', 28),
    ('George Costanza', '5557890123', 37),
    ('Helen Mirren', '5558901234', 42),
    ('Ian Fleming', '5559012345', 33),
    ('Jack Sparrow', '5550123456', 50);

INSERT INTO Card (number, user_id) VALUES
    ('1234567812345678', (SELECT id FROM Users WHERE name = 'Alice Johnson')),
    ('2345678923456789', (SELECT id FROM Users WHERE name = 'Bob Smith')),
    ('3456789034567890', (SELECT id FROM Users WHERE name = 'Charlie Brown')),
    ('4567890145678901', (SELECT id FROM Users WHERE name = 'Diana Prince')),
    ('5678901256789012', (SELECT id FROM Users WHERE name = 'Ethan Hunt')),
    ('6789012367890123', (SELECT id FROM Users WHERE name = 'Fiona Gallagher')),
    ('7890123478901234', (SELECT id FROM Users WHERE name = 'George Costanza')),
    ('8901234589012345', (SELECT id FROM Users WHERE name = 'Helen Mirren')),
    ('9012345690123456', (SELECT id FROM Users WHERE name = 'Ian Fleming')),
    ('0123456701234567', (SELECT id FROM Users WHERE name = 'Jack Sparrow'));

INSERT INTO Transaction (card_id, amount, transaction_type) VALUES
    ((SELECT id FROM Card WHERE number = '1234567812345678'), 10000, 'credit'),
    ((SELECT id FROM Card WHERE number = '1234567812345678'), 5000, 'debit'),
    ((SELECT id FROM Card WHERE number = '2345678923456789'), 20000, 'credit'),
    ((SELECT id FROM Card WHERE number = '2345678923456789'), 7500, 'debit'),
    ((SELECT id FROM Card WHERE number = '3456789034567890'), 30000, 'credit'),
    ((SELECT id FROM Card WHERE number = '3456789034567890'), 15000, 'debit'),
    ((SELECT id FROM Card WHERE number = '4567890145678901'), 40000, 'credit'),
    ((SELECT id FROM Card WHERE number = '4567890145678901'), 12500, 'debit'),
    ((SELECT id FROM Card WHERE number = '5678901256789012'), 50000, 'credit'),
    ((SELECT id FROM Card WHERE number = '5678901256789012'), 20000, 'debit'),
    ((SELECT id FROM Card WHERE number = '6789012367890123'), 60000, 'credit'),
    ((SELECT id FROM Card WHERE number = '6789012367890123'), 30000, 'debit'),
    ((SELECT id FROM Card WHERE number = '7890123478901234'), 70000, 'credit'),
    ((SELECT id FROM Card WHERE number = '7890123478901234'), 25000, 'debit'),
    ((SELECT id FROM Card WHERE number = '8901234589012345'), 80000, 'credit'),
    ((SELECT id FROM Card WHERE number = '8901234589012345'), 35000, 'debit'),
    ((SELECT id FROM Card WHERE number = '9012345690123456'), 90000, 'credit'),
    ((SELECT id FROM Card WHERE number = '9012345690123456'), 40000, 'debit'),
    ((SELECT id FROM Card WHERE number = '0123456701234567'), 100000, 'credit'),
    ((SELECT id FROM Card WHERE number = '0123456701234567'), 45000, 'debit');

INSERT INTO Station (name) VALUES
    ('Central Station'),
    ('North Station'),
    ('East Station'),
    ('South Station'),
    ('West Station');

INSERT INTO Terminal (station_id) VALUES
    ((SELECT id FROM Station WHERE name = 'Central Station')),
    ((SELECT id FROM Station WHERE name = 'North Station')),
    ((SELECT id FROM Station WHERE name = 'East Station')),
    ((SELECT id FROM Station WHERE name = 'South Station')),
    ((SELECT id FROM Station WHERE name = 'West Station'));
