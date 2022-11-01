#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -U postgres -d weather <<-EOSQL
CREATE TABLE IF NOT EXISTS cities
(
    city_id serial primary key,
    name varchar(50) not null,
    lat real  not null,
    lon real  not null,
    country varchar(2)  not null,
    state varchar(255) not null,
    UNIQUE (name, country, state)
);

create table if not exists weather
(
    id serial primary key,
    city_id int references cities (city_id),
    date varchar(10) not null,
    temp real not null,
    detail JSONB not null
);

create table if not exists favor
(
    id serial primary key,
    user_name text unique not null,
    favor text[] not null default '{}'::text[]
);
EOSQL