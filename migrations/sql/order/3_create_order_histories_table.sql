-- +migrate Up
CREATE TABLE IF NOT EXISTS order_histories (
    id bigint(20) unsigned primary key,
    order_id bigint(20) unsigned not null,
    status varchar(50) not null,
    remark varchar(150) null,
    created_at timestamp not null default current_timestamp(),

    key status (status),

    foreign key (order_id)
        references orders (id)
        on update cascade
        on delete no action
) engine = InnoDB
    auto_increment = 0
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS order_histories;