CREATE TABLE users
(
    id            uuid primary key,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
    created_at    timestamp(6) not null,
    updated_at    timestamp(6) not null,
);

CREATE TABLE subscriptions
(
    id            uuid primary key,
    user_id       string not null,
    created_at    timestamp(6) not null,
    updated_at    timestamp(6) not null,
);

-- create table images
-- (
--     id                   bigserial
--         primary key,
--     listing_id           bigint
--         constraint fk_rails_2a2257c8bb
--             references listings,
--     external_link        varchar,
--     created_at           timestamp(6) not null,
--     updated_at           timestamp(6) not null,
--     failed_to_load_count integer default 0,
--     status               integer default 0
-- );