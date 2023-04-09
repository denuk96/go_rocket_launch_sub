CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id            uuid primary key DEFAULT uuid_generate_v4(),
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at    timestamp(6) not null,
    updated_at    timestamp(6) not null
);

CREATE TABLE subscriptions
(
    id            uuid primary key DEFAULT uuid_generate_v4(),
    user_id       varchar(255) not null,
    created_at    timestamp(6) not null,
    updated_at    timestamp(6) not null
);
