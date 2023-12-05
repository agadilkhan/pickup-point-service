create table user_tokens
(
    id            bigserial
        primary key,
    user_id       bigint,
    token         text,
    refresh_token text,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now()
);