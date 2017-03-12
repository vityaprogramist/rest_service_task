-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

COPY film.film FROM 'D:/Go/gopath/src/github.com/rest_service_task/migrations/data/films.dat' WITH DELIMITER ';' CSV;
COPY film.genre FROM 'D:/Go/gopath/src/github.com/rest_service_task/migrations/data/genres.dat' WITH DELIMITER ';' CSV;
COPY film.film_genre FROM 'D:/Go/gopath/src/github.com/rest_service_task/migrations/data/film_genres.dat' WITH DELIMITER ';' CSV;
COPY users.user FROM 'D:/Go/gopath/src/github.com/rest_service_task/migrations/data/users.dat' WITH DELIMITER ';' CSV;
COPY rental.rental FROM 'D:/Go/gopath/src/github.com/rest_service_task/migrations/data/rental.dat' WITH DELIMITER ';' CSV;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM rental.rental;
DELETE FROM users.user;
DELETE FROM film.film_genre;
DELETE FROM film.genre;
DELETE FROM film.film;
