create table logins
(
    id          bigserial,
    user_id     bigint       not null,
    ip          inet         not null,
    user_agent  varchar(255) not null,
    created_at  timestamp    not null,

    primary key (id)
);

create index idx_logins_user_id on logins (user_id);

create table sessions
(
    id             char(64)  not null,
    user_id        bigint    not null,
    login_id       bigint,
    grant_id       text,
    last_active_at timestamp not null,
    last_active_ip inet      not null,
    created_at     timestamp not null,
    signed_out_at  timestamp,

    primary key (id),
    unique (login_id)
);
