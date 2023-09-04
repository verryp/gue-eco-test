-- +migrate Up
CREATE TABLE IF NOT EXISTS order_details (
    id int(11) unsigned auto_increment primary key,
    order_id bigint(20) unsigned not null,
    item_id bigint(20) not null,
    item_name varchar(150) not null,
    quantity int(11) not null,
    item_price decimal(20, 2) not null,
    total_amount decimal(20, 2) not null,
    customer_note text null,

    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null,

    key deleted_at (deleted_at),
    key item_id (item_id),

    foreign key (order_id)
        references orders (id)
        on update cascade
        on delete no action
) engine = InnoDB
    auto_increment = 1
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS order_details;