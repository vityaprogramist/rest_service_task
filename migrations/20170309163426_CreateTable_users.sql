-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SCHEMA users;

CREATE TABLE IF NOT EXISTS "users"."users" (
  id        BIGSERIAL PRIMARY KEY NOT NULL,
  name      VARCHAR NOT NULL,
  login     VARCHAR NOT NULL,
  password  VARCHAR NOT NULL,
  age       INT,
  phone     BIGINT
)
WITH (OIDS = FALSE);

CREATE UNIQUE INDEX pk_users_ids ON users.users USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS "users"."users";
DROP SCHEMA IF EXISTS "users";
