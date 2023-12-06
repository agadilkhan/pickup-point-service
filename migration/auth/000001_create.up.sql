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

create table outbox_messages
(
    id           bigserial
        primary key,
    user_email   text,
    code         text,
    is_processed boolean
);
