-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.rented_by_user_id(user_id BIGINT, IN page_size INT, IN page_number INT)   RETURNS TABLE (film_id INT, title VARCHAR, genres TEXT, release_year INT) AS $$ DECLARE   page INT; BEGIN   IF (page_size IS NOT NULL AND page_number IS NOT NULL) THEN     page := (page_size * (page_number - 1));   ELSE     RAISE EXCEPTION 'passed wrong arguments.';   END IF;    RETURN QUERY     SELECT       film.film.id,       film.film.title,       string_agg(film.genre.name, ', ') AS genres,       film.film.release_year::INT     FROM       film.film     INNER JOIN       film.film_genre     ON       film.film.id = film.film_genre.film_id     INNER JOIN       film.genre     ON       film.genre.id = film.film_genre.genre_id     INNER JOIN       rental.rental     ON       film.film.id = rental.rental.film_id     INNER JOIN       users.user     ON       users.user.id = rental.rental.user_id     WHERE users.user.id = $1     GROUP BY users.user.id, film.film.id     ORDER BY film.film.id     LIMIT page_size     OFFSET page; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.rented_by_user_id(BIGINT, INT, INT);
