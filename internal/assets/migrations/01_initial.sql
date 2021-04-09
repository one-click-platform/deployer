-- +migrate Up

create table accounts (
    id bigserial primary key,
    email text not null unique,
    password_hash bytea not null
);

create table environments (
    id bigserial primary key,
    name text not null,
    account_id bigserial,
    foreign key (account_id) references accounts (id),
    unique(name, account_id)
);

-- +migrate Down

drop table accounts;
drop table projects;