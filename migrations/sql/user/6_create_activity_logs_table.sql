-- +migrate Up
CREATE TABLE IF NOT EXISTS activity_logs (
    id int(11) unsigned auto_increment primary key,
    user_id bigint(20) not null,
    ip_address varchar(50),
    user_agent varchar(50),
    path_endpoint varchar(150) not null,
    created_at timestamp not null default current_timestamp(),

    key user_id (user_id)
) engine = InnoDB
    auto_increment = 1
    default charset = utf8
    collate = utf8_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS activity_logs;