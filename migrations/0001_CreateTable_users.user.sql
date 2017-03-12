-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- schema 'users' for 'user' table
CREATE SCHEMA IF NOT EXISTS users;
COMMENT ON SCHEMA users IS 'schema for storing ''users'' table and private functions.';

-- drop sequence if it exists
-- DROP SEQUENCE IF EXISTS users.user_id_seq;

-- create sequence for column 'id' in 'users' table
CREATE SEQUENCE IF NOT EXISTS users.user_id_seq
  INCREMENT 1
  MINVALUE 1
  NO MAXVALUE
  START 1
  CACHE 1;

-- drop table if it exists
DROP TABLE IF EXISTS users.user;

-- create 'users' table
CREATE TABLE users.user (
  id              BIGINT DEFAULT nextval('users.user_id_seq') NOT NULL,
  first_name      VARCHAR NOT NULL,
  last_name       VARCHAR NOT NULL,
  login           VARCHAR NOT NULL,
  password_hash   VARCHAR NOT NULL,
  age             INT NOT NULL,
  phone           BIGINT NOT NULL
)
WITH (OIDS = FALSE);

ALTER SEQUENCE users.user_id_seq OWNED BY users.user.id;
ALTER TABLE users.user ADD PRIMARY KEY (id);
ALTER TABLE users.user ADD CONSTRAINT unique_login UNIQUE (login);
ALTER TABLE users.user ADD CONSTRAINT age_check CHECK (age > 0);
ALTER TABLE users.user ADD CONSTRAINT phone_check CHECK (phone > 0);
ALTER TABLE users.user OWNER TO postgres;
CREATE UNIQUE INDEX pk_user_id ON users.user USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users.user DROP CONSTRAINT age_check;
ALTER TABLE users.user DROP CONSTRAINT phone_check;
ALTER TABLE users.user DROP CONSTRAINT unique_login;
DROP INDEX IF EXISTS users.pk_user_id;
DROP TABLE IF EXISTS users.user;
DROP SEQUENCE IF EXISTS users.user_id_seq;
DROP SCHEMA IF EXISTS users;