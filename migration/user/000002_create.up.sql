create table roles
(
    id          bigserial
        primary key,
    name        varchar(50) not null,
    description text
);

create table users
(
    id           bigserial
        primary key,
    role_id      bigint not null
        constraint fk_users_role
            references roles
            on update cascade on delete cascade,
    first_name   varchar(255),
    last_name    varchar(255),
    email        varchar(255) UNIQUE,
    phone        varchar(15),
    login        varchar(255) UNIQUE,
    password     varchar(255),
    is_confirmed boolean                  default false,
    is_deleted   boolean,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);
