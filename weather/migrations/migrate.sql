/*DO
    $do$
    DECLARE
  _db TEXT := 'weather';
_user TEXT := 'postgres';
_password TEXT := 'masterkey';
BEGIN
CREATE EXTENSION IF NOT EXISTS dblink; -- enable extension
IF EXISTS (SELECT 1 FROM pg_database WHERE datname = _db) THEN
    RAISE NOTICE 'Database already exists';
ELSE
    PERFORM dblink_connect('host=localhost user=' || _user || ' password=' || _password || ' dbname=' || current_database());
PERFORM dblink_exec('CREATE DATABASE ' || _db);
END IF;
END
$do$;*/

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