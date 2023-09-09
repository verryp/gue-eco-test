-- +migrate Up
CREATE TABLE IF NOT EXISTS clients (
    id int(11) unsigned auto_increment primary key,
    `name` varchar(50) not null,
    api_key varchar(32) not null unique,
    algorithm varchar(5) not null,
    location varchar(100) not null,
    public_cert varchar(100) not null,
    private_cert varchar(100) not null,
    created_at timestamp not null default current_timestamp(),
    updated_at timestamp not null default current_timestamp() on update current_timestamp(),
    deleted_at timestamp null
) engine = InnoDB
    auto_increment = 1
    default charset = utf8
    collate = utf8_unicode_ci;

INSERT INTO
    clients (name, api_key, algorithm, location, public_cert, private_cert, created_at)
VALUES
    ("GUE Ecosystem", "116wIdYjuZUEF0OrJpaGIFP099uwhSXF", "RS256", "/go/src/github.com/verryp/gue-eco-test", "/deployment/auth/sample_key.pub", "/deployment/auth/sample_key.txt", now());

-- +migrate Down
DROP TABLE IF EXISTS clients;