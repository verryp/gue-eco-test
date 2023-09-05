-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigint(20) unsigned primary key,
    `name` varchar(150) not null,
    email varchar(150) not null unique,
    password varchar(100) not null,
    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null,

    key deleted_at (deleted_at)
) engine = InnoDB
    auto_increment = 0
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS users;