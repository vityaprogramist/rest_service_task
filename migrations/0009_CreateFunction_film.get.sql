-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION film.get(VARCHAR = NULL, VARCHAR = NULL, INT = NULL)   RETURNS TABLE (id INT, name VARCHAR, genre_list TEXT, release_year INT) AS $$ BEGIN   CREATE TEMP TABLE IF NOT EXISTS temp_table AS     SELECT       film.film.id, film.film.title, string_agg(film.genre.name, ', ') AS genre_list, film.film.release_year::INT     FROM       film.film     INNER JOIN       film.film_genre     ON       film.film.id = film.film_genre.film_id     INNER JOIN       film.genre     ON       film.genre.id = film.film_genre.genre_id     GROUP BY film.id;    IF ($1 IS NOT NULL) THEN     DELETE FROM temp_table WHERE temp_table.title != $1;   END IF;    IF ($2 IS NOT NULL) THEN     DELETE FROM temp_table WHERE strpos(temp_table.genre_list, $2) = 0;   END IF;    IF ($3 IS NOT NULL AND $3 > 0) THEN     DELETE FROM temp_table WHERE  temp_table.release_year != $3;   END IF;    RETURN QUERY     SELECT       *     FROM       temp_table;   DROP TABLE temp_table; END; $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS film.get(VARCHAR, VARCHAR, INT);
