create table product
(
    id       uuid primary key,
    name     varchar,
    category varchar,
    cost     int
);

-- 1-index

create index index_name on product(name);
drop index index_name;
-- 2-index

create index index_id_cost on product(id, cost);

-- 3-index

create index index_category_name on product(name, category);

explain (analyse )
select * from product where name  = 'Elsa' and category = 'Becker' -- id = 'b2457480-75c7-481b-8b76-c40795ec8ff0';
