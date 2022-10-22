DO
    $do$ DECLARE
    _db TEXT := 'weather';
    _user TEXT := 'postgres';
    _password TEXT := 'masterkey';
BEGIN
CREATE
EXTENSION IF NOT EXISTS dblink;
IF EXISTS (SELECT 1 FROM pg_database WHERE datname = _db) THEN
    RAISE NOTICE 'Database already exists';
ELSE
    PERFORM dblink_connect('host=localhost user=' || _user || ' password=' || _password || ' dbname=' || current_database());
PERFORM dblink_exec('CREATE DATABASE ' || _db);
END IF;
END
$do$;