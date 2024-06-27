create table users(
    id uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    phone varchar not null,
    age int not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deteled_at timestamp
);