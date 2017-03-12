-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- schema 'film' for 'genre' table
CREATE SCHEMA IF NOT EXISTS rental;

DROP TABLE IF EXISTS rental.rental;
CREATE TABLE IF NOT EXISTS rental.rental(
  film_id     INTEGER NOT NULL,
  user_id     BIGINT NOT NULL
)
WITH (OIDS = FALSE);

ALTER TABLE rental.rental ADD PRIMARY KEY (film_id, user_id);

ALTER TABLE rental.rental
  ADD CONSTRAINT rental_film_id_fkey
    FOREIGN KEY (film_id) REFERENCES film.film(id) ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE rental.rental
  ADD CONSTRAINT rental_user_id_fkey
    FOREIGN KEY (user_id) REFERENCES users.user(id) ON UPDATE CASCADE ON DELETE RESTRICT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE rental.rental DROP CONSTRAINT rental_film_id_fkey;
ALTER TABLE rental.rental DROP CONSTRAINT rental_user_id_fkey;
DROP TABLE IF EXISTS rental.rental;