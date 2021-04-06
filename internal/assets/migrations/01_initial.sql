-- +migrate Up

create table accounts (
    id bigserial primary key,
    login varchar(50) not null,
    password varying(1024) not null,
);

-- +migrate Down

drop table settings;
