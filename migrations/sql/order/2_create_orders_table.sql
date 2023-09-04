-- +migrate Up
CREATE TABLE IF NOT EXISTS orders (
    id bigint(20) unsigned primary key,
    order_serial varchar(100) not null,
    customer_name varchar(150) not null,
    customer_email varchar(150) not null,
    status varchar(50) not null,
    total_amount decimal(20, 2) not null,

    expired_at timestamp not null,
    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null,

    unique order_serial_id (id, order_serial),
    key deleted_at (deleted_at),
    key status (status)
) engine = InnoDB
    auto_increment = 0
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS orders;