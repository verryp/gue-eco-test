-- +migrate Up
CREATE TABLE IF NOT EXISTS item_quotas (
    id int(11) unsigned auto_increment primary key,
    item_id bigint(20) unsigned not null,
    date_limiter timestamp not null default current_timestamp(),
    quota_remaining int(11) not null,

    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null,

    key deleted_at (deleted_at),

    foreign key (item_id)
        references items (id)
        on update cascade
        on delete no action
) engine = InnoDB
    auto_increment = 1
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS item_quotas;