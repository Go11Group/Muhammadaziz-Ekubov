create table cards(
    id uuid primary key default gen_random_uuid() not null,
    number int not null,
    user_id uuid not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deteled_at timestamp
);

create type transaction_types as enum('credit', 'debit');

create table transactions(
    id uuid primary key default gen_random_uuid() not null,
    card_id uuid references card(id) not null,
    amount int not null,
    terminal_id uuid default null,
    transaction_type transaction_types not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deteled_at timestamp
);

create table stations(
    id uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deteled_at timestamp
);

create table terminals(
    id uuid primary key default gen_random_uuid() not null,
    station_id uuid not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deteled_at timestamp
);