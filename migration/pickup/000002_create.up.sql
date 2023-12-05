create table customers
(
    id         bigserial
        primary key,
    first_name varchar(255),
    last_name  varchar(255),
    email      varchar(255),
    phone      varchar(15)
);

create table companies
(
    id            bigserial
        primary key,
    name          varchar(255),
    contact_email varchar(255),
    contact_phone varchar(255)
);

create table pickup_points
(
    id      bigserial
        primary key,
    name    varchar(255),
    address varchar(255)
);

create table orders
(
    id           bigserial
        primary key,
    customer_id  bigint
        constraint fk_orders_customer
            references customers
            on update cascade on delete cascade,
    company_id   bigint
        constraint fk_orders_company
            references companies
            on update cascade on delete cascade,
    point_id     bigint
        constraint fk_orders_point
            references pickup_points
            on update cascade on delete cascade,
    code         varchar(50),
    status       varchar(50),
    is_paid      boolean,
    total_amount numeric,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);

create table products
(
    id          bigserial
        primary key,
    name        varchar(255),
    description text,
    price       numeric
);

create table order_items
(
    id            bigserial
        primary key,
    order_id      bigint
        constraint fk_orders_order_items
            references orders,
    product_id    bigint
        constraint fk_order_items_product
            references products
            on update cascade on delete cascade,
    quantity      bigint,
    sub_total     numeric,
    is_accept     boolean,
    num_of_refund bigint
);

create table transactions
(
    id               bigserial
        primary key,
    user_id          bigint,
    order_id         bigint
        constraint fk_transactions_order
            references orders,
    transaction_type text,
    created_at       timestamp with time zone default now(),
    updated_at       timestamp with time zone default now()
);
