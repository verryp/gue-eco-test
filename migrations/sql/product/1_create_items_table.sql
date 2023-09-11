-- +migrate Up
CREATE TABLE IF NOT EXISTS items (
    id bigint(20) unsigned primary key,
    name varchar(150) not null,
    quota_per_days int(11) not null,
    quantity int(11) not null,
    category varchar(50) not null,
    price decimal(20, 2) not null,

    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null,

    key deleted_at (deleted_at)
) engine = InnoDB
    auto_increment = 0
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS items;