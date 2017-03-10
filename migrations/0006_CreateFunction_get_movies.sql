-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION get_movies(IN page_size INT = NULL, IN page_number INT = NULL, IN filter VARCHAR = NULL)   RETURNS TABLE(id BIGINT, name VARCHAR, genre films.genre_type[], release TIMESTAMP) AS $$ DECLARE   page BIGINT; BEGIN   IF (page_size IS NOT NULL AND page_number IS NOT NULL) THEN     page := (page_size * (page_number - 1));   END IF;    IF (filter IS NOT NULL) THEN     RETURN QUERY       SELECT         films.films.id, films.films.name, films.films.genre, films.films.release       FROM         films.films       WHERE         filter = ANY(films.films.genre::VARCHAR[])       ORDER BY films.films.id       LIMIT page_size       OFFSET page;   ELSE     RETURN QUERY       SELECT         films.films.id, films.films.name, films.films.genre, films.films.release       FROM         films.films       ORDER BY films.films.id       LIMIT page_size       OFFSET page;   END IF;   EXCEPTION     WHEN OTHERS THEN       RAISE; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS get_movies(INT,INT,VARCHAR);
