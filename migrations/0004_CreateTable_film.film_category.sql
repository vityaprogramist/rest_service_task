-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- schema 'film' for 'genre' table
CREATE SCHEMA IF NOT EXISTS film;

DROP TABLE IF EXISTS film.film_genre;
CREATE TABLE IF NOT EXISTS film.film_genre (
  film_id     INTEGER NOT NULL,
  genre_id    SMALLINT NOT NULL
)
WITH (OIDS = FALSE);

ALTER TABLE film.film_genre ADD PRIMARY KEY (film_id, genre_id);
ALTER TABLE film.film_genre
  ADD CONSTRAINT film_genre_genre_id_fkey
    FOREIGN KEY (genre_id) REFERENCES film.genre(id) ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE film.film_genre
  ADD CONSTRAINT film_genre_film_id_fkey
    FOREIGN KEY (film_id) REFERENCES film.film(id) ON UPDATE CASCADE ON DELETE RESTRICT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- ALTER TABLE film.genre DROP CONSTRAINT unique_genre;
-- DROP INDEX IF EXISTS film.pk_genre_id;
ALTER TABLE film.film_genre DROP CONSTRAINT film_genre_film_id_fkey;
ALTER TABLE film.film_genre DROP CONSTRAINT film_genre_genre_id_fkey;
DROP TABLE IF EXISTS film.film_genre;