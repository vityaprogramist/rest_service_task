-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE OR REPLACE FUNCTION public.end_rent(f_id INT, u_id BIGINT) RETURNS VOID AS $$ BEGIN   IF EXISTS(SELECT * FROM rental.rental WHERE film_id = f_id AND user_id = u_id) THEN     DELETE FROM rental.rental WHERE film_id = $1 AND user_id = $2;   ELSE     RAISE EXCEPTION 'Произошла ошибка во время возврата фильма.';   END IF; END $$ LANGUAGE plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION IF EXISTS public.end_rent(INT, BIGINT);
