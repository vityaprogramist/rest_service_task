-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SCHEMA IF NOT EXISTS rent;

CREATE TABLE IF NOT EXISTS rent.users_films (
  id_user BIGINT,
  id_film BIGINT,
  CONSTRAINT fk_users FOREIGN KEY (id_user) REFERENCES users.users(id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_films FOREIGN KEY (id_film) REFERENCES films.films(id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT uniq UNIQUE (id_user,id_film)
)
WITH (OIDS=FALSE);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS rent.users_films;

DROP SCHEMA IF EXISTS rent;
