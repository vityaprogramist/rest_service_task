-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS "users"."users" (
  id              BIGSERIAL PRIMARY KEY NOT NULL,
  name            VARCHAR NOT NULL,
  login           VARCHAR NOT NULL UNIQUE,
  password_hash   VARCHAR NOT NULL,
  age             INT NOT NULL,
  phone           BIGINT NOT NULL,
  CHECK (age > 0),
  CHECK (phone > 0)
)
WITH (OIDS = FALSE);

CREATE UNIQUE INDEX pk_users_ids ON users.users USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS "users"."users";
DROP SCHEMA IF EXISTS "users";
