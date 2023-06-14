create TABLE if not exists users
(
    id         UUID        DEFAULT gen_random_uuid() not null unique,
    login      TEXT                                  not null unique,
    password   TEXT                                  not null,
    created_at TIMESTAMPTZ default now()
);

create index if not exists index_login_users on users (login);