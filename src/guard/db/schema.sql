create table grants
(
    id           uuid               default gen_random_uuid(),
    user_id      bigint    not null,
    action       smallint  not null,
    foreign_id   bigint    not null,
    ip           inet      not null,
    user_agent   text      not null,
    otp_id       bigint,
    token        text,
    void         bool      not null default false,
    voided_by_id bigint,
    finalised_at timestamp,
    created_at   timestamp not null,
    expires_at   timestamp not null,

    primary key (id),

    unique (otp_id),
    unique (token)
);

create index idx_grants_action_foreign_id on grants (action, foreign_id);
create index idx_grants_id_ip on grants (id, ip);

create table otps
(
    id              bigserial,
    channel         smallint  not null,
    email           text,
    phone           text,
    code            text,
    send_count      smallint  not null,
    last_sent_at    timestamp,
    attempts        smallint  not null default 0,
    last_attempt_at timestamp,
    created_at      timestamp not null,

    primary key (id)
);
