-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- schema 'users' for 'users' table
CREATE SCHEMA IF NOT EXISTS film;
COMMENT ON SCHEMA film IS 'schema for storing ''users'' table, private functions and auxiliary entities.';

-- drop sequence if it exists
DROP SEQUENCE IF EXISTS film.film_id_seq;
-- create sequence for column 'film_id' in 'film' table
CREATE SEQUENCE film.film_id_seq
  INCREMENT 1
  MINVALUE 1
  NO MAXVALUE
  START 1
  CACHE 1;

-- create especially domain for using in 'film' table as release year value
CREATE DOMAIN film.Year AS INT
  CONSTRAINT ckeck_year CHECK (VALUE > 1900 AND VALUE <= date_part('year', 'now'::DATE));


DROP TABLE IF EXISTS film.film;
CREATE TABLE IF NOT EXISTS film.film (
  id              INTEGER DEFAULT nextval('film.film_id_seq') NOT NULL,
  title           VARCHAR NOT NULL,
  release_year    film.Year
)
WITH (OIDS = FALSE);

ALTER SEQUENCE film.film_id_seq OWNED BY film.film.id;
ALTER TABLE film.film ADD PRIMARY KEY (id);
ALTER TABLE film.film ADD CONSTRAINT unique_film UNIQUE (title, release_year);
ALTER TABLE film.film OWNER TO postgres;
CREATE UNIQUE INDEX pk_film_id ON film.film USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE film.film DROP CONSTRAINT unique_film;
DROP INDEX IF EXISTS film.pk_film_id;
DROP TABLE IF EXISTS film.film;
DROP DOMAIN IF EXISTS film.Year;
DROP SCHEMA IF EXISTS film;
