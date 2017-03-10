-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SCHEMA films;

CREATE TYPE films.genre_type AS ENUM (
  'comedy',
  'action',
  'drama',
  'horror',
  'crime',
  'historical',
  'science',
  'western',
  'adventure',
  'fantasy'
);

CREATE TABLE IF NOT EXISTS "films"."films" (
  id        BIGSERIAL PRIMARY KEY NOT NULL,
  name      VARCHAR NOT NULL,
  genre     films.genre_type[],
  release   TIMESTAMP
)
WITH (OIDS = FALSE);

CREATE UNIQUE INDEX pk_films_ids ON films.films USING BTREE (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS "films"."films";
DROP TYPE IF EXISTS "films"."genre_type";
DROP SCHEMA IF EXISTS "films";
