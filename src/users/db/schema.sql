create schema users;

create table users
(
    id             bigint    not null,
    type           smallint  not null,
    email          text,
    email_verified boolean   not null default false,
    phone          text,
    phone_verified boolean   not null default false,
    password       text,
    signup_ip      inet,
    created_at     timestamp not null,

    primary key (id),
    unique (type, email),
    unique (type, phone)
);

create table users.password_reset_requests
(
    id         bigserial,
    user_id    bigint    not null,
    identifier text,
    ip         inet      not null,
    user_agent text,
    created_at timestamp not null,

    primary key (id)
);

create table users.password_resets
(
    id           bigserial,
    user_id      bigint    not null,
    old_password text,
    new_password text      not null,
    request_id   bigint,
    grant_id     uuid,
    created_at   timestamp not null,

    primary key (id),
    unique (grant_id),
    unique (request_id)
);
