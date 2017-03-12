-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.get_films(IN page_size INT = NULL, IN page_number INT = NULL, IN film_genre VARCHAR = NULL, IN film_release_year INT = NULL)   RETURNS TABLE(id INT, name VARCHAR, genre TEXT, release_year INT) AS $$ DECLARE   page INT; BEGIN   IF (page_size IS NOT NULL AND page_number IS NOT NULL) THEN     page := (page_size * (page_number - 1));   ELSE     RAISE EXCEPTION 'passed wrong arguments.';   END IF;    RETURN QUERY     SELECT * FROM film.get(NULL, $3, $4)     ORDER BY id     LIMIT page_size     OFFSET page; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.get_films(INT,INT,VARCHAR,INT);
