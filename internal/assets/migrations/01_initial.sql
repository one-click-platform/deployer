-- +migrate Up

create table accounts (
    id bigserial primary key,
    email text not null unique,
    password_hash bytea not null
);

-- +migrate Down

drop table accounts;
