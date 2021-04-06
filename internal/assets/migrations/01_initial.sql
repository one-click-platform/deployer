-- +migrate Up

create table accounts (
    id bigserial primary key,
    email varchar(25) not null unique,
    password bytea not null
);

-- +migrate Down

drop table accounts;
