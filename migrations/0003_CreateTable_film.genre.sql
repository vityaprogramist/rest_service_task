-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- schema 'film' for 'genre' table
CREATE SCHEMA IF NOT EXISTS film;

-- drop sequence if it exists
DROP SEQUENCE IF EXISTS film.genre_id_seq;
-- create sequence for column 'id' in 'genre' table
CREATE SEQUENCE film.genre_id_seq
  INCREMENT 1
  MINVALUE 1
  NO MAXVALUE
  START 1
  CACHE 1;

DROP TABLE IF EXISTS film.genre;
CREATE TABLE IF NOT EXISTS film.genre (
  id    SMALLINT DEFAULT nextval('film.genre_id_seq') NOT NULL,
  name  VARCHAR NOT NULL
)
WITH (OIDS = FALSE);

ALTER SEQUENCE film.genre_id_seq OWNED BY film.genre.id;
ALTER TABLE film.genre ADD PRIMARY KEY (id);
ALTER TABLE film.genre ADD CONSTRAINT unique_genre UNIQUE (name);
ALTER TABLE film.genre OWNER TO postgres;
CREATE UNIQUE INDEX pk_genre_id ON film.genre USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE film.genre DROP CONSTRAINT unique_genre;
DROP INDEX IF EXISTS film.pk_genre_id;
DROP TABLE IF EXISTS film.genre;