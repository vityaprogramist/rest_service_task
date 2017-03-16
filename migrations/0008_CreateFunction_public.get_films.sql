-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

 CREATE FUNCTION get_films (page_size integer DEFAULT NULL::integer, page_number integer DEFAULT NULL::integer, film_genre character varying DEFAULT NULL::character varying, film_release_year integer DEFAULT NULL::integer) RETURNS TABLE(id integer, name character varying, genre text, release_year integer) 	LANGUAGE plpgsql AS $$  DECLARE   page INT;  BEGIN    IF (COALESCE(page_size, 0) > 0 AND COALESCE(page_number, 0) > 0) THEN      page := (page_size * (page_number - 1));    ELSE      RAISE EXCEPTION 'Неверные параметры запроса.';    END IF;    RETURN QUERY       SELECT * FROM film.get(NULL, $3, $4)       ORDER BY id       LIMIT page_size       OFFSET page; END; $$

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.get_films(INT,INT,VARCHAR,INT);
